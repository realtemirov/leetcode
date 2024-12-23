package problem1572

func DiagonalSum(mat [][]int) int {
    sum := 0

    l, r := 0, len(mat)-1
    for _, nums := range mat {
        if l == r {
            sum += nums[l]
        } else {
            sum += nums[l] + nums[r]
        }

        l++
        r--
    }

    return sum
}