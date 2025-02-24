package models

type Product struct {
	Name        string
	Category    string
	Price       float64
	Description string
}

func InitializeProducts() []Product {
	return []Product{
		{
			Name:        "Encrypted USB Drive",
			Category:    "hardware",
			Price:       49.99,
			Description: "256-bit AES encrypted USB drive with physical write-protect switch",
		},
		{
			Name:        "Security Key Bundle",
			Category:    "hardware",
			Price:       79.99,
			Description: "FIDO2 security key set for two-factor authentication",
		},
		{
			Name:        "Privacy Screen",
			Category:    "accessories",
			Price:       29.99,
			Description: "Anti-spy laptop privacy screen filter",
		},
		{
			Name:        "VPN License",
			Category:    "software",
			Price:       59.99,
			Description: "Annual license for premium VPN service",
		},
		{
			Name:        "Hardware Firewall",
			Category:    "hardware",
			Price:       199.99,
			Description: "Open-source network security appliance",
		},
	}
}