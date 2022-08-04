package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("wrong number of arguments")
		return
	}

	file := filepath.Join(arg[0], "text.txt")

	/*d, err := os.ReadFile(file)
	if err != nil {
		fmt.Print("Wrong path")
	}
	data := string(d)
	fmt.Println(data)*/

	f, err := os.Open(file) // opening file
	if err != nil {
		fmt.Print(file)
		fmt.Println("Wrong path")
		return
	}
	defer f.Close()

	d := bufio.NewScanner(f)

	preData := ""

	for d.Scan() {
		preData += d.Text()
	}

	data := ""
	tmp := ""

	for i := 0; i < len(preData); i++ {
		data += string(preData[i])
		tmp += string(preData[i])
		if len(tmp)%16 == 0 {
			data += " "
		}
	}

	sorted := strings.Fields(data)
	/*
		for i := 0; i < len(sorted); i++ {
			if !tetris.Check(sorted[i]) {
				fmt.Println("ERROR")
				return
			}
		}
		fmt.Println("valid data")
	*/
	fmt.Println(sorted)
}