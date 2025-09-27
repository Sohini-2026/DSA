package heapUtil

/*
Given an unsorted array and two numbers x and k, find k closest values to x.
Input : arr[] = {10, 2, 14, 4, 7, 6}, x = 5, k = 3 .
*/

import (
	"container/heap"
	"errors"
	"math"

	"github.com/Sohini-2026/DSA/baseDS"
)

func KClosestElements(arr []int, x int, k int) ([]int, error) {
	if k > len(arr) {
		return nil, errors.New("k is larger than array size")
	}

	h := &baseDS.PairMaxHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		diff := int(math.Abs(float64(arr[i] - x)))
		heap.Push(h, baseDS.Pair{Variable: diff, Val: arr[i]})
	}

	for i := k; i < len(arr); i++ {
		diff := int(math.Abs(float64(arr[i] - x)))
		if diff < (*h)[0].Variable {
			heap.Pop(h)
			heap.Push(h, baseDS.Pair{Variable: diff, Val: arr[i]})
		}
	}

	result := []int{}
	for h.Len() > 0 {
		elem := heap.Pop(h).(baseDS.Pair)
		result = append(result, elem.Val)
	}

	return result, nil
}
