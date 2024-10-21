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
	test()
	fmt.Println(x)
}

func test() {
	//function scope
	//if have function scope variable with same name as global variable compiler uses function scope variable
	var x = 20
	fmt.Println(x)
}
