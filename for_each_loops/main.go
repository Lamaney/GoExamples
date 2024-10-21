package main

import "fmt"

func main() {

	var numbers = []int{1, 2, 3, 4}

	/*for index := 0; index < len(numbers); index++ {
		fmt.Println(numbers[index])
	}*/

	for _, value := range numbers {
		fmt.Println(value)
	}

	var language = "Golang"

	for _, v := range language {
		fmt.Printf("%c\n", v)
	}
}
