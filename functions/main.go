package main

import "fmt"

func main() {
	// this is the main function
	fmt.Println("Fuctions")
	// call an external function
	sum := Adder(1, 3)
	fmt.Println(sum)
	// call a function of two arguments
	a, b := GetDescription("baraka")
	fmt.Println(a, b)
	// exluding a value returen using an underscore
	_, c := GetDescription("baraka")
	fmt.Println(c)
	// or
	d, _ := GetDescription("baraka")
	fmt.Println(d)
}

//  functions with a return type explicitly typed
// a. adding two numbers
func Adder(a int, b int) int {
	return a + b
}

// multiple return types
func GetDescription(name string) (string, int) {
	a := "Goodmorning " + name + " the time is"
	b := 8
	return a, b
}
