// Package main provides ...
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Slices")
	// slice initialization
	ages := []int{4, 5, 6, 2, 4, 3, 7, 5, 8, 6, 8}
	fmt.Println(ages)

	// SUBSLICING
	// ages[lowerbound:upper bound]
	fmt.Println(ages[2:5])
	//missing uper or lower bound
	// assumes as from 0
	fmt.Println(ages[:5])
	// assumes the last value
	fmt.Println(ages[5:])
	//one can append or coppy from a slice
	// append
	number1 := []int{1, 5, 4, 3, 6}
	// number2 := []int{85, 45, 6}

	numbers := append(number1, 5, 6, 8)
	fmt.Println("before append ", number1)

	fmt.Println("after append", numbers)
	// copy
	// copied := copy(number1, number2)
	// fmt.Println("copied", copied)
	/* create a slice numbers1 with double the capacity of earlier slice*/
	b := make([]int, len(numbers), (cap(numbers))*2)

	/* copy content of numbers to numbers1 */
	a := copy(b, numbers)
	fmt.Println(a)
	ran := rand.Intn(100) + 5
	fmt.Println(ran)
}
