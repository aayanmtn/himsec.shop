package ui

import (
	"fmt"
	"strings"

	"himsec.shop/models"
	"himsec.shop/styles"
)

func RenderMainView(products []models.Product, selected int, category string) string {
	var sb strings.Builder

	// Add extra padding before logo
	sb.WriteString("\n")

	// Render logo with style
	sb.WriteString(styles.LogoStyle.Render(Logo))
	sb.WriteString("\n")

	// Get unique categories
	categories := getUniqueCategories(products)

	// Render products by category
	for _, cat := range categories {
		// Category header
		sb.WriteString(styles.CategoryStyle.Render("— " + cat))
		sb.WriteString("\n")

		// Products in this category
		for _, product := range products {
			if product.Category == cat {
				price := fmt.Sprintf("$%.2f", product.Price)
				wishlistMark := " "
				if product.WishList {
					wishlistMark = "♥"
				}
				productLine := fmt.Sprintf("%s %s %s ★",
					product.Name,
					styles.PriceStyle.Render(price),
					wishlistMark)
				sb.WriteString(styles.ProductStyle.Render(productLine))
				sb.WriteString("\n")
			}
		}
		sb.WriteString("\n")
	}

	// Render controls
	sb.WriteString(styles.ProductStyle.Render("[q] Quit | [↑/↓] Navigate | [enter] Select | [w] Wishlist | [b] Buy"))

	return sb.String()
}

func RenderDetailView(product models.Product) string {
	var sb strings.Builder

	// Add extra padding before logo
	sb.WriteString("\n")

	// Render logo with style
	sb.WriteString(styles.LogoStyle.Render(Logo))
	sb.WriteString("\n")

	// Product details
	sb.WriteString(styles.CategoryStyle.Render(product.Name))
	sb.WriteString("\n")
	sb.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Price: $%.2f", product.Price)))
	sb.WriteString("\n")
	sb.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Description: %s", product.Description)))
	sb.WriteString("\n")
	if product.WishList {
		sb.WriteString(styles.ProductStyle.Render("Status: In Wishlist ♥"))
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

	// Controls
	sb.WriteString(styles.ProductStyle.Render("[b] Back | [w] Toggle Wishlist | [p] Purchase | [q] Quit"))

	return sb.String()
}

func RenderCheckoutView(products []models.Product) string {
	var sb strings.Builder

	// Add extra padding before logo
	sb.WriteString("\n")

	// Render logo with style
	sb.WriteString(styles.LogoStyle.Render(Logo))
	sb.WriteString("\n")

	// Calculate total
	var total float64
	var itemCount int
	sb.WriteString(styles.CategoryStyle.Render("— Cart"))
	sb.WriteString("\n")

	for _, product := range products {
		if product.WishList {
			itemCount++
			total += product.Price
			sb.WriteString(styles.ProductStyle.Render(fmt.Sprintf("%s %s",
				product.Name,
				styles.PriceStyle.Render(fmt.Sprintf("$%.2f", product.Price)))))
			sb.WriteString("\n")
		}
	}

	if itemCount == 0 {
		sb.WriteString(styles.ProductStyle.Render("Cart is empty"))
		sb.WriteString("\n")
	} else {
		sb.WriteString("\n")
		sb.WriteString(styles.CategoryStyle.Render("— Summary"))
		sb.WriteString("\n")
		sb.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Total Items: %d", itemCount)))
		sb.WriteString("\n")
		sb.WriteString(styles.ProductStyle.Render(fmt.Sprintf("Total: %s",
			styles.PriceStyle.Render(fmt.Sprintf("$%.2f", total)))))
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	// Controls
	sb.WriteString(styles.ProductStyle.Render("[b] Back | [p] Proceed to Payment | [q] Quit"))

	return sb.String()
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