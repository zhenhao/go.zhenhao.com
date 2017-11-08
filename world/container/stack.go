package container

type Stack struct {
	l *List
}

func NewStack() *Stack {
	l := NewList()
	return &Stack{l}
}

func (s *Stack) Push(ele Element) {
	s.l.Push(ele)
}

func (s *Stack) Pop() Element {
	return s.l.Pop()
}

func (s *Stack) Empty() bool {
	return s.l.Empty()
}
