# Minimum Size Subarray Sum

Musbat butun sonlar massivi `nums` va musbat butun sonli `target` berilgan, *yig'indisi `target`dan katta yoki teng bo'lgan pastki qatorning minimal uzunligini* qaytaring. Agar bunday pastki qator bo'lmasa, uning o'rniga `0` ni qaytaring.

#### Example 1:
```
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: Kichik massiv [4,3] muammo cheklovi ostida minimal uzunlikka ega.
```

#### Example 2:
```
Input: target = 4, nums = [1,4,4]
Output: 1
```

#### Example 3:
```
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
```

#### Cheklovlar:

* `1 <= target <= 10^9`
* `1 <= nums.length <= 10^5`
* `1 <= nums[i] <= 10^4`

**Kuzatish**: Agar siz `O(n)` yechimini topgan boʻlsangiz, vaqt murakkabligi `O(n log(n))` boʻlgan boshqa yechimni kodlab koʻring.

```go
func minSubArrayLen(target int, nums []int) int {
	l, s := 0, 0

	m := math.MaxInt32

	for r := 0; r < len(nums); r++ {
		s += nums[r]

		for s >= target {
			m = min(m, r-l+1)
			s -= nums[l]
			l++
		}
	}

	if m == math.MaxInt32 {
		return 0
	} else {
		return m
	}
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1299/)