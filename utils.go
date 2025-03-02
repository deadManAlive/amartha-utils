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

func EmptyMatrixGenerator(y, x int) [][]int {
	mat := make([][]int, y)
	for i := range y {
		mat[i] = make([]int, x)
	}
	return mat
}

// embed other matrix inside other matrix (!!no bound checking!!)
func EmbedMatrix(source, destination [][]int, istart, jstart int) {
	for i, u := range source {
		for j, v := range u {
			destination[istart+i][jstart+j] = v
		}
	}
}
