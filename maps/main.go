package main

import "fmt"

func main() {
	/*var names map[string]int

	names = make(map[string]int, 0)

	names["Omer"] = 1
	names["Faruk"] = 2
	names["Aydin"] = 3

	fmt.Println(names)*/

	names := map[string]int{
		"Omer":  1,
		"Faruk": 2,
		"Aydin": 3,
	}

	fmt.Println(names)

	delete(names, "Recep")

	fmt.Println(names)
}
