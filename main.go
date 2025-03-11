package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bwish "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/muesli/termenv"

	// Local imports - assuming these packages exist in your project
	"himsec.shop/models"
	"himsec.shop/ui"
)

const (
	host = "0.0.0.0"
	port = "22"
)

type model struct {
	products     []models.Product
	wishes       []models.Wish
	selected     int
	currentView  string
	width        int
	height       int
	border       lipgloss.Border
	name         string
	address      string
	phone        string
	country      string
	state        string
	city         string
	currentField int
}

func InitialModel() model {
	return model{
		products:    models.InitializeProducts(),
		selected:    0,
		currentView: "main",
		border:      lipgloss.RoundedBorder(),
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
					if m.currentView == "checkout" && m.currentField > 0 {
						m.currentField--
					} else if m.currentView == "main" && m.selected > 0 {
						m.selected--
					}
				case "down", "j":
					if m.currentView == "checkout" {
						m.currentField = (m.currentField + 1) % 6
					} else if m.currentView == "main" && m.selected < len(m.products)-1 {
						m.selected++
					}
				case "tab":
					if m.currentView == "checkout" {
						m.currentField = (m.currentField + 1) % 6
					}
				case "enter":
					m.currentView = "detail"
				case "w":
					if m.currentView == "main" || m.currentView == "detail" {
						newWish := models.NewWish(m.products[m.selected].Name)
						m.wishes = append(m.wishes, newWish)
					}
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
				default:
					if m.currentView == "checkout" {
						if msg.Type == tea.KeyRunes {
							switch m.currentField {
								case 0:
									m.name += msg.String()
								case 1:
									m.address += msg.String()
								case 2:
									m.phone += msg.String()
								case 3:
									m.country += msg.String()
								case 4:
									m.state += msg.String()
								case 5:
									m.city += msg.String()
							}
						} else if msg.Type == tea.KeyBackspace {
							switch m.currentField {
								case 0:
									if len(m.name) > 0 {
										m.name = m.name[:len(m.name)-1]
									}
								case 1:
									if len(m.address) > 0 {
										m.address = m.address[:len(m.address)-1]
									}
								case 2:
									if len(m.phone) > 0 {
										m.phone = m.phone[:len(m.phone)-1]
									}
								case 3:
									if len(m.country) > 0 {
										m.country = m.country[:len(m.country)-1]
									}
								case 4:
									if len(m.state) > 0 {
										m.state = m.state[:len(m.state)-1]
									}
								case 5:
									if len(m.city) > 0 {
										m.city = m.city[:len(m.city)-1]
									}
							}
						}
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
			s = ui.RenderCheckoutView(m.products[m.selected:m.selected+1], m.wishes, m.currentField, m.name, m.address, m.phone, m.country, m.state, m.city)
	}
	return fmt.Sprintf("\n%s\n", s)
}

func shopBubbleteaMiddleware() wish.Middleware {
	newProg := func(m tea.Model, opts ...tea.ProgramOption) *tea.Program {
		return tea.NewProgram(m, opts...)
	}

	teaHandler := func(s ssh.Session) *tea.Program {
		pty, _, active := s.Pty()
		if !active {
			wish.Fatalln(s, "no active terminal, skipping")
			return nil
		}

		m := InitialModel()
		m.width = pty.Window.Width
		m.height = pty.Window.Height

		return newProg(m, append(bwish.MakeOptions(s), tea.WithAltScreen())...)
	}

	return bwish.MiddlewareWithProgramHandler(teaHandler, termenv.ANSI256)
}

func main() {
	// Ensure the SSH host key exists
	keyPath := ".ssh/id_ed25519"
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		if _, err := os.Stat(".ssh"); os.IsNotExist(err) {
			if err := os.Mkdir(".ssh", 0700); err != nil {
				log.Fatalf("Failed to create .ssh directory: %v", err)
			}
		}

		log.Printf("SSH host key not found at %s", keyPath)
		log.Printf("Generate a key with: ssh-keygen -t ed25519 -f %s -N \"\"", keyPath)
		log.Fatalf("SSH host key required")
	}

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
				 wish.WithHostKeyPath(keyPath),
				 wish.WithMiddleware(
					 shopBubbleteaMiddleware(),
						     logging.Middleware(),
				 ),
	)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH shop server on %s:%s", host, port)

	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Printf("Could not start server: %v", err)
			done <- nil
		}
	}()

	<-done
	log.Println("Stopping SSH server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Printf("Could not stop server: %v", err)
	}
}
