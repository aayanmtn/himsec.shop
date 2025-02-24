package main

import (
	"fmt"

	"himsec.shop/models"
	"himsec.shop/ui"
)

func main() {
	// Initialize the program state
	program := NewProgram()

	// Start the main loop
	for {
		// Clear screen
		fmt.Print("\033[H\033[2J")

		// Render current view
		fmt.Println(program.View())

		// Get input
		var input string
		fmt.Scanln(&input)

		// Handle input
		if input == "q" {
			break
		}
		program.HandleInput(input)
	}
}

type Program struct {
	currentView string
	products    []models.Product
	selected    int
	category    string
}

func NewProgram() *Program {
	return &Program{
		currentView: "main",
		products:    models.InitializeProducts(),
		selected:    0,
		category:    "all",
	}
}

func (p *Program) View() string {
	switch p.currentView {
	case "main":
		return ui.RenderMainView(p.products, p.selected, p.category)
	case "detail":
		return ui.RenderDetailView(p.products[p.selected])
	default:
		return ui.RenderMainView(p.products, p.selected, p.category)
	}
}

func (p *Program) HandleInput(input string) {
	switch p.currentView {
	case "main":
		switch input {
		case "n":
			if p.selected < len(p.products)-1 {
				p.selected++
			}
		case "p":
			if p.selected > 0 {
				p.selected--
			}
		case "d":
			p.currentView = "detail"
		case "c":
			p.cycleCategoryFilter()
		}
	case "detail":
		if input == "b" {
			p.currentView = "main"
		}
	}
}

func (p *Program) cycleCategoryFilter() {
	categories := []string{"all", "hardware", "software", "accessories"}
	for i, cat := range categories {
		if cat == p.category {
			p.category = categories[(i+1)%len(categories)]
			return
		}
	}
}