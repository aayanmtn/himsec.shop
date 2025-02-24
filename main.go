package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"

	"himsec.shop/models"
	"himsec.shop/ui"
)

const (
	host = "0.0.0.0"
	port = 2222
)

type model struct {
	products    []models.Product
	selected    int
	currentView string
	width      int
	height     int
	border     lipgloss.Border
}

func initialModel() model {
	return model{
		products:    models.InitializeProducts(),
		selected:    0,
		currentView: "main",
		border:     lipgloss.RoundedBorder(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			}
		case "down", "j":
			if m.selected < len(m.products)-1 {
				m.selected++
			}
		case "enter":
			m.currentView = "detail"
		case "w":
			m.products[m.selected].WishList = !m.products[m.selected].WishList
		case "b":
			if m.currentView == "detail" {
				m.currentView = "main"
			} else {
				m.currentView = "checkout"
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m model) View() string {
	s := ""
	switch m.currentView {
	case "main":
		s = ui.RenderMainView(m.products, m.selected, "all")
	case "detail":
		s = ui.RenderDetailView(m.products[m.selected])
	case "checkout":
		s = ui.RenderCheckoutView(m.products)
	}

	return fmt.Sprintf("\n%s\n", s)
}

func ensureSSHDirectory() error {
	sshDir := ".ssh"
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		return fmt.Errorf("failed to create .ssh directory: %w", err)
	}
	return nil
}

func main() {
	// Ensure .ssh directory exists
	if err := ensureSSHDirectory(); err != nil {
		log.Fatalf("Failed to setup SSH directory: %v", err)
	}

	hostKeyPath := filepath.Join(".ssh", "term_info_ed25519")

	// Set up SSH server with more detailed error handling
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatalf("Failed to create SSH server: %v", err)
	}

	// Log server startup
	log.Printf("Starting SSH server on %s:%d\n", host, port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("SSH server failed: %v", err)
	}
}

func teaHandler(s wish.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		fmt.Println("No active terminal, skipping")
		return nil, nil
	}

	m := initialModel()
	return m, []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
		tea.WithInput(pty),
		tea.WithOutput(pty),
	}
}