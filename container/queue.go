package container

type Queue struct {
	l *List
}

func NewQueue() *Queue {
	l := NewList()
	return &Queue{l}
}

func (q *Queue) Push(ele Element) {
	q.l.Append(ele)
}

func (q *Queue) Pop() Element {
	return q.l.Pop()
}

func (q *Queue) Empty() bool {
	return q.l.Empty()
}
