package problem0217

func ContainsDuplicate(nums []int) bool {
    if len(nums) < 2 {
        return false
    }

    m := make(map[int]bool)

    for _, num := range nums {
        if _, ok := m[num]; ok {
            return true
        }
        m[num]=true
    }

    return false
}