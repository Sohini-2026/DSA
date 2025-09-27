package heapUtil

/*Given a list of points on the 2-D plane and an integer K. The task is to find K closest points to the origin and print them.

Note: The distance between two points on a plane is the Euclidean distance.

Example:
Input : point = [[3, 3], [5, -1], [-2, 4]], K = 2 .
Output : [[3,3],[-2,4]]
Explanation: The Euclidean distance between (x1, y1) and (x2, y2) is sqrt((x1 - x2)^2 + (y1 - y2)^2). So the distance between each point and the origin is:
Distance between (3, 3) and the origin = sqrt(3^2 + 3^2) = sqrt(18)
Distance between (5, -1) and the origin = sqrt(5^2 + -1^2) = sqrt(26)
Distance between (-2, 4) and the origin = sqrt((-2)^2 + 4^2) = sqrt(20)
Thus the closest 2 points to the origin are (3, 3) and (-2, 4)
*/

import (
	"container/heap"
	"errors"
	"fmt"
	"math"

	"github.com/Sohini-2026/DSA/baseDS"
)

func KClosestPointsToOrigin(points [][]int, k int) ([][]int, error) {
	if k > len(points) {
		return nil, errors.New("k is larger than array size")
	}

	h := &baseDS.PairMaxHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		dist := int(math.Sqrt(float64(points[i][0]*points[i][0] + points[i][1]*points[i][1])))
		heap.Push(h, baseDS.Pair{Variable: dist, Val: i})
	}

	fmt.Printf("Initial max-heap: %+v\n", *h) //[10 7 4]

	for i := k; i < len(points); i++ {
		dist := int(math.Sqrt(float64(points[i][0]*points[i][0] + points[i][1]*points[i][1])))
		if dist < (*h)[0].Variable {
			heap.Pop(h)
			heap.Push(h, baseDS.Pair{Variable: dist, Val: i})
			fmt.Printf("Updated max-heap: %+v\n", *h)
		}
	}

	result := [][]int{}
	for h.Len() > 0 {
		elem := heap.Pop(h).(baseDS.Pair)
		result = append(result, points[elem.Val])
	}

	return result, nil
}
