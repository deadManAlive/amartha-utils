package utils

import "fmt"

func PrintSlice(s []int) {
	for _, v := range s {
		fmt.Printf("%4d ", v)
	}
	fmt.Println()
}

func PrintSlice2(s [][]int) {
	for _, v := range s {
		for _, w := range v {
			if w == 0 {
				fmt.Printf("%4s ", ".")
			} else {
				fmt.Printf("%4d ", w)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func EmptySquareMatrixGenerator(n int) [][]int {
	mat := make([][]int, n)
	for i := range n {
		mat[i] = make([]int, n)
	}
	return mat
}
