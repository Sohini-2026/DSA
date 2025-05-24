package recursive

import "github.com/Sohini-2026/DSA/baseDS"

func ReverseStack(s *baseDS.Stack) {
	if s.IsEmpty() {
		return
	}

	top := s.Pop()

	ReverseStack(s)

	insertAtBottom(s, top)
}
func insertAtBottom(s *baseDS.Stack, item int) {
	if s.IsEmpty() {
		s.Push(item)
		return
	}

	temp := s.Pop()
	insertAtBottom(s, item)
	s.Push(temp)
}
