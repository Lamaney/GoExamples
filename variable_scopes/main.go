package main

import "fmt"

var x = 10

func main() {
	var condition = true

	if condition {
		//block scope
		var x = 10
		fmt.Println(x)
	}

	//function scope
	fmt.Println(condition)

	//globla scope
	fmt.Println(x)
}
