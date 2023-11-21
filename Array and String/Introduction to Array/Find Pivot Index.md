# Find Pivot Index

Butun sonlar `nums` massivi berilgan bo'lsa, ushbu massivning `pivot indeksi`ni hisoblang.

`Pivot indeksi` indeksning chap tomonidagi barcha raqamlarning yig'indisi indeksning o'ng tomonidagi barcha raqamlarning yig'indisiga teng bo'lgan indeksdir.

Agar indeks massivning chap chetida bo'lsa, chap yig'indi `0` ga teng, chunki chap tomonda elementlar yo'q. Bu massivning o'ng chetiga ham tegishli.

`Eng chap burchak indeksini` qaytaring. Agar bunday indeks mavjud bo'lmasa, `-1` ni qaytaring.

#### Example 1:
```
Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
Pivot indeksi 3 ga teng.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11
```

#### Example 2:
```
Input: nums = [1,2,3]
Output: -1
Explanation:
Muammo bayonotidagi shartlarni qondiradigan indeks yo'q.
```

#### Example 3:

```
Input: nums = [2,1,-1]
Output: 0
Explanation:
Pivot indeksi 0 ga teng.
Left sum = 0 (0 indeksining chap tomonida hech qanday element yo'q)
Right sum = nums[1] + nums[2] = 1 + -1 = 0
```

#### Cheklovlar:

* `1 <= nums.length <= 104`

* `-1000 <= nums[i] <= 1000`


> Eslatma: Bu savol 1991 yilgi savol bilan bir xil: [link](https://leetcode.com/problems/find-the-middle-index-in-array/)

> Hint-1: 
> 
> sumLeft massivi yarating, bunda sumLeft[i] i indeksining chap tomonidagi barcha raqamlar yig'indisidir.

> Hint-2:
>
> SumRight massivi yarating, bunda sumRight[i] i indeksining o'ng tomonidagi barcha raqamlar yig'indisidir.

> Hint-3: 
>
> Har bir i indeksi uchun sumLeft[i] sumRight[i] ga teng ekanligini tekshiring. Agar topilmasam, -1 ni qaytaring.

```go
func pivotIndex(nums []int) int {
	l, s := 0, 0

	for _, v := range nums {
		s += v
	}

	for i, _ := range nums {
		if l == s-l-nums[i] {
			return i
		}
		l += nums[i]
	}

	return -1
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/201/introduction-to-array/1144/)