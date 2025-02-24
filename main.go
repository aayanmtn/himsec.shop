package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"himsec.shop/models"
	"himsec.shop/ui"
)

type model struct {
	products    []models.Product
	wishes      []models.Wish
	selected    int
	currentView string
	width      int
	height     int
	border     lipgloss.Border
	name       string
	address    string
	phone      string
	country    string
	state      string
	city       string
	currentField int
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

func main() {
	// Start SSH server in a goroutine
	go ssh.StartSSHServer()

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}