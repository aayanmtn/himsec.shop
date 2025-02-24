
package models

type Wish struct {
	ProductName string
	Status     string
}

func NewWish(productName string) Wish {
	return Wish{
		ProductName: productName,
		Status:     "pending",
	}
}
