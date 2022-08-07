package pkg

import (
	"testing"
)

type MakeFieldTest struct {
	arg int
	ans [][]string
}

var MakeFieldTests = []MakeFieldTest{
	{
		3,
		[][]string{
			{".", ".", "."},
			{".", ".", "."},
			{".", ".", "."},
		},
	},
}

func TestMakeField(t *testing.T) {
	for _, test := range MakeFieldTests {
		output := MakeField(test.arg)
		if len(output) != len(test.ans) {
			t.Errorf("Output %s not equal to expected %s", output, test.ans)
		}

	}
}
