# Plus One

Sizga butun sonlar qatori `digits` sifatida berilgan katta butun son berilgan, bu erda har bir `digits[i]` butun sonning <code>i<sup>th</sup></code> raqamidir. Raqamlar chapdan o'ngga tartibda eng muhimdan eng muhimgacha tartiblangan. Katta butun sonda bosh `0` lar mavjud emas.

Katta butun sonni bittaga oshiring va natijada `digits` qatorini qaytaring.

#### Example 1:
```
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: Massiv 123 butun sonni ifodalaydi.
Birga ko'paytirilsa, 123 + 1 = 124 hosil bo'ladi. 
Shunday qilib, natija [1,2,4] bo'lishi kerak.
```

#### Example 2:
```
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: Massiv 4321 butun sonini ifodalaydi.
Birga ko'paytirilsa, 4321 + 1 = 4322 hosil bo'ladi. 
Shunday qilib, natija [4,3,2,2] bo'lishi kerak.
```

#### Example 3:
```
Input: digits = [9]
Output: [1,0]
Explanation: Massiv butun 9 ni ifodalaydi.
Birga ko'paytirilsa, 9 + 1 = 10 bo'ladi.
Shunday qilib, natija [1,0] bo'lishi kerak.
```

Constraints:

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.

```go
func plusOne(digits []int) []int {
	n := len(digits) - 1

	for i := n; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i]++
			break
		}

	}

	return digits
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/201/introduction-to-array/1148/)