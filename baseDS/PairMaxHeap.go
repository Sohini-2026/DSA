package baseDS

type Pair struct {
	Variable int
	Val      int
}

// PairMaxHeap is a max-heap of Pairs based on the Diff field.
type PairMaxHeap []Pair

func (h PairMaxHeap) Len() int           { return len(h) }
func (h PairMaxHeap) Less(i, j int) bool { return h[i].Variable > h[j].Variable }
func (h PairMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PairMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *PairMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *PairMaxHeap) Peek() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}
