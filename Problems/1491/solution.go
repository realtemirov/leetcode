package problem1491

func Average(salary []int) float64 {
	var maxNum, minNum = salary[0], salary[0]
	var sumArr int

	for _, val := range salary {
		sumArr += val

		if val > maxNum {
			maxNum = val
		}

		if val < minNum {
			minNum = val
		}
	}

	return float64(sumArr-minNum-maxNum) / float64(len(salary)-2)
}