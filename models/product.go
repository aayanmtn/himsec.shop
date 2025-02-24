package models

type Product struct {
	Name        string
	Category    string
	Price       float64
	Description string
	SKU         string   // Unique product identifier for payment gateway
	ImageURL    string   // Product image URL
}

func InitializeProducts() []Product {
	return []Product{
		{
			Name:        "Vim Keychain",
			Category:    "accessories",
			Price:       9.99,
			Description: "Metal keychain with Vim logo",
			SKU:         "ACC-VIM-001",
		},
		{
			Name:        "Git Sticker Pack",
			Category:    "stickers",
			Price:       7.99,
			Description: "Set of 5 Git-themed vinyl stickers",
			SKU:         "STK-GIT-001",
		},
		{
			Name:        "Linux Penguin T-Shirt",
			Category:    "clothing",
			Price:       25.99,
			Description: "Premium Tux penguin t-shirt in black",
			SKU:         "CLT-LNX-001",
		},
		{
			Name:        "Firefox Logo Cap",
			Category:    "clothing",
			Price:       19.99,
			Description: "Baseball cap with Mozilla Firefox logo",
			SKU:         "CLT-FFX-001",
		},
	}
}