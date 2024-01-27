package sudoku

//ToDO
//write func CreateGrid(args []string)([9][9] rune,bool)
//kanuarbe

func CreateGrid(args []string) ([9][9]rune, bool) {
	num := [9][9]rune{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if len(args[i]) != 9 {
				return num, false
			}
			num[i][j] = rune(args[i][j])
		}
	}
	return num, true
}
