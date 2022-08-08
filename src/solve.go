package src

func Solve(field *[][]string, shapes []Shape, n int, chars []string) bool {
	numOfFigures := len(shapes)

	putOk := false

	var s Shape
	s.x1 = shapes[n].x1 // copying shape
	s.x2 = shapes[n].x2
	s.x3 = shapes[n].x3
	s.x4 = shapes[n].x4
	s.y1 = shapes[n].y1
	s.y2 = shapes[n].y2
	s.y3 = shapes[n].y3
	s.y4 = shapes[n].y4
	if CheckExeed(field, shapes) {
		return false
	}

	for { // looping untill all positions in field tried to be put in

		if (*field)[s.x1][s.y1] == "." && (*field)[s.x2][s.y2] == "." && (*field)[s.x3][s.y3] == "." && (*field)[s.x4][s.y4] == "." { // if place is free, put tetro
			(*field)[s.x1][s.y1] = chars[n]
			(*field)[s.x2][s.y2] = chars[n]
			(*field)[s.x3][s.y3] = chars[n]
			(*field)[s.x4][s.y4] = chars[n]
			putOk = true
		} else {
			putOk = false
		}
		if (n == numOfFigures-1) && (putOk) { // condition for quiting from last recursion
			return true
		}

		if putOk {
			n++
			ok := Solve(field, shapes, n, chars) // calling recursive func for putting next tetro
			if ok {
				return true
			} else {
				n--
			}
		}

		if putOk {
			(*field)[s.x1][s.y1] = "." // removing previous tetro
			(*field)[s.x2][s.y2] = "."
			(*field)[s.x3][s.y3] = "."
			(*field)[s.x4][s.y4] = "."
		}
		if s.y1 != len((*field)[0])-1 && s.y2 != len((*field)[0])-1 && s.y3 != len((*field)[0])-1 && s.y4 != len((*field)[0])-1 {
			ShiftRight(&s) // changing position right to one dot, if possible
		} else {
			if s.x1 != len(*(field))-1 && s.x2 != len((*field))-1 && s.x3 != len((*field))-1 && s.x4 != len((*field))-1 {
				BackY(&s, &shapes[n])
				ShiftDown(&s) // changing down right to one dot, if possible
			} else {
				BackY(&s, &shapes[n]) // returning origin
				BackX(&s, &shapes[n])
				break
			}
		}

	}

	return false
}

func ShiftRight(s *Shape) {
	(*s).y1 += 1
	(*s).y2 += 1
	(*s).y3 += 1
	(*s).y4 += 1
}

func ShiftDown(s *Shape) {
	(*s).x1 += 1
	(*s).x2 += 1
	(*s).x3 += 1
	(*s).x4 += 1
}

func BackY(s, old *Shape) {
	(*s).y1 = (*old).y1
	(*s).y2 = (*old).y2
	(*s).y3 = (*old).y3
	(*s).y4 = (*old).y4
}

func BackX(s, old *Shape) {
	(*s).x1 = (*old).x1
	(*s).x2 = (*old).x2
	(*s).x3 = (*old).x3
	(*s).x4 = (*old).x4
}

func CheckExeed(field *[][]string, shapes []Shape) bool {
	num := len((*field)) - 1

	for i := 0; i < len(shapes); i++ {
		if shapes[i].x1 > num || shapes[i].x2 > num || shapes[i].x3 > num || shapes[i].x4 > num {
			return true
		} else if shapes[i].y1 > num || shapes[i].y2 > num || shapes[i].y3 > num || shapes[i].y4 > num {
			return true
		}
	}
	return false
}
