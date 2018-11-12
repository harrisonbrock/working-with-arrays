package main

import "fmt"

func main() {

	// create slice
	numbers := []int{8, 1, 10, 5, 3, 2, 6, 7, 4, 9}

	// get the length of a slice
	length := len(numbers)

	// print the length of the slice
	fmt.Println("The length is:", length)

	// iterating over an slice
	print(numbers)

	// get sum of numbers
	sum := sum(numbers)
	fmt.Println("Sum of numbers:", sum)

	// get average
	avg := average(numbers)
	fmt.Println("Average of numbers:", avg)

}

func print(arr []int) {
	for index, number := range arr {
		fmt.Printf("Index: %v and Number: %v\n", index, number)

	}
}

func sum(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum
}

func average(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum / len(arr)

}
