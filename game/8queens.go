package game

import (
	"fmt"
)

var QN int = 8
var qc int = 0

func Queens() {
	qs := make([]int, QN)
	setq(qs, 0)
}

func setq(qs []int, row int) {
	if row >= QN {
		drawQueens(qs)
		return
	}

	for i := 0; i < QN; i++ {
		if !attack(qs, row, i) {
			qs[row] = i
			setq(qs, row+1)
		}
	}
}

func attack(qs []int, row int, col int) bool {
	for i := 0; i < row; i++ {
		q := qs[i]
		if q == col || i+q == row+col || q-i == col-row {
			return true
		}
	}
	return false
}

func drawQueens(qs []int) {
	qc++
	fmt.Printf("======= %d =======\n", qc)
	for _, q := range qs {
		for i := 0; i < QN; i++ {
			if i == q {
				fmt.Print("Q")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}
