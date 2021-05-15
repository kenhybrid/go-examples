package main

import (
	"fmt"
)

func main() {

	fmt.Println("conditionals")
	// if statements
	// simple
	err := true
	// err = true
	if err {
		fmt.Println("an error ocured")
	}
	// if/else
	age := 4
	age2 := 3
	if age > age2 {
		fmt.Printf("%d is greater than %d\n", age, age2)

	} else {
		fmt.Printf("%d is greater than %d\n", age2, age)
	}
	// if/else if
	color := "red"
	if color == "red" {
		fmt.Println("RED")
	} else if color == "blue" {
		fmt.Println("BLUE")
	} else {
		fmt.Println("ITS NEITHER BLUE OR RED")
	}
	// switch
	switch color {
	case "red":
		fmt.Println("RED")
	case "blue":
		fmt.Println("BLUE")
	case "green":
		fmt.Println("GREEN")

	}

}
