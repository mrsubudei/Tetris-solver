package src

import (
	"testing"
)

type CheckTest struct {
	arg string
	ans bool
	num int
}

var CheckTests = []CheckTest{
	{
		"#...#...#...#...",
		true,
		1,
	},
	{
		"......#.###.....",
		true,
		7,
	},
	{
		".##.##..........",
		true,
		12,
	},
	{
		".#...##..#......",
		true,
		17,
	},
	{
		"#...#....",
		false,
		0,
	},
	{
		"............",
		false,
		0,
	},
	{
		"dgdfhjkjhg",
		false,
		0,
	},
}

func TestCheckValid(t *testing.T) {
	for _, test := range CheckTests {
		output, numb := CheckValid(test.arg)
		if output != test.ans {
			t.Errorf("Output %t not equal to expected %t", output, test.ans)
		}
		if numb != test.num {
			t.Errorf("Output %d not equal to expected %d", numb, test.num)
		}
	}
}

type GetShapeTest struct {
	arg int
	ans Shape
}

var GetShapeTests = []GetShapeTest{
	{
		1,
		Shape{0, 0, 1, 0, 2, 0, 3, 0},
	},
	{
		2,
		Shape{0, 0, 0, 1, 0, 2, 0, 3},
	},
	{
		12,
		Shape{0, 1, 0, 2, 1, 0, 1, 1},
	},
	{
		13,
		Shape{0, 0, 1, 0, 1, 1, 2, 1},
	},
}

func TestGetShape(t *testing.T) {
	for _, test := range GetShapeTests {
		output := GetShape(test.arg)
		if output != test.ans {
			t.Errorf("Output %d not equal to expected %d", output, test.ans)
		}
	}
}
