package main

import "fmt"

// Package level variable
// You can't use short declaration operator here, it must be declared with "var"
var test = 10

func main() {
	fmt.Printf("test: %v %T\n", test, test)

	// Short declaration operator, auto type inference
	// And these are function/block level variables
	x := 10
	y := "Good morning"

	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("y: %v %T\n", y, y)

	// Assignment operator
	x = 20 // Reassigning value
	fmt.Printf("x: %v %T\n", x, x)

	// x := 30 // error - x is already declared

	// Multiple assignment - Even though x is already declared, it can be reassigned like this
	x, z := 30, 40
	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("z: %v %T\n", z, z)

	w := 20 + 20
	fmt.Printf("w: %v %T\n", w, w)

	k := 10 == 10
	fmt.Printf("k: %v %T\n", k, k)
	k = 10 > 20
	fmt.Printf("k: %v %T\n", k, k)
}
