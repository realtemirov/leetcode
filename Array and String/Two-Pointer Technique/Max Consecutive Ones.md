# Max Consecutive Ones

Ikkilik massiv `nums` berilgan bo'lsa, *massivdagi ketma-ket `1` ning maksimal sonini* qaytaring.

#### Example 1:
```
Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: Birinchi ikkita raqam yoki oxirgi uchta raqam ketma-ket 1 lardir. Ketma-ket 1 ning maksimal soni 3 ta.
```

#### Example 2:
```
Input: nums = [1,0,1,1,0,1]
Output: 2
```

#### Cheklovlar:

* `1 <= nums.length <= 10^5`
* `nums[i] 0 yoki 1 dir.`

> Hint-1
>
> Har qanday oynaga kelsak, siz ikkita narsa haqida o'ylashingiz kerak. Ulardan biri deraza uchun boshlang'ich nuqtadir. Yangi 1s oynasi boshlanganini qanday aniqlash mumkin? Keyingi qism bu oynaning tugash nuqtasini aniqlaydi. Mavjud oynaning tugash nuqtasini qanday aniqlash mumkin? Agar siz ushbu ikki narsani aniqlasangiz, ketma-ket oynalarning oynalarini aniqlay olasiz. Keyinchalik eng uzun oynani topish va o'lchamini qaytarish qoladi.

```go
func findMaxConsecutiveOnes(nums []int) int {
    
    max := 0
    min := 0

    for i:=0; i < len(nums); i++ {
        if nums[i] == 1 {
            min++
        } else {
            if min > max {
                max = min
            }
            min = 0
        }
    }
    if min > max {
        max = min
    }
    return max
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1301/)