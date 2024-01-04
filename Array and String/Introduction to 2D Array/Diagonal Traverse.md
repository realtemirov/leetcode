# Diagonal Traverse

`m x n` matritsa mat berilgan bolsa, diagonal tartibda massivning barcha elementlaridan iborat massivni qaytaring.

#### Example 1:

![matrix](<image-2 (1).png>)

```
Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,4,7,5,3,6,8,9]
```

#### Example 2:

```
Input: mat = [[1,2],[3,4]]
Output: [1,2,3,4]
```

#### Cheklovlar:

* `m == mat.length`
* `n == mat[i].length`
* `1 <= m, n <= 104`
* `1 <= m * n <= 104`
* `-105 <= mat[i][j] <= 105`

```go
func findDiagonalOrder(mat [][]int) []int {

	up := true
	c, l, r := 0, 0, 0
	m := len(mat)
	n := len(mat[0])
	res := make([]int, m*n)
	for c != m*n {
		res[c] = mat[l][r]
		c++
		if (l == 0 || r == n-1) && up {
			up = false
			if l == 0 && r != n-1 {
				r++
			} else {
				l++
			}
		} else if (r == 0 || l == m-1) && !up {
			up = true
			if r == 0 && l != m-1 {
				l++
			} else {
				r++
			}
		} else {

			if up {
				l--
				r++
			} else {
				l++
				r--
			}
		}
	}
	return res
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/202/introduction-to-2d-array/1167/)
