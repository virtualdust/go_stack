package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
    	s.Push(233, 666, 1024)
	s.Print()
	for !s.Empty() {
		s.Pop()
	}
	s.Pop()
	s.Print()
}
