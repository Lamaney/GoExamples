package main

import "fmt"

func main() {
	//stack var lifo var
	defer fmt.Println("1.defer")
	defer fmt.Println("2.defer")
	defer fmt.Println("3.defer")

	x := 10
	y := 20
	defer fmt.Println(x + y)

	fmt.Println("Main started")

	x = 100
	y = 200

	fmt.Println(x + y)

	var condition = true
	defer cleanUp()

	if condition {
		panic("Something went wrong")
	}

	defer fmt.Println("4.defer")
	if condition {
		return
	}
	defer fmt.Println("5.defer")
}

func cleanUp() {
	fmt.Println("Cleaning up")
}
