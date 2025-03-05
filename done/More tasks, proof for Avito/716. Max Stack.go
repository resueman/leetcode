package main

type MaxStack struct {
	stack    []int
	maxStack []int
}

func ConstructorMaxStack() *MaxStack {
	return &MaxStack{make([]int, 0), make([]int, 0)}
}

func (s *MaxStack) Push(number int) {
	s.stack = append(s.stack, number)
	if len(s.maxStack) == 0 || number >= s.maxStack[len(s.maxStack)-1] {
		s.maxStack = append(s.maxStack, number)
	}
}

func (s *MaxStack) Pop() int {
	if len(s.stack) == 0 {
		return -1
	}

	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if s.maxStack[len(s.maxStack)-1] == v {
		s.maxStack = s.maxStack[:len(s.maxStack)-1]
	}
	return v
}

func (s *MaxStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MaxStack) PeekMax() int {
	return s.maxStack[len(s.maxStack)-1]
}

func (s *MaxStack) PopMax() int {
	if len(s.stack) == 0 {
		return -1
	}

	buffer := make([]int, 0, len(s.stack))
	max := s.maxStack[len(s.maxStack)-1]
	for s.stack[len(s.stack)-1] != max {
		v := s.Pop()
		buffer = append(buffer, v)
	}

	s.maxStack = s.maxStack[:len(s.maxStack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	for _, v := range buffer {
		s.Push(v)
	}

	return max
}
