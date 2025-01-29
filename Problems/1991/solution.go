package problem1991

func FindMiddleIndex(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }

    left := 0
    for i, num := range nums {
        sum -= num
        if left == sum {
            return i
        }
        left += num
    }

    return -1
}