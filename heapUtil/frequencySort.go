package heapUtil

/*
Print the elements of an array in the increasing frequency if 2 numbers have same frequency then print the one which came first.

Example:
Input : arr[] = {2, 5, 2, 8, 5, 6, 8, 8}
Output : arr[] = {8, 8, 8, 2, 2, 5, 5, 6} . */

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Sohini-2026/DSA/baseDS"
)

func FrequencySort(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("array is empty")
	}

	freqMap := make(map[int]int)

	for _, num := range arr {
		freqMap[num]++
	}

	h := &baseDS.PairMaxHeap{}
	heap.Init(h)

	for num, freq := range freqMap {
		heap.Push(h, baseDS.Pair{Variable: freq, Val: num})
	}

	fmt.Printf("Initial min-heap: %+v\n", *h)

	result := []int{}
	for h.Len() > 0 {
		elem := heap.Pop(h).(baseDS.Pair)
		for i := 0; i < elem.Variable; i++ {
			result = append(result, elem.Val)
		}
	}

	return result, nil
}
