package main

import (
	"fmt"

	"https://github.com/Sohini-2026/DSA/recursive"
)

func main() {
	fmt.Println("What did u return ?", recursive.Print1ToN(5))
	recursive.Print1ToN(5)
	fmt.Println("Fibonacci of 5 is ", recursive.Fibonacci(5))
	fmt.Println("Height of BT is ", recursive.FindHeightOfBT(&recursive.Node{Data: 1, Left: &recursive.Node{Data: 2}, Right: &recursive.Node{Data: 3}}))
}
