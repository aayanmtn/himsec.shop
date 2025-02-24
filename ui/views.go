package ui

import (
	"fmt"
	"strings"

	"himsec.shop/models"
	"himsec.shop/styles"
)

func RenderMainView(products []models.Product, selected int, category string) string {
	var sb strings.Builder

	// Render logo
	sb.WriteString(styles.TitleStyle.Render(Logo))
	sb.WriteString("\n")

	// Render category filter
	sb.WriteString(styles.InfoStyle.Render(fmt.Sprintf("Category: %s", category)))
	sb.WriteString("\n\n")

	// Render products
	for i, product := range products {
		if category != "all" && product.Category != category {
			continue
		}

		productText := fmt.Sprintf("%s - $%.2f", product.Name, product.Price)
		if i == selected {
			sb.WriteString(styles.SelectedStyle.Render("> " + productText))
		} else {
			sb.WriteString(styles.ProductStyle.Render("  " + productText))
		}
		sb.WriteString("\n")
	}

	// Render controls
	sb.WriteString("\n")
	sb.WriteString(styles.InfoStyle.Render("Controls: [n]ext [p]revious [d]etails [c]ategory [q]uit"))

	return sb.String()
}

func RenderDetailView(product models.Product) string {
	var sb strings.Builder

	// Render logo
	sb.WriteString(styles.TitleStyle.Render(Logo))
	sb.WriteString("\n")

	// Render product details
	sb.WriteString(styles.SelectedStyle.Render(product.Name))
	sb.WriteString("\n\n")
	sb.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Category: %s\n", product.Category)))
	sb.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Price: $%.2f\n", product.Price)))
	sb.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Description: %s\n", product.Description)))

	// Render controls
	sb.WriteString("\n")
	sb.WriteString(styles.InfoStyle.Render("Controls: [b]ack [q]uit"))

	return sb.String()
}