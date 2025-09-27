package heapUtil

/*
Given an array of integers and two numbers k1 and k2. Find the sum of all elements between given two k1’th and k2’th smallest elements of the array. It may  be assumed that all elements of array are distinct.

Example :
Input : arr[] = {20, 8, 22, 4, 12, 10, 14},  k1 = 3,  k2 = 6
Output : 26
         3rd smallest element is 10. 6th smallest element
         is 20. Sum of all element between k1 & k2 is
         12 + 14 = 26 .
*/
import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Sohini-2026/DSA/baseDS"
)

func SumOfElementsBetweenKthSmallestElements(arr []int, k1 int, k2 int) (int, error) {
	if k1 > len(arr) || k2 > len(arr) || k1 >= k2 {
		return -1, errors.New("invalid k1 or k2 values")
	}

	h := &baseDS.IntMinHeap{}
	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
	}

	fmt.Printf("Initial min-heap: %+v\n", *h) //[2 3 5 6]

	sum := 0
	for i := 1; i < k2; i++ {
		minElem := heap.Pop(h).(int)
		if i > k1 {
			fmt.Println("Adding:", minElem)
			sum += minElem
		}
	}

	return sum, nil
}
