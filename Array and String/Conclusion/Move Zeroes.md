# Move Zeroes

Butun massiv `nums` berilgan bo'lsa, `0`ga teng bo'lmagan elementlarning nisbiy tartibini saqlab, barcha 0 larni oxiriga ko'chiring.

**Shuni esda tutingki**, siz massivning nusxasini yaratmasdan buni joyida qilishingiz kerak.

#### Example 1:
```
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
```

#### Example 2:
```
Input: nums = [0]
Output: [0]
```

#### Cheklovlar:
* `1 <= nums.length <= 10^4`
* `-2^31 <= nums[i] <= 2^31 - 1`

**Kuzatish**: bajarilgan operatsiyalarning umumiy sonini minimallashtira olasizmi?

> Hint-1
>
> Joyida, biz qo'shimcha massiv uchun hech qanday joy ajratmasligimiz kerakligini anglatadi. Lekin biz mavjud massivni o'zgartirishga ruxsat berilgan. Biroq, birinchi qadam sifatida, qo'shimcha joydan foydalanadigan yechimni topishga harakat qiling. Bu muammo uchun ham avval muhokama qilingan g'oyani qo'shimcha massiv yordamida qo'llang va o'z joyidagi yechim oxir-oqibat ochiladi.

> Hint-2
> 
> **Ikki nuqtali** yondashuv bu erda foydali bo'lishi mumkin. Massivni takrorlash uchun bitta ko'rsatgich va massivning nolga teng bo'lmagan elementlarida ishlaydigan boshqa ko'rsatgich bo'lishi kerak.

```go
func moveZeroes(nums []int)  {
    k := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[k]=nums[i]
            k++
        }
    }
    if k != 0 {
      for i:=k; i < len(nums); i++ {
          nums[i]=0
      }
    }
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1174/)