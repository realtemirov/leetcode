# Linear Search

Barcha elementlarni birma-bir tekshirish orqali elementni topish usuli `chiziqli qidiruv(linear search)` algoritmi deb nomlanadi. **Eng yomon holat**da, chiziqli qidiruv butun massivni tekshirish bilan yakunlanadi. Shuning uchun chiziqli qidiruv uchun vaqt murakkabligi `O(N)`ga teng.

![linear search](https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/Figures/Array_Explore/Array_Search_1.png)


```go
func linearSearch(nums []int, target int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            return i
        }
    }
    return -1
}
```