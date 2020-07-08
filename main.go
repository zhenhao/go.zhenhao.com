package main

import (
	"fmt"
	"github.com/zhenhao/gofun/container"
	"github.com/zhenhao/gofun/sort"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var x [10]int

	t := container.NewTree()
	for i := 0; i < 10; i++ {
		x[i] = rand.Intn(100)
		t.Insert(x[i])
	}

	t.Traversal()
	t.TraversalRecursive()
	sort.QuickSort(x[:])
	fmt.Println(x)
}
