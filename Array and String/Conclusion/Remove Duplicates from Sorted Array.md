# Remove Duplicates from Sorted Array

**Kamaymaslik tartibida** `nums` tartiblangan butun sonli massivni hisobga olsak, har bir noyob element faqat **bir marta** paydo bo'lishi uchun dublikatlarni joyida olib tashlang. Elementlarning **nisbiy tartibi** **bir xil** saqlanishi kerak. Keyin *noyob elementlar sonini* nums bilan qaytaring.

`nums`ning noyob elementlari soni `k` bo'lishini ko'rib chiqing, qabul qilish uchun siz quyidagi amallarni bajarishingiz kerak:
* Massiv `nums`ni shunday o'zgartiringki, `nums`ning birinchi `k` elementi dastlab `nums`da mavjud bo'lgan tartibda noyob elementlarni o'z ichiga oladi. `nums`ning qolgan elementlari raqamlarning kattaligi kabi muhim emas.
* `k` ni qaytaring.

Sudya sizning yechimingizni quyidagi kod bilan sinab ko'radi:
```python
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```
Agar barcha tasdiqlar o'tib ketsa, sizning yechimingiz **qabul qilinadi**.

#### Example 1:
```
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Sizning funktsiyangiz k = 2 ni qaytarishi kerak, sonlarning birinchi ikkita elementi mos ravishda 1 va 2 bo'ladi.
Qaytarilgan k dan keyin nima qoldirganingiz muhim emas (shuning uchun ular pastki chiziq).
```

#### Example 2:
```
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Sizning funktsiyangiz k = 5 ni qaytarishi kerak, sonlarning birinchi besh elementi mos ravishda 0, 1, 2, 3 va 4.
Qaytarilgan k dan keyin nima qoldirganingiz muhim emas (shuning uchun ular pastki chiziq)..
```

#### Cheklovlar:

* `1 <= nums.length <= 3 * 10^4`
* `-100 <= nums[i] <= 100`
* `nums` is sorted in **non-decreasing** order.

> Hint-1
> 
> Ushbu muammoda e'tibor berish kerak bo'lgan asosiy nuqta - tartiblangan kirish massivi. Ikki nusxadagi elementlarga kelsak, berilgan massiv tartiblanganda ularning massivdagi joylashuvi qanday? Javob uchun yuqoridagi rasmga qarang. Agar biz elementlardan birining o'rnini bilsak, barcha takrorlanuvchi elementlarning joylashishini ham bilamizmi?
> ![image](https://assets.leetcode.com/uploads/2019/10/20/hint_rem_dup.png)

> Hint-2
>
> Biz massivni joyida o'zgartirishimiz kerak va yakuniy massivning o'lchami kirish massivining o'lchamidan kichikroq bo'lishi mumkin. Demak, bu yerda ikki nuqtali yondashuvdan foydalanishimiz kerak. Birinchisi, bu asl massivdagi joriy elementni, ikkinchisi esa faqat noyob elementlarni kuzatib boradi.

> Hint-3
> 
> Aslida, biror elementga duch kelganingizdan so'ng, siz uning dublikatlarini chetlab o'tib, keyingi noyob elementga o'tishingiz kerak.

```go
func removeDuplicates(nums []int) int {
	k := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k+1
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1173/)