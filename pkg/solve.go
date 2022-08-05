package pkg

func PutFigure(field [][]string, s []Shape) [][]string {
	pieces := 0
	tetro := s
	numOfFig := len(s)

	var chars []string

	for i := 0; i < numOfFig; i++ {
		chars = append(chars, string(rune(i+65)))
	}
	putFig := 0
	u := 0

	for putFig != numOfFig {

		for i := 0; i < len(field); i++ {
			for j := 0; j < len(field[i]); j++ {
				if tetro[u].x1 == i && tetro[u].y1 == j && field[i][j] == "." {
					field[i][j] = chars[u]
					pieces++
				}
			}
		}

		for i := 0; i < len(field); i++ {
			for j := 0; j < len(field[i]); j++ {
				if tetro[u].x2 == i && tetro[u].y2 == j && field[i][j] == "." {
					field[i][j] = chars[u]
					pieces++
				}
			}
		}

		for i := 0; i < len(field); i++ {
			for j := 0; j < len(field[i]); j++ {
				if tetro[u].x3 == i && tetro[u].y3 == j && field[i][j] == "." {
					field[i][j] = chars[u]
					pieces++
				}
			}
		}

		for i := 0; i < len(field); i++ {
			for j := 0; j < len(field[i]); j++ {
				if tetro[u].x4 == i && tetro[u].y4 == j && field[i][j] == "." {
					field[i][j] = chars[u]
					pieces++
				}
			}
		}

		if pieces != 4 {
			pieces = 0
			field[tetro[u].x1][tetro[u].y1] = "."
			field[tetro[u].x2][tetro[u].y2] = "."
			if len(field) >= 3 {
				field[tetro[u].x3][tetro[u].y3] = "."
			}
			if len(field) >= 4 {
				field[tetro[u].x4][tetro[u].y4] = "."
			}

			if tetro[u].y1 != len(field[0]) && tetro[u].y2 != len(field[0]) && tetro[u].y3 != len(field[0]) && tetro[u].y4 != len(field[0]) {
				tetro[u] = ShiftRight(tetro[u])
			} else {
				if tetro[u].x1 != len(field) && tetro[u].x2 != len(field) && tetro[u].x3 != len(field) && tetro[u].x4 != len(field) {
					tetro[u].y1 = s[u].y1
					tetro[u].y2 = s[u].y2
					tetro[u].y3 = s[u].y3
					tetro[u].y4 = s[u].y4
					tetro[u] = ShiftDown(tetro[u])
				} else {
					field = IncreaseField(field)
				}
			}
		} else if pieces == 4 {
			pieces = 0
			putFig++
			u++
		}
	}
	return field
}

func ShiftRight(s Shape) Shape {
	s.y1 += 1
	s.y2 += 1
	s.y3 += 1
	s.y4 += 1
	return s
}

func ShiftDown(s Shape) Shape {
	s.x1 += 1
	s.x2 += 1
	s.x3 += 1
	s.x4 += 1
	return s
}

func IncreaseField(field [][]string) [][]string {
	for i := 0; i < len(field); i++ {
		field[i] = append(field[i], ".")
	}

	num := len(field[0])
	var s []string
	for i := 0; i < num; i++ {
		s = append(s, ".")
	}

	field = append(field, s)
	return field
}

/* isEmpty := true

if isEmpty {
	return true
}
for num := 1; num <= 9; num++ {
	if CheckNumber(arr, row, col, num) {
		(arr)[row][col] = num

		if SolvePuzzle(arr) {
			return true
		} else {
			(arr)[row][col] = 0
		}
	}
}*/
