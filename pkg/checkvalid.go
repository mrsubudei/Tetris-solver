package pkg

type Shape struct {
	x1 int
	y1 int
	x2 int
	y2 int
	x3 int
	y3 int
	x4 int
	y4 int
}

func Check(s string) bool {
	a1 := "#...#...#...#..."
	a2 := ".#...#...#...#.."
	a3 := "..#...#...#...#."
	a4 := "...#...#...#...#"
	a5 := "####............"
	a6 := "....####........"
	a7 := "........####...."
	a8 := "............####"

	b1 := "##..##.........."
	b2 := ".##..##........."
	b3 := "..##..##........"
	b4 := "....##..##......"
	b5 := ".....##..##....."
	b6 := "......##..##...."
	b7 := "........##..##.."
	b8 := ".........##..##."
	b9 := "..........##..##"

	c1 := "#...#...##......"
	c2 := ".#...#...##....."
	c3 := "..#...#...##...."
	c4 := "....#...#...##.."
	c5 := ".....#...#...##."
	c6 := "......#...#...##"
	c7 := "###.#..........."
	c8 := ".###.#.........."
	c9 := "....###.#......."
	c10 := ".....###.#......"
	c11 := "........###.#..."
	c12 := ".........###.#.."
	c13 := "##...#...#......"
	c14 := ".##...#...#....."
	c15 := "..##...#...#...."
	c16 := "....##...#...#.."
	c17 := ".....##...#...#."
	c18 := "......##...#...#"
	c19 := "..#.###........."
	c20 := "...#.###........"
	c21 := "......#.###....."
	c22 := ".......#.###...."
	c23 := "..........#.###."
	c24 := "...........#.###"
	c25 := ".#...#..##......"
	c26 := "..#...#..##....."
	c27 := "...#...#..##...."
	c28 := ".....#...#..##.."
	c29 := "......#...#..##."
	c30 := ".......#...#..##"
	c31 := "#...###........."
	c32 := ".#...###........"
	c33 := "....#...###....."
	c34 := ".....#...###...."
	c35 := "........#...###."
	c36 := ".........#...###"
	c37 := "##..#...#......."
	c38 := ".##..#...#......"
	c39 := "..##..#...#....."
	c40 := "....##..#...#..."
	c41 := ".....##..#...#.."
	c42 := "......##..#...#."
	c43 := "###...#........."
	c44 := ".###...#........"
	c45 := "....###...#....."
	c46 := ".....###...#...."
	c47 := "........###...#."
	c48 := ".........###...#"

	d1 := ".##.##.........."
	d2 := "..##.##........."
	d3 := ".....##.##......"
	d4 := "......##.##....."
	d5 := ".........##.##.."
	d6 := "..........##.##."
	d7 := "#...##...#......"
	d8 := ".#...##...#....."
	d9 := "..#...##...#...."
	d10 := "....#...##...#.."
	d11 := ".....#...##...#."
	d12 := "......#...##...#"

	e1 := ".#..###........."
	e2 := "..#..###........"
	e3 := ".....#..###....."
	e4 := "......#..###...."
	e5 := ".........#..###."
	e6 := "..........#..###"
	e7 := "#...##..#......."
	e8 := ".#...##..#......"
	e9 := "..#...##..#....."
	e10 := "....#...##..#..."
	e11 := ".....#...##..#.."
	e12 := "......#...##..#."
	e13 := "###..#.........."
	e14 := ".###..#........."
	e15 := "....###..#......"
	e16 := ".....###..#....."
	e17 := "........###..#.."
	e18 := ".........###..#."
	e19 := ".#..##...#......"
	e20 := "..#..##...#....."
	e21 := "...#..##...#...."
	e22 := ".....#..##...#.."
	e23 := "......#..##...#."
	e24 := ".......#..##...#"

	slice := []string{
		a1,
		a2,
		a3,
		a4,
		a5,
		a6,
		a7,
		a8,
		b1,
		b2,
		b3,
		b4,
		b5,
		b6,
		b7,
		b8,
		b9,
		c1,
		c2,
		c3,
		c4,
		c5,
		c6,
		c7,
		c8,
		c9,
		c10,
		c11,
		c12,
		c13,
		c14,
		c15,
		c16,
		c17,
		c18,
		c19,
		c20,
		c21,
		c22,
		c23,
		c24,
		c25,
		c26,
		c27,
		c28,
		c29,
		c30,
		c31,
		c32,
		c33,
		c34,
		c35,
		c36,
		c37,
		c38,
		c39,
		c40,
		c41,
		c42,
		c43,
		c44,
		c45,
		c46,
		c47,
		c48,
		d1,
		d2,
		d3,
		d4,
		d5,
		d6,
		d7,
		d8,
		d9,
		d10,
		d11,
		d12,
		e1,
		e2,
		e3,
		e4,
		e5,
		e6,
		e7,
		e8,
		e9,
		e10,
		e11,
		e12,
		e13,
		e14,
		e15,
		e16,
		e17,
		e18,
		e19,
		e20,
		e21,
		e22,
		e23,
		e24,
	}

	match := false
	for i := 0; i < len(slice); i++ {
		if s == slice[i] {
			match = true
		}
	}
	return match
}

