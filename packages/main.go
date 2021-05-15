package main

import (
	"fmt"
	"go-examples/packages/strutil"
	"math"

	"github.com/lessons/packages/strutil"
)

func main() {
	fmt.Println("Packages")
	// import the inbuilt package
	val := math.Sqrt(45)
	fmt.Println(val)
	// import custom built packages
	// a package to reverse a string
	fmt.Println(strutil.Reverse("HELLO THERE"))
}
