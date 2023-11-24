# Array Partition I

`2n` ta butun sonli `nums` massivi berilgan bo'lsa, bu butun sonlarni `n` ta juftlik `(a1, b1), (a2, b2), ..., (an, bn)` ga shunday guruhlangki, barcha `i` uchun `min(ai, bi)` yig'indisi `maksimal` bo'lsin. *Maksimal summani* qaytaring.


#### Example 1:

```
Input: nums = [1,4,3,2] 
Output: 4
Explanation: Barcha mumkin bo'lgan juftliklar (elementlarning tartibini e'tiborsiz qoldiring) quyidagilardir:
1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
Shunday qilib, maksimal mumkin bo'lgan summa 4 ga teng.
```

#### Example 2:

```
Input: nums = [6,2,6,5,1,2]
Output: 9
Explanation: Optimal juftlik (2, 1), (2, 5), (6, 6). min(2, 1) + min(2, 5) + min(6, 6) = 1 + 2 + 6 = 9.
```
 

#### Cheklovlar:

* <code>1 <= n <= 10<sup>4</sup></code>
* `nums.length == 2 * n`
* <code>-10<sup>4</sup> <= nums[i] <= 10<sup>4</sup></code>

> Hint-1
>
> Shubhasiz, **brute force** bu erda yordam bermaydi. Boshqa narsani o'ylab ko'ring, 1,2,3,4 kabi misollarni oling.

> Hint-2
> 
> Natijaga erishish uchun qanday qilib juftlik yaratasiz? Biror naqsh bo'lishi kerak.

> Hint-3
>
> Siz buni kuzatdingizmi - Minimal element maksimal elementni qurbon qilish natijasida natijaga qo'shiladi.

> Hint-4
>
> Hali ham juftliklarni topa olmayapsizmi? Massivni tartiblang va naqshni topishga harakat qiling.

```go
func arrayPairSum(nums []int) int {

	slices.Sort(nums)
    
	s := 0
	for i, v := range nums {
		if i%2 == 0 {
			s += v
		}
	}

	return s
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1154/)