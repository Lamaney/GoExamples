package main

import "fmt"

func main() {

	// fixed array. sabit boyutlu
	var names = []string{"Omer", "Faruk", "Aydin"}

	names = append(names, "Faruk")

	fmt.Println(names)

}
