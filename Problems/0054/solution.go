package problem0054

func SpiralOrder(matrix [][]int) []int {
    m, n := len(matrix), len(matrix[0])
    l, r, up, down := 0, n-1, 0, m-1
    nums := make([]int, 0, m*n)

    for {
        for i:= l; i<= r; i++ {
            nums = append(nums, matrix[up][i])
        }
        up++
        if up > down {
            break
        }

        for i:=up; i<=down; i++ {
            nums = append(nums, matrix[i][r])
        }
        r--
        if l>r {
            break
        }

        for i:=r; i>=l; i-- {
            nums = append(nums, matrix[down][i])
        }
        down--
        if down < up {
            break
        }

        for i:= down; i>=up; i-- {
            nums = append(nums, matrix[i][l])
        }
        l++
        if l > r {
            break
        }
    }

    return nums
}