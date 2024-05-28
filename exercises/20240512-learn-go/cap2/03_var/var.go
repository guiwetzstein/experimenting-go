package main

import "fmt"

var y = 10 // Package level scope

func main() {
	whatever(y)
}

func whatever(x int) {
	fmt.Println(x)
	fmt.Println(y)
}
