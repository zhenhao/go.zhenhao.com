package main

import (
	"fmt"
	"go.zhenhao.com/world/container"
)

func main() {
	s := container.NewQueue()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	for !s.Empty() {
		fmt.Println(s.Pop())
	}
}
