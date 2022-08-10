package main

import "fmt"

func main() {
	var s Sudoku

	for i := 0; i < 9; i++ {
		s.SetInitialValue(i, i, 1)
	}

	fmt.Println(s.ToString())
}
