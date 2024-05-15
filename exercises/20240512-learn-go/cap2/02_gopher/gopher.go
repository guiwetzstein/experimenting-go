package main

import "fmt"

func main() {
	// Short declaration operator, auto type inference
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
}
