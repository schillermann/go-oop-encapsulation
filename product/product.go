package shop

import (
	"fmt"
)

type Product struct {
	id    int
	Name  string
	Price float32
}

func (p Product) Print() {
	fmt.Printf("Id: %d, Name: %s, Price: %.2f", p.id, p.Name, p.Price)
}
