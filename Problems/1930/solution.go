package problem1930

func CountPalindromicSubsequence(s string) int {
    firstIdx := make(map[byte]int, len(s))
    lastIdx := make(map[byte]int, len(s))

    for i:=0; i<len(s); i++{
        if _, ok := firstIdx[s[i]]; !ok {
            firstIdx[s[i]]=i
        }

        lastIdx[s[i]]=i
    }

    count := 0

    for char, start := range firstIdx {
        end := lastIdx[char]
        if start < end {
            uniqueChars := make(map[byte]bool, end-start)
            for i:=start+1; i<end; i++ {
                uniqueChars[s[i]]=true
            }

            count+=len(uniqueChars)
        }
    }

    return count
}