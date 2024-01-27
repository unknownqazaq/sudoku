package sudoku

// ToDO
// write func Checker(grid [9][9]rune,row,col int,num rune)bool
// tyergazy
func Checker(grid [9][9]rune, row, col int, num rune) bool {
	for i := 0; i < 9; i++ { //Checking horizontally is able to put number in grid
		if grid[row][i] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ { // Checking vertally is able to put number in grid
		if grid[i][col] == num {
			return false
		}
	}

	// row - row % 3 = to get where box row starts
	// col - col % 3 = to get where box col starts
	for i := 0; i < 3; i++ { // Checking in box[3][3] is able to put number in grid
		for j := 0; j < 3; j++ {
			if grid[i+(row-row%3)][j+(col-col%3)] == num {
				return false
			}
		}
	}

	return true
}
