package main

import (
	"fmt"
	"sort"

	"github.com/Sohini-2026/DSA/baseDS"
	"github.com/Sohini-2026/DSA/binarySearch"
	"github.com/Sohini-2026/DSA/dP"
	"github.com/Sohini-2026/DSA/recursive"
)

func main() {

	//recursion : start
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

	fmt.Printf("Permutation With Spaces %v\n", "abc")
	recursive.PermutationWithSpaces("a", "bc")

	fmt.Printf("Permutation With Case Change %v\n", "abc")
	recursive.PermutationWithCaseChange(" ", "abc")
	fmt.Printf("Letter Case Permutation %v\n", "a1B2")
	recursive.LetterCasePermutation("", "a1B2")

	fmt.Printf("Balanced Parenthesis for n=3\n")
	recursive.CreateBalancedParanthesis("", "", 3, 3)
	fmt.Printf("Is Balanced Parenthesis for string %s\n", "({[]})")
	fmt.Println(recursive.IsBalancedParanthesis("({[]})"))

	fmt.Printf("Get N bit Binary Number for n=3\n")
	recursive.NBitBinaryNumber(5, "", 0, 0)

	fmt.Printf("Get last index in josephus problem for n=7 and k=3\n")
	fmt.Println(recursive.JosepheusProblem(7, 3))

	// recursion : end

	//binary search : start
	fmt.Printf("Binary Search for element 5 in array %v\n", []int{1, 2, 3, 4, 5, 6, 7})
	res := binarySearch.SearchNumberInSortedArray([]int{1, 2, 3, 4, 5, 6, 7}, 5)
	fmt.Println("Element found at index:", res)
	res = binarySearch.SearchNumberInReverseSortedArray([]int{7, 6, 5, 3, 2, 1}, 1)
	fmt.Println("Element found at index:", res)

	res = binarySearch.SearchNumberWithUnknownSorting([]int{1, 2, 3}, 1)
	fmt.Println("Element found at index:", res)

	count := binarySearch.FindCountOfOccurrences([]int{1, 2, 3, 4, 5, 5, 5, 6}, 5)
	fmt.Println("Count of occurrences of 5 is:", count)

	res = binarySearch.CountNoOfTimesArrayIsRotated([]int{6, 7, 1, 2, 3, 4, 5})
	fmt.Println("Count of times array is rotated is:", res)
	res = binarySearch.CountNoOfTimesArrayIsRotated([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("For non rotated array - Count of times array is rotated is:", res)

	//res = binarySearch.CountNoOfTimesArrayIsRotated([]int{7, 6, 5, 4, 3, 2, 1})
	//fmt.Println("Completely rotated - Count of times array is rotated is:", res)

	res = binarySearch.FindElementInRotatedSortedArray([]int{6, 7, 1, 2, 3, 4, 5}, 3)
	fmt.Println("Element 3 found at index:", res)

	res = binarySearch.FindElementInNearlySortedArray([]int{5, 10, 30, 20, 40}, 20)
	fmt.Println("Element 5 found at index:", res)

	res = binarySearch.FindFloorOfElementInSortedArray([]int{1, 2, 3, 4, 8, 10}, 5)
	fmt.Println("Floor index::", res)

	res = binarySearch.FindCeilOfElementInSortedArray([]int{1, 2, 3, 4, 8, 10}, 5)
	fmt.Println("Ceil index::", res)

	ress := binarySearch.FindNextAlphabeticalElement([]string{"a", "b", "c", "d"}, "c")
	fmt.Println("Next alphabetical element after 'c' is:", ress)

	res = binarySearch.FindPositionOfElementInInfiniteSortedArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20}, 5)
	fmt.Println("Position of element 5 in infinite sorted array is:", res)

	res = binarySearch.FindIndexOf1InBSInfiniteArray([]int{0, 0, 0, 0, 0, 1, 1, 1, 1})
	fmt.Println("Index of first 1 in infinite binary sorted array is:", res)

	res = binarySearch.FindMinDiffElementInSortedArray([]int{1, 2, 3, 4, 7, 8, 9}, 5)
	fmt.Println("Min diff element in sorted array is found", res)

	resArr := binarySearch.FindPeakElement([]int{5, 10, 20, 90, 80})
	fmt.Printf("Peak element in array is found at index::%+v\n", resArr)

	res = binarySearch.FindElementInBitonicArray([]int{1, 3, 8, 12, 4, 2}, 4)
	fmt.Println("Element 4 found in bitonic array at index:", res)

	row, col := binarySearch.FindElementInRowColumnSortedMatrix([][]int{
		{10, 20, 30, 40},
		{15, 25, 35, 45},
		{27, 29, 37, 48},
		{32, 33, 39, 50},
	}, 29)
	fmt.Println("Element 29 found in row-column sorted matrix", row, col)

	res = binarySearch.AllocateMinNoOfPages([]int{10, 20, 30, 40}, 2)
	fmt.Println("Minimum number of pages allocated is:", res)
	//binary search : end

	// dynamic programming on trees : start
	root := &baseDS.TreeNode{Value: 1,
		Left: &baseDS.TreeNode{Value: 2,
			Left:  &baseDS.TreeNode{Value: 4},
			Right: &baseDS.TreeNode{Value: 5}},
		Right: &baseDS.TreeNode{Value: 3,
			Left:  &baseDS.TreeNode{Value: 6},
			Right: &baseDS.TreeNode{Value: 7}}}
	resDiameter := 0
	dP.DiameterOfBinaryTree(root, &resDiameter)
	fmt.Println("Diameter of binary tree is:", resDiameter)
}
