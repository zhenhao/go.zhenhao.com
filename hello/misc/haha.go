package misc

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

func Add(a int, b int) int {
	return a + b
}

type Abser interface {
	Abs() float64
}

type Point struct {
	X, Y float64
}

func (v *Point) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Point) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Point) String() string {
	return fmt.Sprintf("{X=%f,Y=%f}", v.X, v.Y)
}

func init() {
	fmt.Println("haha.init.1")
}

func init() {
	fmt.Println("haha.init.2")
}

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func fibonacci2(c, quit chan uint) {
	var x, y uint = 0, 1
	for {
		select {
		case c <- y:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func game() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100)
	var x int = -1
	for x != target {
		fmt.Print("Input a number(0~100): ")
		fmt.Scanf("%d", &x)
		switch {
		case x < target:
			fmt.Println("Less")
			break
		case x > target:
			fmt.Println("Bigger")
			break
		default:
			fmt.Printf("Bingo! the target is %d\n", x)
			break
		}
	}
}

func startServe() {
	var h Hello
	err := http.ListenAndServe("127.0.0.1:8080", h)
	if err != nil {
		log.Fatal(err)
	}
}
