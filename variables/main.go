package main

import "fmt"

func main() {
	// variable declaration methods
	var name = "Kennedy baraka"
	var age int
	age = 2

	// short hand
	course := "information technology"
	fmt.Println(name, age, course)
	// or
	fmt.Printf("%s %d %s\n", name, age, course)
	// data types in go
	max := 45.33
	fullname := "john terry"
	score := 45
	loged := false

	// printf for datatypes checkin
	fmt.Printf("%T\n", max)      //float64
	fmt.Printf("%T\n", fullname) //string
	fmt.Printf("%T\n", score)    //int
	fmt.Printf("%T\n", loged)    //bool

	// hybrid variable declaratoion
	a, b, c := 3, 3.34, "hey"
	fmt.Println(a, b, c)

	// constants dont change their value
	const yes = false
	// yes = true
	// cannot assign to yes its a constant
	fmt.Println(yes)
	// variables declared without initialization have a 0 value
	var minAge int
	fmt.Println(minAge)
}
