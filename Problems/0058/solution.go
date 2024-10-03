package problem0058

func LengthOfLastWord(s string) int {
    cnt := 0
    existEmpty := false
    for _, v := range s {
        if v == ' ' {
            existEmpty = true
        } else {
            if existEmpty {
                cnt = 0
                existEmpty = false
            }
            cnt++
        }
    }

    return cnt
}