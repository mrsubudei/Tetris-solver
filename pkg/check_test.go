package pkg

import "testing"

func BenchmarkCheckTetro(b *testing.B) {
	s := [4][4]string{
		{"#", "#", ".", "."},
		{"#", "#", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
	}

	for i := 0; i < b.N; i++ {
		CheckTetro(s)
	}
}
