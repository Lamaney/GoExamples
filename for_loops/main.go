package main

import "fmt"

func main() {

	/*index := 1

	for index <= 10 {
		fmt.Println(index)
		index++
	}*/
	/*total := 0
	for index := 1; index <= 10; index++ {
		total += index
		fmt.Println(total)
	}*/

	/*index := 0
	var names = [3]string{"Omer", "Faruk", "Aydin"}

	for index < len(names) {
		fmt.Print(names[index])
		index++
	}*/

	/*for index := 0; index <= 10; index++ {
		if index == 3 || index == 5 {
			continue
		}
		fmt.Println(index)
	}*/

	var names = map[string]int{
		"Omer":  10,
		"Faruk": 20,
		"Aydin": 30,
	}

	for key, value := range names {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
