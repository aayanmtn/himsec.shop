package main

import (
	"fmt"
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"himsec.shop/models"
	"himsec.shop/ui"
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
		case "b":
			if m.currentView == "detail" {
				m.currentView = "main"
			} else if m.currentView == "main" {
				m.currentView = "checkout"
			}
		case "p":
			if m.currentView == "detail" {
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
		s = ui.RenderCheckoutView(m.products[m.selected:m.selected+1]) // Pass only selected product
	}

	return fmt.Sprintf("\n%s\n", s)
}

func main() {
	// Start HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HimSec Shop API Server Running")
	})
	go func() {
		log.Printf("Starting HTTP server on port 5000\n")
		if err := http.ListenAndServe("0.0.0.0:5000", nil); err != nil {
			log.Printf("HTTP server error: %v\n", err)
		}
	}()

	// Start terminal UI
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}