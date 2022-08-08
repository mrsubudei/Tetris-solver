package src

import (
	"fmt"
	"math"
)

func MakeField(n int) [][]string {
	sqRoot := n * 4
	power := math.Sqrt(float64(sqRoot))
	rounded := int(power)

	if rounded*rounded < n*4 {
		rounded++
	}

	field := make([][]string, rounded)

	for i := 0; i < rounded; i++ {
		for j := 0; j < rounded; j++ {
			field[i] = append(field[i], ".")
		}
	}

	return field
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

func PrintAns(slice [][]string) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Print(slice[i][j])
		}
		fmt.Println()
	}
}
