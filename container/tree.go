package container

import (
	"fmt"
)

type treeNode struct {
	left   *treeNode
	right  *treeNode
	parent *treeNode
	ele    Element
	level  uint
}

func newTreeNode(ele Element) *treeNode {
	return &treeNode{
		ele:    ele,
		left:   nil,
		right:  nil,
		level:  0,
		parent: nil,
	}
}

func (n *treeNode) insert(node *treeNode) {
	if node.ele.(int) > n.ele.(int) {
		if nil == n.right {
			n.right = node
			node.parent = n
			node.level = n.level + 1
		} else {
			n.right.insert(node)
		}
	} else {
		if nil == n.left {
			n.left = node
			node.parent = n
			node.level = n.level + 1
		} else {
			n.left.insert(node)
		}
	}
}

func (n *treeNode) traversalRecursive() {
	if nil != n.left {
		n.left.traversalRecursive()
	}
	fmt.Printf("%d(%d) ", n.ele, n.level)
	if nil != n.right {
		n.right.traversalRecursive()
	}
}

func (n *treeNode) traversal() {
	s := NewStack()
	p := n
	for nil != p {
		s.Push(p)
		p = p.left
	}

	for !s.Empty() {
		node := s.Pop().(*treeNode)
		fmt.Printf("%d(%d) ", node.ele, node.level)
		p = node.right
		for nil != p {
			s.Push(p)
			p = p.left
		}
	}
}

func (n *treeNode) breadthFirst() {
	q := NewQueue()
	q.Push(n)

	for !q.Empty() {
		node := q.Pop().(*treeNode)
		if nil == node.parent {
			fmt.Printf("%d ", node.ele)
		} else {
			fmt.Printf("%d(%d) ", node.ele, node.parent.ele)
		}
		if nil != node.left {
			q.Push(node.left)
		}

		if nil != node.right {
			q.Push(node.right)
		}
	}
}

type Tree struct {
	root *treeNode
}

func NewTree() *Tree {
	return &Tree{nil}
}

func (t *Tree) Insert(ele Element) {
	node := newTreeNode(ele)

	if nil == t.root {
		node.level = 1
		t.root = node
	} else {
		t.root.insert(node)
	}
}

func (t *Tree) TraversalRecursive() {
	if nil == t.root {
		return
	}

	t.root.traversalRecursive()
	fmt.Println()
}

func (t *Tree) Traversal() {
	if nil == t.root {
		return
	}

	t.root.traversal()
	fmt.Println()
}

func (t *Tree) BreadthFirst() {
	if nil == t.root {
		return
	}

	t.root.breadthFirst()
	fmt.Println()
}
