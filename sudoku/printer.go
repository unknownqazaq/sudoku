package sudoku

import "github.com/01-edu/z01"

// ToDO
// write func Print() and funcErrorPrint()
// dmukassa
const err = "Error\n"

func Print(arr [9][9]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(rune(arr[i][j]))
			if j != 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}

func ErrorPrint() {
	for _, v := range err {
		z01.PrintRune(v)
	}
}
