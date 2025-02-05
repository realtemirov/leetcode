package problem1790

func AreAlmostEqual(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    var (
        has, equal bool
        idx int
    )

    for i, s := range s1 {
        if s == rune(s2[i]) {
            continue
        }    

        if has {
            if equal {
                return false
            }
    
            if rune(s2[idx]) == s && s2[i] == s1[idx] {
                equal = true
            } else {
                return false
            }
        }

        has = true
        idx = i
    }
    if has {
        return equal
    }
    return true
}