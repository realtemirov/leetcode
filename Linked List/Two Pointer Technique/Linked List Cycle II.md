# Linked List Cycle II

`head` berilgan, bog'langan ro'yxatning `head` holda, *tsikl boshlangan tugunni* qaytaring. Agar tsikl bo'lmasa, `null`ni qaytaring.

Agar ro'yxatda `keyingi` ko'rsatkichni doimiy ravishda kuzatib borish orqali yana erishish mumkin bo'lgan ba'zi tugun bo'lsa, bog'langan ro'yxatda tsikl mavjud.
Ichkarida, `pos` tailning `keyingi` ko'rsatkichi ulangan tugun indeksini belgilash uchun ishlatiladi (**0-indekslangan**). Agar tsikl bo'lmasa `-1` bo'ladi.

**E'tibor bering, `pos` parametr sifatida o'tkazilmaydi.**

Bog'langan ro'yxatni **o'zgartirmang**.

## Example 1:
![Linked List Cycle](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

> **Input**: head = [3,2,0,-4], pos = 1 \
> **Output**: tail 1-tugun indexiga ulanadi \
> **Explanation**: Bog'langan ro'yxatda tail ikkinchi tugunga ulanadigan tsikl mavjud.

## Example 2:
![Linked List Cycle](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)

> **Input**: head = [1,2], pos = 0 \
> **Output**: tail 0-tugun indexiga ulanadi \
> **Explanation**: Bog'langan ro'yxatda tail birinchi tugunga ulanadigan tsikl mavjud.

## Example 3:
![Linked List Cycle](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)

> **Input**: head = [1], pos = -1 \
> **Output**: tsikl yo'q \
> **Explanation**: Bog'langan ro'yxatda tsikl mavjud emas.

## Constraints:

* Ro'yxatdagi tugunlar soni <code>[0, 10<sup>4</sup>]</code> oralig'ida.
* <code>-10<sup>5</sup> <= Node.val <= 10<sup>5</sup></code>
* Bog'langan ro'yxatdagi `pos` `-1` yoki **valid** indeks.

**Follow up**: `O(1)` (ya'ni doimiy) xotiradan foydalanib, uni hal qila olasizmi?

## My Solution

```go
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	hasCycle := false

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

    for slow != head {
        slow = slow.Next
        head = head.Next
    }

	return head
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1214/)