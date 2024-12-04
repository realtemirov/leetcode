package problem0219

func ContainsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int)

    for i, num := range nums {
        if j, ok := m[num]; ok && abs(i-j) <= k{
            return true
        }
        m[num]=i
    }

    return false
}

func abs(x int) int {
    if x < 0 {
        x *= -1
    }
    return x
}