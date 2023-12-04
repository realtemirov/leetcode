# Pascal's Triangle II

Butun son `rowIndex` berilgan bo'lsa, **Paskal uchburchagining** `rowIndex^th` **(0-indeksli)** qatorini qaytaring.

**Paskal uchburchagida** har bir raqam ko'rsatilganidek, to'g'ridan-to'g'ri ustidagi ikkita raqamning yig'indisidir:

![triangle](PascalTriangleAnimated2.gif)

#### Example 1:
```
Input: rowIndex = 3
Output: [1,3,3,1]
```

#### Example 2:
```
Input: rowIndex = 0
Output: [1]
```

#### Example 3:
```
Input: rowIndex = 1
Output: [1,1]
```

#### Cheklovlar
* `0 <= rowIndex <= 33`
 

**Kuzatish**: Algoritmingizni faqat O (rowIndex) qo'shimcha joydan foydalanish uchun optimallashtira olasizmi?

```go
func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1,1}
	} else if rowIndex == 2 {
		return []int{1,2,1}
	}

  res := [][]int{{1}} 
	for i := 2; i <= rowIndex+1; i++ {
		temp := []int{1}
		for j := 1; j < i-1; j++ {
			temp = append(temp, res[len(res)-1][j]+res[len(res)-1][j-1])
		}
		temp = append(temp, 1)
		res = append(res, temp)
	}
	return res[rowIndex]
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1171/)