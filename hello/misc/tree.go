package misc

import "fmt"

type treeNode struct {
	Left  *treeNode
	Value int
	Right *treeNode
}

func (n *treeNode) insert(ele int) {
	if ele >= n.Value {
		if n.Right == nil {
			n.Right = &treeNode{nil, ele, nil}
		} else {
			n.Right.insert(ele)
		}
	} else {
		if n.Left == nil {
			n.Left = &treeNode{nil, ele, nil}
		} else {
			n.Left.insert(ele)
		}
	}
}

func (n *treeNode) mSequence(c chan int) {
	if nil != n.Left {
		n.Left.mSequence(c)
	}
	c <- n.Value
	if nil != n.Right {
		n.Right.mSequence(c)
	}
}

type Tree struct {
	root *treeNode
	len  int
}

func (t *Tree) Insert(ele int) {
	if t.root == nil {
		t.root = &treeNode{nil, ele, nil}
	} else {
		t.root.insert(ele)
	}
	t.len++
}

func (t *Tree) MSequence() {
	if t.root == nil {
		return
	}
	c := make(chan int, t.len)
	t.root.mSequence(c)
	for i := 0; i < t.len; i++ {
		fmt.Println(<-c)
	}
}

func (t *Tree) Len() int {
	return t.len
}

func (t *Tree) String() string {
	ch := make(chan int, t.len)
	t.root.mSequence(ch)
	ret := "["
	for i := 0; i < t.len; i++ {
		ret += fmt.Sprintf("%d, ", <-ch)
	}
	return ret + "]"
}
