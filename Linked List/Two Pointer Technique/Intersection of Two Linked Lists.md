# Intersection of Two Linked Lists

Ikkita bog'langan ro'yxatninig boshlari `headA` va `headB` berilgan, *ikkita ro'yxat kesishgan tugunni* qaytaring. Agar ikkita bog'langan ro'yxatning kesishishi bo'lmasa, `nullni` qaytaring.

Masalan, quyidagi ikkita bog'langan ro'yxat `c1` tugunida kesishishni boshlaydi:

![Intersection of linked lists](https://assets.leetcode.com/uploads/2021/03/05/160_statement.png)

Sinov holatlari shunday yaratilganki, butun bog'langan tuzilmaning hech bir joyida tsikllar bo'lmaydi.

**E'tibor bering**, bog'langan ro'yxatlar funktsiya qaytgandan keyin **asl tuzilishini saqlab qolishlari** kerak.

**Maxsus hakam:**
**Sudyaga** kiritilgan ma'lumotlar quyidagicha beriladi (dasturingizga ushbu ma'lumotlar berilmagan):

- `intersectVal` - kesishish sodir bo'lgan tugunning qiymati. Agar kesishgan tugun bo'lmasa, bu `0` ga teng.
- `listA` - birinchi bog'langan ro'yxat.
- `listB` - ikkinchi bog'langan ro'yxat.
- `skipA` - kesishgan tugunga o'tish uchun `A` ro'yxatida oldinga o'tish uchun tugunlar soni (boshdan boshlab).
- `skipB` - kesishgan tugunga o'tish uchun `B` ro'yxatida oldinga o'tish uchun tugunlar soni (boshdan boshlab).
Keyin sudya ushbu ma'lumotlar asosida bog'langan tuzilmani yaratadi va ikkita boshni, `headA` va `headB`ni dasturingizga uzatadi. Agar siz kesishgan tugunni to'g'ri qaytarsangiz, sizning yechimingiz `qabul qilinadi`.

## Example 1:
![Intersection of Linked Lists](https://assets.leetcode.com/uploads/2021/03/05/160_example_1_1.png)

> **Input**: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3 \
> **Output**: `8` da kesishgan \
> **Explanation**: Kesilgan tugunning qiymati 8 ga teng (ikkita ro'yxat kesishsa, bu 0 bo'lmasligi kerakligiga e'tibor bering). A ning boshidan [4,1,8,4,5] deb o'qiladi. B ning boshidan [5,6,1,8,4,5] deb o'qiladi. A da kesishgan tugundan oldin 2 ta tugun mavjud; B ning kesishgan tugunidan oldin 3 ta tugun mavjud. \
E'tibor bering, kesishgan tugunning qiymati 1 emas, chunki A va B da 1 qiymatiga ega bo'lgan tugunlar (A va B dagi 3-tugun) turli tugun havolalari hisoblanadi. Boshqacha qilib aytganda, ular xotiradagi ikki xil joyni ko'rsatadi, A va B da 8 qiymatiga ega bo'lgan tugunlar (Ada 3-tugun va Bda 4-tugun) xotiradagi bir xil joyni ko'rsatadi.

## Example 2:
![Intersection of Linked Lists](https://assets.leetcode.com/uploads/2021/03/05/160_example_2.png)

> **Input**: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1 \
> **Output**: `2` da kesishgan \
> **Explanation**: Kesilgan tugunning qiymati 2 ga teng (ikkita ro'yxat kesishsa, bu 0 bo'lmasligi kerakligiga e'tibor bering). A ning boshidan [1,9,1,2,4] deb o'qiladi. B ning boshidan [3,2,4] deb o'qiladi. A da kesishgan tugun oldidan 3 ta tugun bor; B ning kesishgan tugunidan oldin 1 ta tugun mavjud.

## Example 3:
![Intersection of Linked Lists](https://assets.leetcode.com/uploads/2021/03/05/160_example_3.png)

> **Input**: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2 \
> **Output**: Kesishmagan \
> **Explanation**: A ning boshidan [2,6,4] deb o'qiladi. B ning boshidan [1,5] deb o'qiladi. Ikkala ro'yxat kesishmaganligi sababli, intersectVal 0 bo'lishi kerak, skipA va skipB esa ixtiyoriy qiymatlar bo'lishi mumkin. \
Izoh: Ikki ro'yxat kesishmaydi, shuning uchun nullni qaytaring.

## Constraints:

* `A` ro'yxatidagi tugunlar soni `m` da.
* `B` ro'yxatidagi tugunlar soni `n` da.
* <code>1 <= m, n <= 3 * 10<sup>4</sup></code>
* <code>1 <= Node.val <= 10<sup>5</sup></code>
* `0 <= skipA < m`
* `0 <= skipB < n`
* `A` va `B` ro'yxatlar kesishmasa `intersectVal` `0`.
* `A` va `B` ro'yxatlar kesishsa `intersectVal == listA[skipA] == listB[skipB]`

**Follow up**: `O (m + n)` vaqtida ishlaydigan va faqat `O (1)` xotiradan foydalanadigan yechim yoza olasizmi?

## My Solutions

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	listMap := make(map[*ListNode]struct{}, 0)
    curr := headA

    for curr != nil {
        listMap[curr] = struct{}{}
        curr = curr.Next
    }

    curr = headB
    for curr != nil {
        if _, ok := listMap[curr]; ok {
            return curr
        }

        curr = curr.Next
    }

    return nil
}
```

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB

    for nodeA != nodeB {
        nodeA = nodeA.Next
        nodeB = nodeB.Next

        if nodeA == nil && nodeB != nil {
            nodeA = headB
        } else if nodeA != nil && nodeB == nil {
            nodeB = headA
        }
    }

    return nodeA
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1215/)