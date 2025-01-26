package problem0905

func SortArrayByParity(nums []int) []int {
    idx := 0

    for i := range nums {
        if nums[i] % 2 == 0 {
            nums[i], nums[idx] = nums[idx], nums[i]
            idx++
        }
    }

    return nums
}