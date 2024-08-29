# Remove Nth Node From End of List

Bog'langan ro'yxatning `head`i berilgan, ro'yxat oxiridan `n-tugunni` olib tashlang va `head`ni qaytaring.

## Example 1:
![Remove Nth Node From End of List](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)

> **Input**: head = [1,2,3,4,5], n = 2 \
> **Output**: [1,2,3,5]

## Example 2:
> **Input**: head = [1], n = 1 \
> **Output**: []

## Example 3:
> **Input**: head = [1,2], n = 1
> **Output**: [1]

## Constraints:

* Ro'yxatdagi tugunlar soni `sz`.
* `1 <= sz <= 30`
* `0 <= Node.val <= 100`
* `1 <= n <= sz`

**Follow up**: Buni bir o'tishda amalga oshira olasizmi?

## My Solution

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next:head}
    l, r := dummy, head
    
    for n > 0 && r != nil {
        r = r.Next
        n--
    }
    
    for r != nil {
        r = r.Next
        l = l.Next
    }
    
    l.Next = l.Next.Next
    
    return dummy.Next
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1296/)