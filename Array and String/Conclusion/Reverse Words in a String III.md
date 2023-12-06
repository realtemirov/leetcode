# Reverse Words in a String III

`s` qatori berilgan boʻlsa, boʻsh joy va boshlangʻich soʻz tartibini saqlab qolgan holda, gap ichidagi har bir soʻzdagi belgilar tartibini teskari oʻzgartiring.

#### Example 1:
```
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
```

#### Example 2:
```
Input: s = "Mr Ding"
Output: "rM gniD"
```

#### Cheklovlar

* `1 <= s.length <= 5 * 10^4`
* `s` bosib chiqarish mumkin bo'lgan **ASCII** belgilarni o'z ichiga oladi.
* `s` hech qanday oldingi yoki keyingi bo'shliqlarni o'z ichiga olmaydi.
* `s` ichida **kamida bitta so'z** bor.
* `S` harfidagi barcha so'zlar bitta bo'sh joy bilan ajratiladi

```go
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	var result strings.Builder
	for _, v := range arr {

		if v != "" {
			result.WriteString(reverse(v) + " ")
		}
	}

	return strings.TrimSpace(result.String())
}

func reverse(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteByte(s[i])
	}
	return result.String()
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1165/)