package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	// declaration
	var a [5]int
	a[0] = 45
	// printing an array at an index
	fmt.Println(a[0])

	// iterating an array
	numbers := [5]int{12, 45, 78, 65, 32}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d ", numbers[i])

	}
	// sort using conditioning
	fmt.Println()

	fmt.Println("Sorted above 40")

	for i := 0; i < len(numbers); i++ {

		if numbers[i] > 40 {
			fmt.Printf("%d ", numbers[i])

		}

	}
	fmt.Println()

}
