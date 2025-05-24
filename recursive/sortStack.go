package recursive

import "github.com/Sohini-2026/DSA/baseDS"

func SortStack(s *baseDS.Stack) {
	if s.IsEmpty() {
		return
	}

	top := s.Pop()

	SortStack(s)

	insertSorted(s, top)
}

func insertSorted(s *baseDS.Stack, item int) {
	if s.IsEmpty() || item < s.Top() {
		s.Push(item)
		return
	}

	temp := s.Pop()
	insertSorted(s, item)
	s.Push(temp)
}
