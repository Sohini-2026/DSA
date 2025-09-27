package heapUtil

/*
Given an unsorted array of numbers, find the top 'K' frequently occurring numbers in it.
Example:
Input :  arr[] = {5, 2, 1,1,1, 3, 2}
k = 2
Output : {1,2}
*/
import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Sohini-2026/DSA/baseDS"
)

func TopKFrequentNumbers(arr []int, k int) ([]int, error) {
	if k > len(arr) {
		return nil, errors.New("k is larger than array size")
	}

	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
	}

	h := &baseDS.PairMinHeap{}
	heap.Init(h)

	for num, freq := range freqMap {
		heap.Push(h, baseDS.Pair{Variable: freq, Val: num})
		if h.Len() > k {
			heap.Pop(h)
		}
		fmt.Printf("Current min-heap: %+v\n", *h)
	}

	result := []int{}
	for h.Len() > 0 {
		elem := heap.Pop(h).(baseDS.Pair)
		result = append(result, elem.Val)
	}

	return result, nil
}
