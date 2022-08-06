package pkg

import (
	"fmt"
	"math"
)

func MakeField(n int) [][]string {
	q := n * 4
	s := math.Sqrt(float64(q))
	size := int(s)

	field := make([][]string, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			field[i] = append(field[i], ".")
		}
	}

	return field
}

func Printans(slice [][]string) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Print(slice[i][j])
		}
		fmt.Println()
	}
}

func removeMinos(field [][]string, s Shape) [][]string {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if s.x1 == i && s.y1 == j {
				field[i][j] = "."
			}
		}
	}

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if s.x2 == i && s.y2 == j {
				field[i][j] = "."
			}
		}
	}

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if s.x3 == i && s.y3 == j {
				field[i][j] = "."
			}
		}
	}

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if s.x4 == i && s.y4 == j {
				field[i][j] = "."
			}
		}
	}
	return field
}
