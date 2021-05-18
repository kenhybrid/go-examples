package main

import "fmt"

func main() {
	fmt.Println("Loops")
	// while like for loop
	a := 1
	for a < 10 {
		fmt.Println(a)
		a++
	}

	// for	normal loop
	for i := 0; i < 10; i++ {
		fmt.Println("The number is at", i)
	}
	// fizzbuzz
	for b := 0; b < 50; b++ {
		if b%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if b%3 == 0 {
			fmt.Println("buzz")
		} else if b%5 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(b)
		}
	}
	// pattern
	for c := 0; c <= 10; c++ {
		fmt.Printf("\n")
		for f := c; f <= 10; f++ {
			fmt.Printf("%d", f)
		}

		for f := 0; f <= c; f++ {
			fmt.Printf("%d", f)
		}

	}
}
