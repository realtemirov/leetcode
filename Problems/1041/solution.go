package problem1041

func IsRobotBounded(instructions string) bool {
	x, y, degree := 0, 0, 0

	for _, step := range instructions {
		if step == 'L' {
			degree = (degree + 3) % 4
			continue
		} else if step == 'R' {
			degree = (degree + 1) % 4
			continue
		}

		switch degree {
		case 0:
			y++
		case 1:
			x++
		case 2:
			y--
		case 3:
			x--
		}
	}

	return degree != 0 || x == 0 && y == 0
}
