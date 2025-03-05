package main

// https://leetcode.com/problems/implement-queue-using-stacks/description/
type Stack struct {
	stack []int
	size  int
}

func StackConstructor() Stack {
	return Stack{make([]int, 0), 0}
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		return 0
	}
	l := len(s.stack)
	last := s.stack[l-1]
	s.stack = s.stack[:l-1]
	s.size--
	return last
}

func (s *Stack) Peek() int {
	if s.size == 0 {
		return 0
	}
	return s.stack[len(s.stack)-1]

}

func (s *Stack) Push(x int) {
	s.stack = append(s.stack, x)
	s.size++
}

type MyQueue struct {
	pushStack Stack
	popStack  Stack
}

func Constructor11() MyQueue {
	pushStack, popStack := Stack{make([]int, 0), 0}, Stack{make([]int, 0), 0}
	return MyQueue{pushStack, popStack}
}

func (this *MyQueue) Push(x int) {
	this.pushStack.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.popStack.size == 0 {
		for this.pushStack.size != 0 {
			p := this.pushStack.Pop()
			this.popStack.Push(p)
		}
	}
	return this.popStack.Pop()
}

func (this *MyQueue) Peek() int {
	if this.popStack.size == 0 {
		for this.pushStack.size != 0 {
			p := this.pushStack.Pop()
			this.popStack.Push(p)
		}
	}
	return this.popStack.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.popStack.size == 0 && this.pushStack.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
