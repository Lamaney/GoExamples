package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

func (b *Book) ShippingCost() int {
	return b.desi + 5*2
}

func (e *Electronic) ShippingCost() int {
	return e.desi + 10*2
}

func main() {

	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Electronic{desi: 20},
		&Book{desi: 30},
	}

	fmt.Println(CalculateShippingCost(products))
}

func CalculateShippingCost(products []IShippable) int {
	cost := 0
	for _, product := range products {
		cost += product.ShippingCost()
	}
	return cost
}