func CheckTetro(tetromino [4][4]string) bool {
	c, d := 0, 0
	for a, elem := range tetromino {
		for b, elem2 := range elem {
			if elem2 != "." {
				d++
				if a+1 < 4 && tetromino[a+1][b] != "." {
					c++
				}
				if a-1 >= 0 && tetromino[a-1][b] != "." {
					c++
				}
				if b+1 < 4 && tetromino[a][b+1] != "." {
					c++
				}
				if b-1 >= 0 && tetromino[a][b-1] != "." {
					c++
				}
			}
		}
	}
	if d != 4 {
		return false
	}
	if c == 6 || c == 8 {
		return true
	}
	return false
}

func GetShape(n int) Shape {
	shape := Shape{}
	switch n {
	case 1:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 1
		shape.y2 = 0
		shape.x3 = 2
		shape.y3 = 0
		shape.x4 = 3
		shape.y4 = 0
	case 2:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 0
		shape.y2 = 1
		shape.x3 = 0
		shape.y3 = 2
		shape.x4 = 0
		shape.y4 = 3
	case 3:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 0
		shape.y2 = 1
		shape.x3 = 1
		shape.y3 = 0
		shape.x4 = 1
		shape.y4 = 1
	case 4:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 1
		shape.y2 = 0
		shape.x3 = 2
		shape.y3 = 0
		shape.x4 = 2
		shape.y4 = 1
	case 5:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 0
		shape.y2 = 1
		shape.x3 = 0
		shape.y3 = 2
		shape.x4 = 1
		shape.y4 = 0
	case 6:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 0
		shape.y2 = 1
		shape.x3 = 1
		shape.y3 = 1
		shape.x4 = 2
		shape.y4 = 1
	case 7:
		shape.x1 = 0
		shape.y1 = 2
		shape.x2 = 1
		shape.y2 = 0
		shape.x3 = 1
		shape.y3 = 1
		shape.x4 = 1
		shape.y4 = 2
	case 8:
		shape.x1 = 0
		shape.y1 = 1
		shape.x2 = 1
		shape.y2 = 1
		shape.x3 = 2
		shape.y3 = 0
		shape.x4 = 2
		shape.y4 = 1
	case 9:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 1
		shape.y2 = 0
		shape.x3 = 2
		shape.y3 = 0
		shape.x4 = 3
		shape.y4 = 0
	case 10:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 0
		shape.y2 = 1
		shape.x3 = 1
		shape.y3 = 0
		shape.x4 = 2
		shape.y4 = 0
	case 11:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 0
		shape.y2 = 1
		shape.x3 = 0
		shape.y3 = 2
		shape.x4 = 1
		shape.y4 = 2
	case 12:
		shape.x1 = 0
		shape.y1 = 1
		shape.x2 = 0
		shape.y2 = 2
		shape.x3 = 1
		shape.y3 = 0
		shape.x4 = 1
		shape.y4 = 1
	case 13:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 1
		shape.y2 = 0
		shape.x3 = 1
		shape.y3 = 1
		shape.x4 = 2
		shape.y4 = 1
	case 14:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 0
		shape.y2 = 1
		shape.x3 = 1
		shape.y3 = 1
		shape.x4 = 1
		shape.y4 = 2
	case 15:
		shape.x1 = 0
		shape.y1 = 1
		shape.x2 = 1
		shape.y2 = 0
		shape.x3 = 1
		shape.y3 = 1
		shape.x4 = 2
		shape.y4 = 1
	case 16:
		shape.x1 = 0
		shape.y1 = 1
		shape.x2 = 1
		shape.y2 = 0
		shape.x3 = 2
		shape.y3 = 0
		shape.x4 = 3
		shape.y4 = 0
	case 17:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 1
		shape.y2 = 0
		shape.x3 = 1
		shape.y3 = 1
		shape.x4 = 2
		shape.y4 = 0
	case 18:
		shape.x1 = 0
		shape.y1 = 0
		shape.x2 = 0
		shape.y2 = 1
		shape.x3 = 0
		shape.y3 = 2
		shape.x4 = 1
		shape.y4 = 1
	case 19:
		shape.x1 = 0
		shape.y1 = 1
		shape.x2 = 1
		shape.y2 = 0
		shape.x3 = 1
		shape.y3 = 1
		shape.x4 = 2
		shape.y4 = 1
	}

	return shape
}
