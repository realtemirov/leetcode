# Implement strStr()

Ikki torli `igna` va `pichan`ni hisobga olgan holda, `pichan`da `igna`ning birinchi paydo bo'lish `ko'rsatkichi`ni yoki agar `igna` `pichanning bir qismi` bo'lmasa `-1` ni qaytaring.

#### Example 1:

```
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" indeks 0 va 6 da uchraydi.
Birinchi hodisa 0 indeksida, shuning uchun biz 0 ni qaytaramiz.
```

#### Example 2:

```
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" "leetcode" da bo'lmagan, shuning uchun biz -1 ni qaytaramiz.
```

#### Cheklovlar:

* `1 <= haystack.length, needle.length <= 104`
* `haystack va needle faqat inglizcha kichik harflardan iborat.`

```go
func strStr(haystack string, needle string) int {

	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if len(haystack)-i < len(needle) {
				return -1
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}

	return -1
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/203/introduction-to-string/1161/)