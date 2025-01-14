package problem1346

func CheckIfExist(arr []int) bool {
	m := make(map[int]bool, len(arr))

	for _, num := range arr {
		if _, ok := m[num*2]; ok || (num%2 == 0 && m[num/2]) {
			return true
		}
		m[num] = true
	}

	return false
}
