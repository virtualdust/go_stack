package stack

import (
	"fmt"
	"log"
)

type Stack []interface{}

const STACK_DEFAULT_CAP = 1024

func NewStack() *Stack {
	stack := make(Stack, 0, STACK_DEFAULT_CAP)
	return &stack
}

func (s *Stack) Push(e ...interface{}) {
	*s = append(*s, e...)
}

func (s *Stack) Pop() {
	length := len(*s)
	if length == 0 {
		log.Printf("Error:Try to pop an empty stack")
		return
	}
	*s = (*s)[:length-1]
}

func (s *Stack) Top() interface{} {
	length := len(*s)
	if length == 0 {
		return nil
	}
	return (*s)[length-1]
}

func (s *Stack) Empty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Print() {
	fmt.Printf("stack size is %d\n", s.Size())
    	fmt.Printf("stack empty is %v\n", s.Empty())
	fmt.Printf("stack top is %v\n", s.Top())
	fmt.Printf("stack is %s", func() string {
		var str string
		for i := len(*s) - 1; i >= 0; i-- {
			str = fmt.Sprintf("%s %v", str, (*s)[i])
		}
		return fmt.Sprintf("[%s ]\n", str)
	}())
}
