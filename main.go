package main

import (
	"fmt"

	"github.com/Sohini-2026/DSA/recursive"
)

func main() {
	fmt.Println("What did u return ?", recursive.Print1ToN(5))
	recursive.Print1ToN(5)
	fmt.Println("Fibonacci of 5 is ", recursive.Fibonacci(5))
	fmt.Println("Height of BT is ", recursive.FindHeightOfBT(&recursive.Node{Data: 1, Left: &recursive.Node{Data: 2}, Right: &recursive.Node{Data: 3}}))
	fmt.Println("Sorting array using recursion")
	//arr := []int{5, 2, 9, 1, 5, 6}
	arr := []int{1, 5, 2, 0}
	recursive.SortArray(&arr)
	fmt.Println("Sorted array is ", arr)
	fmt.Println("Sorting stack using recursion")
	stack := &recursive.Stack{Items: []int{1, 5, 2, 0}}
	recursive.SortStack(stack)
	fmt.Println("Sorted stack is ", stack.Items)
}
