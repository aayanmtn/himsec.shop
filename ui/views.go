package ui

import (
	"fmt"
	"strings"

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
		Border(lipgloss.RoundedBorder()).
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
				productLine := fmt.Sprintf("%s %s %s",
					product.Name,
					styles.PriceStyle.Render(price),
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
	controls := styles.ProductStyle.Render("[q] Quit | [↑/↓] Navigate | [enter] Select | [b] Checkout")
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
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.PrimaryColor)

	// Product details
	details := fmt.Sprintf(`
%s

Price: %s
Description: %s`,
		styles.CategoryStyle.Render(product.Name),
		styles.PriceStyle.Render(fmt.Sprintf("$%.2f", product.Price)),
		product.Description)

	sb.WriteString(borderStyle.Render(details))
	sb.WriteString("\n\n")

	// Controls
	controls := styles.ProductStyle.Render("[b] Back | [p] Purchase | [q] Quit")
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
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.PrimaryColor)

	// Cart items
	cartContent := &strings.Builder{}
	cartContent.WriteString(styles.CategoryStyle.Render("— Selected Item"))
	cartContent.WriteString("\n")

	// Show selected product
	product := products[0] // For now, just show first product
	cartContent.WriteString(styles.ProductStyle.Render(fmt.Sprintf("%s %s",
		product.Name,
		styles.PriceStyle.Render(fmt.Sprintf("$%.2f", product.Price)))))
	cartContent.WriteString("\n\n")

	cartContent.WriteString(styles.CategoryStyle.Render("— Summary"))
	cartContent.WriteString("\n")
	cartContent.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Total: %s",
		styles.PriceStyle.Render(fmt.Sprintf("$%.2f", product.Price)))))
	cartContent.WriteString("\n")

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