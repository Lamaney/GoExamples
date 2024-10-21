package main

import "fmt"

func main() {
	/*var a int
	a = 10

	var p *int

	fmt.Println(a)

	p = &a

	*p = 20

	fmt.Println(a)*/

	var a = 10
	var b int
	var p *int

	b = a
	p = &a
	*p = 20

	fmt.Println(a, b)

}
