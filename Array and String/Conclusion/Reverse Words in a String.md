# Reverse Words in a String

Kiritilgan `s qatori berilgan bo'lsa, **so'zlarning** tartibini o'zgartiring.

**So'z** bo'sh joy bo'lmagan belgilar ketma-ketligi sifatida aniqlanadi. `S` harfidagi **so'zlar** kamida bitta bo'sh joy bilan ajratiladi.

*Yagona boʻsh joy bilan birlashtirilgan soʻzlar qatorini teskari tartibda* qaytaring.

**E'tibor bering**, `s` ikkita so'z orasida bosh yoki keyingi bo'shliqlar yoki bir nechta bo'shliqlarni o'z ichiga olishi mumkin. Qaytarilgan satrda faqat so'zlarni ajratib turadigan bitta bo'sh joy bo'lishi kerak. Qo'shimcha bo'shliqlarni qo'shmang.

#### Example 1:
```
Input: s = "the sky is blue"
Output: "blue is sky the"
```

#### Example 2:
```
Input: s = "  hello world  "
Output: "world hello"
Explanation: Sizning teskari satringizda oldingi yoki keyingi bo'shliqlar bo'lmasligi kerak.
```

#### Example 3:
```
Input: s = "a good   example"
Output: "example good a"
Explanation: Ikki so'z orasidagi bir nechta bo'shliqni teskari satrda bitta bo'sh joyga qisqartirishingiz kerak.
```

#### Cheklovlar
* `1 <= s.length <= 10^4`
* `s` inglizcha harflarni (katta va kichik harflar), raqamlar va `' '` bo'shliqlarni o'z ichiga oladi.
* `s` ichida **kamida bitta so'z** bor.
 
**Kuzatuv**: Agar string maʼlumotlar turi sizning tilingizda oʻzgaruvchan boʻlsa, `O(1)` qoʻshimcha boʻsh joy bilan **oʻz oʻrnida** yecha olasizmi?

```go
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	var result strings.Builder
	for i := len(arr) - 1; i >= 0; i-- {

		if arr[i] != "" {
			result.WriteString(arr[i])
			if i != 0 {

				result.WriteString(" ")
			}
		}
	}

	return strings.TrimSpace(result.String())
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1164/)