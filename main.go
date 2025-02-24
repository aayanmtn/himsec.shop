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
}

func NewProgram() *Program {
	return &Program{
		currentView: "main",
		products:    models.InitializeProducts(),
		selected:    0,
	}
}

func (p *Program) View() string {
	switch p.currentView {
	case "main":
		return ui.RenderMainView(p.products, p.selected, "all")
	case "detail":
		return ui.RenderDetailView(p.products[p.selected])
	case "checkout":
		return ui.RenderCheckoutView(p.products) // Assumed function in ui package
	default:
		return ui.RenderMainView(p.products, p.selected, "all")
	}
}

func (p *Program) HandleInput(input string) {
	switch p.currentView {
	case "main":
		switch input {
		case "up", "p":
			if p.selected > 0 {
				p.selected--
			}
		case "down", "n":
			if p.selected < len(p.products)-1 {
				p.selected++
			}
		case "enter", "d":
			p.currentView = "detail"
		case "w":
			// Toggle wishlist for selected product
			p.products[p.selected].WishList = !p.products[p.selected].WishList
		case "b":
			p.currentView = "checkout"
		}
	case "detail":
		switch input {
		case "b":
			p.currentView = "main"
		case "w":
			// Toggle wishlist for current product
			p.products[p.selected].WishList = !p.products[p.selected].WishList
		case "p":
			// TODO: Implement payment gateway integration
			p.currentView = "checkout"
		}
	case "checkout":
		// Add checkout handling logic here (e.g., processing order)
		if input == "b" {
			p.currentView = "main"
		}

	}
}