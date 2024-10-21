package main

import "fmt"

func main() {
	/*total := add(10, 20)
	fmt.Println(total)*/
	//total, diffrence := calculation(30, 20)
	//fmt.Println(total, diffrence)
	//var numbers = []int{10, 20, 2, 3, 5}
	//fmt.Println(sumArray(numbers))

	fmt.Println(sum(3, 5, 6, 22, 4, 5, 6, 78, 3, 2, 4))

}

//like c# params
func sum(numbers ...int) int {

	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func sumArray(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func add(x int, y int) int {
	return x + y
}

//like c# tupples
func calculation(x int, y int) (int, int) {
	return x + y, x - y
}
