package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Structs")

	// initialization
	var p = Person{name: "Baraka", age: 20}
	fmt.Println(p)
	// we use the dot notation to acces a structs fields
	fmt.Println(p.name)

}
