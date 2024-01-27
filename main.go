package main

import (
	"os"
	"piscine/sudoku"
)

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		sudoku.ErrorPrint()
		return
	}

	grid, check := sudoku.CreateGrid(args)
	if !check {
		sudoku.ErrorPrint()
		return
	}

	if sudoku.Solver(&grid, 0, 0) == 1 {

		sudoku.Res(&grid, 0, 0)
		sudoku.Print(grid)
	} else {
		sudoku.ErrorPrint()
	}
}
