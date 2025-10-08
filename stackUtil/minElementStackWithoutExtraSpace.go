package stackUtil

/*
Design a Data Structure SpecialStack that supports all the stack operations like push(), pop(), isEmpty(), isFull() and an additional operation getMin() which should return minimum element from the SpecialStack. All these operations of SpecialStack must be O(1). To implement SpecialStack, you should only use standard Stack data structure and no other data structure like arrays, list, .. etc.
*/

import (
	"github.com/Sohini-2026/DSA/baseDS"
)

type MinimumElementStack struct {
	mainStack  *baseDS.Stack
	minElement int
}

func NewMinimumElementStack() *MinimumElementStack {
	return &MinimumElementStack{
		mainStack:  &baseDS.Stack{},
		minElement: -1, // or some sentinel value
	}
}

func (s *MinimumElementStack) Push(value int) {
	if s.mainStack.IsEmpty() {
		s.mainStack.Push(value)
		s.minElement = value
	} else if value >= s.minElement {
		s.mainStack.Push(value)
	} else {
		encodedValue := 2*value - s.minElement
		s.mainStack.Push(encodedValue)
		s.minElement = value
	}
}

func (s *MinimumElementStack) Pop() int {
	if s.mainStack.IsEmpty() {
		return -1 // or some error value
	}
	topValue := s.mainStack.Pop()
	if topValue >= s.minElement {
		return topValue
	} else {
		originalMin := s.minElement
		s.minElement = 2*s.minElement - topValue
		return originalMin
	}
}

func (s *MinimumElementStack) GetMin() int {
	if s.mainStack.IsEmpty() {
		return -1 // or some error value
	}
	return s.minElement
}
