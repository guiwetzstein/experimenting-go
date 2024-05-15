package main

import "fmt"

func main() {
	nBytes, error := fmt.Println("Hello, World!", "Oi, Mundo!", 100)
	fmt.Println(nBytes, error)

	// Use _ (the blank identifier) on values you don't want to use
	_, error2 := fmt.Println("Hello, World!", "Oi, Mundo!", 100)
	fmt.Println(error2)

	// Primitive types
	x := 10
	y := "Hello"
	z := true

	fmt.Println(x, y, z)
}

// To create module: "go mod init hello-world"
// To run: "go run helloworld.go"
