package recursive

import (
	"fmt"

	"github.com/Sohini-2026/DSA/baseDS"
)

func ReverseStack(s *baseDS.Stack) {
	if s.IsEmpty() {
		return
	}

	top := s.Pop()
	fmt.Println("Top element is ", top, " and stack is ", s.Items)
	ReverseStack(s)

	insertAtBottom(s, top)
}
func insertAtBottom(s *baseDS.Stack, item int) {
	if s.IsEmpty() {
		s.Push(item)
		return
	}

	temp := s.Pop()
	fmt.Println("Popped element is ", temp, " and stack is ", s.Items)
	insertAtBottom(s, item)
	s.Push(temp)
	fmt.Println("Pushed element is ", temp, " and stack is ", s.Items)
}
