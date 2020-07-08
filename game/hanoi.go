package game

import "fmt"

func Hanoi(n uint) {
	put('A', 'B', 'C', n)
}

func put(a byte, b byte, c byte, n uint) {
	if n == 1 {
		fmt.Printf("%c -> %c\n", a, c)
	} else if n > 1 {
		put(a, c, b, n-1)
		fmt.Printf("%c -> %c\n", a, c)
		put(b, a, c, n-1)
	}
}
