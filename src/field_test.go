package src

import (
	"testing"
)

type MakeFieldTest struct {
	arg int
	ans [][]string
}

var MakeFieldTests = []MakeFieldTest{
	{
		5,
		[][]string{
			{".", ".", ".", ".", "."},
			{".", ".", ".", ".", "."},
			{".", ".", ".", ".", "."},
			{".", ".", ".", ".", "."},
			{".", ".", ".", ".", "."},
		},
	},
	{
		3,
		[][]string{
			{".", ".", ".", "."},
			{".", ".", ".", "."},
			{".", ".", ".", "."},
			{".", ".", ".", "."},
		},
	},
	{
		4,
		[][]string{
			{".", ".", ".", "."},
			{".", ".", ".", "."},
			{".", ".", ".", "."},
			{".", ".", ".", "."},
		},
	},

	{
		2,
		[][]string{
			{".", ".", "."},
			{".", ".", "."},
			{".", ".", "."},
		},
	},
	{
		6,
		[][]string{
			{".", ".", ".", ".", "."},
			{".", ".", ".", ".", "."},
			{".", ".", ".", ".", "."},
			{".", ".", ".", ".", "."},
			{".", ".", ".", ".", "."},
		},
	},
}

func TestMakeField(t *testing.T) {
	for _, test := range MakeFieldTests {
		strExp := ""
		strAns := ""
		output := MakeField(test.arg)
		for i := 0; i < len(output); i++ {
			for j := 0; j < len(output[i]); j++ {
				strExp += output[i][j]
			}
		}

		for i := 0; i < len(test.ans); i++ {
			for j := 0; j < len(test.ans[i]); j++ {
				strAns += test.ans[i][j]
			}
		}

		if strExp != strAns {
			t.Errorf("Output %s not equal to expected %s", output, test.ans)
		}

	}
}

func TestIncreaseField(t *testing.T) {
	field := [][]string{
		{".", "."},
		{".", "."},
	}
	got := IncreaseField(field)
	want := [][]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}

	strExp := ""
	strAns := ""
	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[i]); j++ {
			strExp += got[i][j]
		}
	}

	for i := 0; i < len(want); i++ {
		for j := 0; j < len(want[i]); j++ {
			strAns += want[i][j]
		}
	}

	if strExp != strAns {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
