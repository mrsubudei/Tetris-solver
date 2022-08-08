package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tetris/src"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("wrong number of arguments")
		return
	}

	f, err := os.Open(arg[0]) // opening file
	if err != nil {
		fmt.Println("Wrong path")
		return
	}
	defer f.Close()

	d := bufio.NewScanner(f)
	preData := ""
	count := 0
	for d.Scan() {
		count++
		preData += d.Text()
		if count == 5 {

			if len(preData)%16 != 0 {
				fmt.Println("ERROR")
				return
			}
			count = 0
		}

	}
	if len(preData) == 0 {
		fmt.Println("At least one tetromino has to be in the file")
		return
	}

	data := ""
	tmp := ""

	for i := 0; i < len(preData); i++ { // dividing peaces
		data += string(preData[i])
		tmp += string(preData[i])
		if len(tmp)%16 == 0 {
			data += " "
		}
	}

	sorted := strings.Fields(data)

	pos := 0
	valid := false
	var sliceMinos []src.Shape

	for i := 0; i < len(sorted); i++ { // checking is data valid
		valid, pos = src.CheckValid(sorted[i])
		if !valid {
			fmt.Println("ERROR")
			return
		}
		sliceMinos = append(sliceMinos, src.GetShape(pos)) // getting shapes for each tetrominoes
	}

	numOfMinos := len(sorted)
	field := src.MakeField(numOfMinos)

	solved := false
	n := 0
	var chars []string

	for i := 0; i < numOfMinos; i++ { // creating slice of character for answer. A, B, C etc.
		chars = append(chars, string(rune(i+65)))
	}

	for !solved {
		solved = src.Solve(&field, sliceMinos, n, chars)
		if !solved {
			field = src.IncreaseField(field)
			n = 0
		}
	}
	src.PrintAns(field)
}
