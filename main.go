package main

import (
	"fmt"

	"github.com/Sohini-2026/DSA/recursive"
)

func main() {
	/*
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
		stack := &baseDS.Stack{Items: []int{1, 5, 2, 0}}
		recursive.SortStack(stack)
		fmt.Println("Sorted stack is ", stack.Items)
		fmt.Println("Deleting middle element of stack using recursion")
		deleteStack := &baseDS.Stack{Items: []int{1, 2, 3, 4, 5}}
		recursive.DeleteMiddleElementOfStack(deleteStack)
		fmt.Println("Stack after deleting middle element is ", deleteStack.Items)
		fmt.Println("Reversing stack using recursion")
		reverseStack := &baseDS.Stack{Items: []int{1, 2, 3, 4, 5}}
		recursive.ReverseStack(reverseStack)
		fmt.Println("Reversed stack is ", reverseStack.Items)

		fmt.Println("Tower of Hanoi with 3 disks")
		recursive.TowerOfHanoi(3, "A", "C", "B")
		fmt.Println("Kth symbol in grammar for n=4 and k=3 is ", recursive.GetKthSymbolInGrammer(4, 1))

		fmt.Printf("Print all subsets of array %v\n", []string{"a", "b", "c"})
		recursive.PrintAllSubsets1("abc", "")

		m := make(map[string]bool)
		result := []string{}
		recursive.PrintUniqueSubsetsLexicographially("abb", "", m, &result)
		sort.Strings(result)
		fmt.Println(result)

	*/

	fmt.Printf("Permutation With Spaces %v\n", "abc")
	recursive.PermutationWithSpaces("a", "bc")

	fmt.Printf("Permutation With Case Change %v\n", "abc")
	recursive.PermutationWithCaseChange(" ", "abc")

}
