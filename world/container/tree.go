package container

import (
	"fmt"
)

type treeNode struct {
	left  *treeNode
	ele   Element
	right *treeNode
}

func newTreeNode(ele Element) *treeNode {
	return &treeNode{nil, ele, nil}
}

func (n *treeNode) insert(ele Element) {
	if ele.(int) > n.ele.(int) {
		if nil == n.right {
			n.right = newTreeNode(ele)
		} else {
			n.right.insert(ele)
		}
	} else {
		if nil == n.left {
			n.left = newTreeNode(ele)
		} else {
			n.left.insert(ele)
		}
	}
}

func (n *treeNode) dump() {
	if nil != n.left {
		n.left.dump()
	}
	fmt.Printf("%d ", n.ele)
	if nil != n.right {
		n.right.dump()
	}
}

type Tree struct {
	root *treeNode
}

func NewTree() *Tree {
	return &Tree{nil}
}

func (t *Tree) Insert(ele Element) {
	if nil == t.root {
		t.root = newTreeNode(ele)
	} else {
		t.root.insert(ele)
	}
}

func (t *Tree) Dump() {
	if nil == t.root {
		return
	} else {
		t.root.dump()
	}
}
