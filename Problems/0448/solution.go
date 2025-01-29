package problem0448

func FindDisappearedNumbers(nums []int) []int {
    arr := make([]bool, len(nums))

    for _, num :=range nums {
        arr[num-1] = true
    }

    idx := 0
    for i, b := range arr {
        if !b {
            nums[idx] = i+1
            idx++
        }
    }

    return nums[:idx]
}