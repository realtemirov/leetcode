# Spiral Matrix

`m x n` matritsa mat berilgan bolsa, matritsaning barcha elementlarini spiral tartibda qaytaring..

#### Example 1:

![matrix](image-3.png)

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

#### Example 2:

![matrix](image-4.png)

```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```

#### Cheklovlar:

* `m == matrix.length`
* `n == matrix[i].length`
* `1 <= m, n <= 10`
* `-100 <= matrix[i][j] <= 100`

> Hint1: 
> 
> Ba'zi muammolar uchun eng yaxshi yo'l simulyatsiya uchun ba'zi algoritmlarni ishlab chiqishdir. Asosan, muammo bizdan nima qilishni talab qilayotganini taqlid qilishingiz kerak.

> Hint2:
>
> Biz chegara bo'ylab chegaradan o'tamiz va ichkariga harakat qilamiz. Bu asosiy operatsiya. Birinchi qator, oxirgi ustun, oxirgi qator, birinchi ustun, keyin esa 1 ga ichkariga o'tamiz va takrorlaymiz. Ana xolos. Bu bizga kerak bo'lgan barcha simulyatsiya.

> Hint3: 
>
> Indekslardan birida taraqqiyotni qachon o'zgartirmoqchi ekanligingizni o'ylab ko'ring. Agar [i, j] dan i ga o'tsangiz, xuddi shu ustunga o'tasiz. Xuddi shunday, j uchun qiymatlarni o'zgartirish orqali siz bir xil qatorga o'tasiz. Shuningdek, chegaraning oxirini kuzatib boring, shunda siz ichkariga qarab harakat qilishingiz va keyin takrorlashni davom ettirasiz. Bitta ustun yoki bitta qator kabi chekka holatlarni taqlid qilish har doim eng yaxshisidir.

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/202/introduction-to-2d-array/1168/)