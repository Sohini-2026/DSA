package recursive

import "github.com/Sohini-2026/DSA/baseDS"

// DeleteMiddleElementOfStack deletes the middle element of a stack using recursion.
func DeleteMiddleElementOfStack(s *baseDS.Stack) {
	if s.IsEmpty() {
		return
	}

	midIndex := (len(s.Items) / 2) + 1
	deleteMiddleHelper(s, midIndex)
}
func deleteMiddleHelper(s *baseDS.Stack, midIndex int) {
	if midIndex == 1 {
		s.Pop() // Remove the middle element
		return
	}

	top := s.Pop()                    // Pop the top element
	deleteMiddleHelper(s, midIndex-1) // Recur for the next element

	s.Push(top) // Push the top element back onto the stack
}
