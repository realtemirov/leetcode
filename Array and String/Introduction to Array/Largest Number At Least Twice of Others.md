# Largest Number At Least Twice of Others

Sizga eng katta butun son `yagona` bo'lgan sonlar qatori `nums` berilgan.

Massivdagi eng katta element massivdagi boshqa raqamlardan `kamida ikki barobar` ko'p yoki yo'qligini aniqlang. Agar shunday bo'lsa, eng katta element `indeksini` qaytaring yoki aks holda `-1` ni qaytaring.

#### Example 1:
```
Input: nums = [3,6,1,0]
Output: 1
Explanation: 6 - eng katta butun son.
X massividagi har bir boshqa son uchun 6 kamida x dan ikki baravar katta. 6 qiymatining indeksi 1 ga teng, shuning uchun biz 1 ni qaytaramiz. 
6 qiymatining indeksi 1 ga teng, shuning uchun biz 1 ni qaytaramiz.
```

#### Example 2:
```
Input: nums = [1,2,3,4]
Output: -1
Explanation: 4 3 qiymatidan ikki baravar kam, shuning uchun biz -1 ni qaytaramiz.
```

#### Constraints:

* `2 <= nums.length <= 50`
* `0 <= nums[i] <= 100`
* `nums`dagi eng katta element noyobdir.

> Hint1:
>
> Noyob eng katta element "m"ni topish uchun massivni skanerlang, uning "maxIndex" indeksini kuzatib boring. Massiv bo'ylab yana skanerlang. Agar biz `m < 2*x` bilan `x != m` ni topsak, biz `-1`ni qaytarishimiz kerak. Aks holda, biz "maxIndex" ni qaytarishimiz kerak.

```go
func dominantIndex(nums []int) int {
	m := nums[0]
	maxIndex := 0
	for i, v := range nums {
		if v > m {
			m = v
			maxIndex = i
		}
	}

	for _, v := range nums {
		if v != m && v*2 > m {
			return -1
		}
	}

	return maxIndex
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/201/introduction-to-array/1147/)