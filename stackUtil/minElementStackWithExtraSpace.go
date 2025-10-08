package stackUtil

/*
Design a Data Structure SpecialStack that supports all the stack operations like push(), pop(), isEmpty(), isFull() and an additional operation getMin() which should return minimum element from the SpecialStack. All these operations of SpecialStack must be O(1). To implement SpecialStack, you should only use standard Stack data structure and no other data structure like arrays, list, .. etc.
*/

import (
	"github.com/Sohini-2026/DSA/baseDS"
)

type MinElementStack struct {
	mainStack *baseDS.Stack
	minStack  *baseDS.Stack
}

func NewMinElementStack() *MinElementStack {
	return &MinElementStack{
		mainStack: &baseDS.Stack{},
		minStack:  &baseDS.Stack{},
	}
}

func (s *MinElementStack) Push(value int) {
	s.mainStack.Push(value)
	if s.minStack.IsEmpty() || value <= s.minStack.Top() {
		s.minStack.Push(value)
	}
}

func (s *MinElementStack) Pop() int {
	if s.mainStack.IsEmpty() {
		return -1 // or some error value
	}
	poppedValue := s.mainStack.Pop()
	if poppedValue == s.minStack.Top() {
		s.minStack.Pop()
	}
	return poppedValue
}

func (s *MinElementStack) GetMin() int {
	if s.minStack.IsEmpty() {
		return -1 // or some error value
	}
	return s.minStack.Top()
}
