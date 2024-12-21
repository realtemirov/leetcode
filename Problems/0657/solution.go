package problem0657

func JudgeCircle(moves string) bool {
	if len(moves)%2 == 1 {
		return false
	}

	x, y := 0, 0
	for _, move := range moves {
		switch move {
		case 'U':
			x++
		case 'D':
			x--
		case 'L':
			y--
		case 'R':
			y++
		}
	}

	return x == 0 && y == 0
}
