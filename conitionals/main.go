package main

import "fmt"

func main() {

	var age int8 = 17

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}
}