package game

import (
	"fmt"
	"math/rand"
	"time"
)

func Guess() {
	rand.Seed(int64(time.Now().Nanosecond()))
	x := rand.Int() % 100
	var in int

Loop:
	for {
		fmt.Print("Input a number(0~99):")
		fmt.Scanf("%d", &in)
		switch {
		case in == x:
			fmt.Printf("Congratuations! The number is %d\n", x)
			break Loop
		case in > x:
			fmt.Println("Bigger")
			break
		case in < x:
			fmt.Println("Smaller")
			break
		}
	}
}

func Fibonacci(n uint) {
	f := gen()

	for i := uint(0); i < n; i++ {
		fmt.Printf("%d ", f())
	}
	fmt.Println()
}

func gen() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
