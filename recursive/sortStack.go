package recursive

type Stack struct {
	Items []int
}

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}
func (s *Stack) Pop() int {
	if len(s.Items) == 0 {
		return -1 // or some error value
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}
func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) Top() int {
	if len(s.Items) == 0 {
		return -1 // or some error value
	}
	return s.Items[len(s.Items)-1]
}

func SortStack(s *Stack) {
	if s.IsEmpty() {
		return
	}

	top := s.Pop()

	SortStack(s)

	insertSorted(s, top)
}

func insertSorted(s *Stack, item int) {
	if s.IsEmpty() || item < s.Top() {
		s.Push(item)
		return
	}

	temp := s.Pop()
	insertSorted(s, item)
	s.Push(temp)
}
