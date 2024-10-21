package main

import (
	"fmt"
)

func main() {

	/*var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang"
	quantity = 10
	discount = 0.38
	isInStock = false

	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(isInStock, reflect.TypeOf(isInStock))*/
	/*var productName string = "Golang"
	fmt.Println(productName, reflect.TypeOf(productName));*/

	/*productName := "Golang"

	fmt.Println(productName)*/

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang"
	quantity = 10
	discount = 0.38
	isInStock = false

	fmt.Println("product name:", productName, "Quantity:", quantity, "Discount:", discount, "IsInStock:", isInStock)
	fmt.Printf("Product Name: %s, Quantity: %d, Discount %f, IsInStock %t", productName, quantity, discount, isInStock)

}
