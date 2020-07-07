package main

import (
	"fmt"
	"github.com/zhenhao/gofun/container"
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
