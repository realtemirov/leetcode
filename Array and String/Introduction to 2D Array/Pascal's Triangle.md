# Pascal's Triangle

Berilgan butun son `numRows`, **Paskal uchburchagining** birinchi numRows ni qaytaring.

**Paskal uchburchagida** har bir raqam, ko'rsatilganidek, to'g'ridan-to'g'ri ustidagi ikkita raqamning yig'indisidir:

![triangle](PascalTriangleAnimated2.gif)

#### Example 1:

```
Input: numRows = 5 
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```

#### Example 2:
```
Input: numRows = 1
Output: [[1]]
```

#### Constraints:

* `1 <= numRows <= 30`

```go
func generate(numRows int) [][]int {
	res := [][]int{{1}}

	for i := 2; i <= numRows; i++ {
		temp := []int{1}
		for j := 1; j < i-1; j++ {
			temp = append(temp, res[len(res)-1][j]+res[len(res)-1][j-1])
		}
		temp = append(temp, 1)
		res = append(res, temp)
	}

	return res
}

```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/202/introduction-to-2d-array/1170/)