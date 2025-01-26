package problem1299

func ReplaceElements(arr []int) []int {
	num := -1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > num {
			num, arr[i] = arr[i], num
		} else {
			arr[i] = num
		}
	}

	return arr
}
