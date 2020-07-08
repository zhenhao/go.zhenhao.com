package main

import (
	"fmt"
	"github.com/zhenhao/gofun/container"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	t := container.NewTree()
	for i := 0; i < 10; i++ {
		t.Insert(rand.Intn(100))
	}

	t.Traversal()
	fmt.Println()
	t.TraversalRecursive()
}
