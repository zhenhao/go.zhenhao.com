package container

type Node struct {
	ele  Element
	prev *Node
	next *Node
}

func newNode(ele Element) *Node {
	return &Node{ele, nil, nil}
}

type List struct {
	head *Node
	tail *Node
	len  uint
}

func NewList() *List {
	return &List{nil, nil, 0}
}

func (l *List) Push(ele Element) {
	if l.Empty() {
		l.head = newNode(ele)
		l.tail = l.head
	} else {
		nn := newNode(ele)
		l.head.prev = nn
		nn.next = l.head
		l.head = nn
	}
	l.len++
}

func (l *List) Append(ele Element) {
	if l.Empty() {
		l.tail = newNode(ele)
		l.head = l.tail
	} else {
		nn := newNode(ele)
		nn.prev = l.tail
		l.tail.next = nn
		l.tail = nn
	}

	l.len++
}

func (l *List) Pop() Element {
	if l.Empty() {
		panic("out of range")
	}
	rn := l.head.ele
	l.head = l.head.next
	if nil != l.head {
		l.head.prev = nil
	}
	l.len--
	return rn
}

func (l *List) Get(idx uint) Element {
	if idx >= l.len {
		panic("out of range")
	}
	p := l.head
	for i := uint(0); i < idx; i++ {
		p = p.next
	}
	return p.ele
}

func (l *List) Delete(idx uint) {
	if idx >= l.len {
		panic("out of range")
	}
	p := l.head
	for i := uint(0); i < idx; i++ {
		p = p.next
	}

	if nil != p.prev {
		p.prev.next = p.next
	} else {
		l.head = p.next
		l.head.prev = nil
	}

	if nil != p.next {
		p.next.prev = p.prev
	} else {
		l.tail = p.prev
		l.tail.next = nil
	}

	l.len--
}

func (l *List) Insert(idx uint, ele Element) {
	if idx < 0 || idx >= l.len {
		panic("out of range")
	}

	p := l.head
	for i := uint(0); i < idx; i++ {
		p = p.next
	}
	nn := newNode(ele)

	if p == l.head {
		l.head = nn
	}

	nn.next = p
	nn.prev = p.prev
	if nil != p.prev {
		p.prev.next = nn
	}
	p.prev = nn
	l.len++
}

func (l *List) Reverse() {
	if l.len < 2 {
		return
	}

	p := l.head
	for p != nil {
		q := p.next

		p.next = p.prev
		p.prev = q
		p = q
	}
	h := l.head
	l.head = l.tail
	l.tail = h
}

func (l *List) Len() uint {
	return l.len
}

func (l *List) Empty() bool {
	return l.len == 0
}
