package problem0383

func CanConstruct(ransomNote string, magazine string) bool {
    m := [26]int{}

    for _, c := range magazine {
        m[c-'a']++
    }

    for _, c := range ransomNote {
        m[c-'a']--
        if m[c-'a'] < 0 {
            return false
        }
    }

    return true
}