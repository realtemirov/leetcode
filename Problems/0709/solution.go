package problem0709

func ToLowerCase(s string) string {
    b := []byte(s)
    for i, c := range b {
        if c >= 'A' && c <= 'Z' {
            b[i] += 32
        }
    }
    return string(b)
}