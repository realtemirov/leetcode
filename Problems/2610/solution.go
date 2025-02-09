package problem2610


func FindMatrix(nums []int) [][]int {
    m := map[int]int{}
    res := [][]int{}

    for _, num := range nums {
        idx := m[num]
        if idx + 1 > len(res) {
            res = append(res, []int{})
        }
        res[idx] = append(res[idx], num)
        m[num]++
    }

    return res
}