package main

import (
	shop "schillermann/go-oop-encapsulation/product"
)

func main() {

	product := shop.Product{
		Name:  "Honey",
		Price: 10.23,
	}

	product.Print()
}
