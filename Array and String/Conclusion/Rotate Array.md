# Rotate Array

Butun sonli massiv `nums` berilgan bo'lsa, massivni `k` qadam bilan o'ngga aylantiring, bu erda `k` manfiy emas.

#### Example 1:
```
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
1 qadam o'ngga aylantiring: [7,1,2,3,4,5,6]
2 qadam o'ngga aylantiring: [6,7,1,2,3,4,5]
3 qadam o'ngga aylantiring: [5,6,7,1,2,3,4]
```

#### Example 2:
```
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation: 
1 qadam o'ngga aylantiring: [99,-1,-100,3]
2 qadam o'ngga aylantiring: [3,99,-1,-100]
```

#### Cheklovlar

* `1 <= nums.length <= 10^5`
* `-2^31 <= nums[i] <= 2^31 - 1`
* `0 <= k <= 10^5`

**Kuzatish:**
* Iloji boricha ko'proq yechim topishga harakat qiling. Ushbu muammoni hal qilishning kamida `uchta` yo'li mavjud.
* Buni `O (1)` qo'shimcha joy bilan joyida qila olasizmi?

> Hint-1
>
> Eng oson yechim qo'shimcha xotiradan foydalanishi mumkin va bu juda yaxshi.

> Hint-2
>
> Haqiqiy hiyla bu muammoni qo'shimcha xotiradan foydalanmasdan hal qilishga urinayotganda paydo bo'ladi. Bu elementlarni harakatlantirish uchun qandaydir tarzda asl massivdan foydalanish kerakligini anglatadi. Endi biz har bir elementni asl joyiga joylashtirishimiz va uning atrofidagi barcha elementlarni sozlash uchun siljitishimiz mumkin, chunki bu juda qimmatga tushadi va kattaroq kirish massivlarida vaqt tugashi mumkin.

> Hint-3
>
> Bir fikr chizig'i kerakli natijaga erishish uchun massivni (yoki uning qismlarini) teskari aylantirishga asoslangan. Qayta tiklash misol yordamida bizga qanday yordam berishi mumkinligini o'ylab ko'ring.

> Hint-4
>
> Boshqa fikrlash yo'nalishi biroz murakkab, lekin aslida u har bir elementni asl holatiga joylashtirish g'oyasiga asoslanadi va elementni dastlab shu holatda kuzatib boradi. Asosan, har bir qadamda biz elementni o'z joyiga joylashtiramiz va u erda yoki qo'shimcha o'zgaruvchiga yozilgan elementni kuzatib boramiz. Biz buni bitta chiziqli o'tishda qila olmaymiz va bu erda g'oya elementlar orasidagi tsiklik bog'liqlikka asoslangan.

```go
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1182/)