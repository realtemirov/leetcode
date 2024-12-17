package problem1822

func ArraySign(nums []int) int {
    var sign int8 = 1

    for _, v := range nums {
        if v == 0 {
            return 0
        }

        if v < 0 {
            sign *= -1
        } 
    }
    return int(sign)
}