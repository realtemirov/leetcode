# Linked List Cycle

`head` berilgan, bog'langan ro'yxatning `head` holda, bog'langan ro'yxatda tsikl bor yoki yo'qligini aniqlang.

Agar roʻyxatda `keyingi` koʻrsatkichni doimiy ravishda kuzatib borish orqali yana yetib borish mumkin boʻlgan tugun boʻlsa, bogʻlangan roʻyxatda tsikl mavjud. Ichkarida `pos` dumining `keyingi` koʻrsatkichi bogʻlangan tugun indeksini belgilash uchun ishlatiladi. 

**Eʼtibor bering, `pos` parametr sifatida oʻtkazilmaydi.**

Agar bog'langan ro'yxatda sikl bo'lsa, `true` qiymatini qaytaring. Aks holda, `false` qaytaring.


## Example 1:
![Linked List Cycle](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

> **Input**: head = [3,2,0,-4], pos = 1 \
> **Output**: true \
> **Explanation**: Bog'langan ro'yxatda tail 1-tugunga (0-indekslangan) ulanadigan tsikl mavjud.

## Example 2:
![Linked List Cycle](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)

> **Input**: head = [1,2], pos = 0 \
> **Output**: true \
> **Explanation**: Bog'langan ro'yxatda tail 0-tugunga (0-indekslangan) ulanadigan tsikl mavjud.

## Example 3:
![Linked List Cycle](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)

> **Input**: head = [1], pos = -1 \
> **Output**: false \
> **Explanation**: Bog'langan ro'yxatda tsikl mavjud emas.

## Constraints:

* Ro'yxatdagi tugunlar soni <code>[0, 10<sup>4</sup>]</code> oralig'ida.
* <code>-10<sup>5</sup> <= Node.val <= 10<sup>5</sup></code>
* Bog'langan ro'yxatdagi `pos` `-1` yoki **valid** indeks.

**Follow up**: `O(1)` (ya'ni doimiy) xotiradan foydalanib, uni hal qila olasizmi?

## My Solution

```go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        
        if fast == slow {
            return true
        }
    }
    
    return false
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1212/)