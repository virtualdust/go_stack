package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
    	s.Push(233, 666, 1024)
	PrintStackInfo(s)
	for !s.Empty() {
		s.Pop()
	}
	s.Pop()
	PrintStackInfo(s)
}

func PrintStackInfo(s *Stack){
	fmt.Println("stack len is", s.Len())
    	fmt.Println("stack empty is", s.Empty())
	fmt.Println("stack top is", s.Top())
	s.Print()
}
