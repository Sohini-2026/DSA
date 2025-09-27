package baseDS

// PairMinHeap is a min-heap of Pairs based on the Diff field.
type PairMinHeap []Pair

func (h PairMinHeap) Len() int           { return len(h) }
func (h PairMinHeap) Less(i, j int) bool { return h[i].Variable < h[j].Variable }
func (h PairMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PairMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *PairMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *PairMinHeap) Peek() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}
