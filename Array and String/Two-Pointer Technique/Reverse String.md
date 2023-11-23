# Reverse String

Satrni teskari aylantiruvchi funksiya yozing. Kirish qatori `s` belgilar massivi sifatida beriladi.

Buni kiritish massivini `O(1)` qo'shimcha xotira bilan joyida o'zgartirish orqali qilishingiz kerak.

#### Example 1:
```
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
```

#### Example 2:
```
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
```

#### Cheklovlar:

* <code>1 <= s.length <= 10<sup>5</sup></code>`
* `s[i] bosma ascii belgisi.`

> Hint-1
>
> Satrni teskari aylantirishning butun mantig'i qarama-qarshi yo'nalishli ikki nuqtali yondashuvdan foydalanishga asoslangan!

```go
func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1183/)