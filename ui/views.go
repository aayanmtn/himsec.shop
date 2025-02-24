package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/border"
	"github.com/charmbracelet/lipgloss"

	"himsec.shop/models"
	"himsec.shop/styles"
)

func RenderMainView(products []models.Product, selected int, category string) string {
	var sb strings.Builder

	// Add the logo
	sb.WriteString(styles.LogoStyle.Render(Logo))
	sb.WriteString("\n")

	// Get unique categories
	categories := getUniqueCategories(products)

	// Create a border style
	borderStyle := lipgloss.NewStyle().
		Border(border.Rounded).
		BorderForeground(styles.PrimaryColor)

	// Render products by category
	for _, cat := range categories {
		// Category header
		sb.WriteString(styles.CategoryStyle.Render("— " + cat))
		sb.WriteString("\n")

		// Products in this category
		for i, product := range products {
			if product.Category == cat {
				price := fmt.Sprintf("$%.2f", product.Price)
				wishlistMark := " "
				if product.WishList {
					wishlistMark = "♥"
				}

				productLine := fmt.Sprintf("%s %s %s %s",
					product.Name,
					styles.PriceStyle.Render(price),
					wishlistMark,
					styles.StarStyle.Render("★"))

				// Highlight selected item
				if i == selected {
					productLine = "> " + productLine
				} else {
					productLine = "  " + productLine
				}

				sb.WriteString(styles.ProductStyle.Render(productLine))
				sb.WriteString("\n")
			}
		}
		sb.WriteString("\n")
	}

	// Render controls with border
	controls := styles.ProductStyle.Render("[q] Quit | [↑/↓] Navigate | [enter] Select | [w] Wishlist | [b] Buy")
	sb.WriteString(borderStyle.Render(controls))

	return borderStyle.Render(sb.String())
}

func RenderDetailView(product models.Product) string {
	var sb strings.Builder

	// Add the logo
	sb.WriteString(styles.LogoStyle.Render(Logo))
	sb.WriteString("\n")

	// Create a border style
	borderStyle := lipgloss.NewStyle().
		Border(border.Rounded).
		BorderForeground(styles.PrimaryColor)

	// Product details
	details := fmt.Sprintf(`
%s

Price: %s
Description: %s
%s`,
		styles.CategoryStyle.Render(product.Name),
		styles.PriceStyle.Render(fmt.Sprintf("$%.2f", product.Price)),
		product.Description,
		product.WishList ? "Status: In Wishlist ♥" : "")

	sb.WriteString(borderStyle.Render(details))
	sb.WriteString("\n\n")

	// Controls
	controls := styles.ProductStyle.Render("[b] Back | [w] Toggle Wishlist | [p] Purchase | [q] Quit")
	sb.WriteString(borderStyle.Render(controls))

	return borderStyle.Render(sb.String())
}

func RenderCheckoutView(products []models.Product) string {
	var sb strings.Builder

	// Add the logo
	sb.WriteString(styles.LogoStyle.Render(Logo))
	sb.WriteString("\n")

	// Create a border style
	borderStyle := lipgloss.NewStyle().
		Border(border.Rounded).
		BorderForeground(styles.PrimaryColor)

	// Calculate total
	var total float64
	var itemCount int

	// Cart items
	cartContent := &strings.Builder{}
	cartContent.WriteString(styles.CategoryStyle.Render("— Cart"))
	cartContent.WriteString("\n")

	for _, product := range products {
		if product.WishList {
			itemCount++
			total += product.Price
			cartContent.WriteString(styles.ProductStyle.Render(fmt.Sprintf("%s %s",
				product.Name,
				styles.PriceStyle.Render(fmt.Sprintf("$%.2f", product.Price)))))
			cartContent.WriteString("\n")
		}
	}

	if itemCount == 0 {
		cartContent.WriteString(styles.ProductStyle.Render("Cart is empty"))
		cartContent.WriteString("\n")
	} else {
		cartContent.WriteString("\n")
		cartContent.WriteString(styles.CategoryStyle.Render("— Summary"))
		cartContent.WriteString("\n")
		cartContent.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Total Items: %d", itemCount)))
		cartContent.WriteString("\n")
		cartContent.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Total: %s",
			styles.PriceStyle.Render(fmt.Sprintf("$%.2f", total)))))
		cartContent.WriteString("\n")
	}

	sb.WriteString(borderStyle.Render(cartContent.String()))
	sb.WriteString("\n\n")

	// Controls
	controls := styles.ProductStyle.Render("[b] Back | [p] Proceed to Payment | [q] Quit")
	sb.WriteString(borderStyle.Render(controls))

	return borderStyle.Render(sb.String())
}

func getUniqueCategories(products []models.Product) []string {
	categories := make(map[string]bool)
	var result []string

	for _, p := range products {
		if !categories[p.Category] {
			categories[p.Category] = true
			result = append(result, p.Category)
		}
	}

	return result
}

var Logo = "  _   _      _ _         \n | | | | ___| | | ___ _ __ \n | |_| |/ _ \\ | |/ _ \\ '__|\n |  _  |  __/ | |  __/ |   \n |_| |_|\\___|_|_|\\___|_|   \n"