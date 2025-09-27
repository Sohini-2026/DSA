package baseDS

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

func (s *Stack) Size() int {
	return len(s.Items)
}
