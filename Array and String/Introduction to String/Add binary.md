# Add binary

Ikkita `a` va `b` ikkilik satrlar berilgan bo'lsa, *ularning yig'indisini ikkilik qator sifatida* qaytaring.

#### Example 1:

```
Input: a = "11", b = "1"
Output: "100"
```

#### Example 2:

```
Input: a = "1010", b = "1011"
Output: "10101"
```
 

#### Cheklovlar:

* `1 <= a.length, b.length <= 104`
* `a va b faqat "0" yoki "1" belgilardan iborat.`
* `Har bir satrda noldan tashqari bosh nollar mavjud emas.`

```go
func addBinary(a string, b string) string {

	str := ""
	var carry int8 = 0

	for i, j := len(a), len(b); i >= 1 || j >= 1; {
		i--
		j--
		var digitA, digitB int8

		if i >= 0 {
			digitA = int8(a[i] - '0')
		}

		if j >= 0 {
			digitB = int8(b[j] - '0')
		}

		total := carry + digitA + digitB

		c := fmt.Sprintf("%d", total%2)
		str = c + str

		carry = total / 2
	}

	if carry > 0 {
		str = "1" + str
	}

	return str
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/203/introduction-to-string/1160/)