# Introduction - Singly Linked List

Singly linked listdagi har bir node nafaqat qiymatdan, balki `reference field`(havola maydoni)dan ham iborat. Shuningdek, singly linked list ketma-ketlikdagi nodelardan tashkil topgan. 

Bu yerda singly linked listni namunasi:

![Singly linked list](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/12/screen-shot-2018-04-12-at-152754.png)

Ko'k o'qlar singly linked listdagi nodelarni birgalikda qanday bog'langanini ko'rsatadi.

## Node structure

Bu yerda singly linked list nodeining odatiy ta'rifi:

### Golang
```go
type SinglyLinkedList struct {
	Value int
	Next  *SinglyLinkedList
}
```

### C++
```c++
// Definition for singly-linked list.
struct SinglyListNode {
    int val;
    SinglyListNode *next;
    SinglyListNode(int x) : val(x), next(NULL) {}
};
```

### Java
```java
// Definition for singly-linked list.
public class SinglyListNode {
    int val;
    SinglyListNode next;
    SinglyListNode(int x) { val = x; }
}
```

Ko'p hollarda biz butun list uchun `head node`dan foydalanamiz.

## Operations

Massivdagi farqli ravishda, biz linked listda o'zgarmas vaqtda tasodifiy elementlarga kira olmaymiz. Agar siz i-elementni olishni xohlasangiz, birma-bir bosh tugun(head)dan boshlab o'tib chiqishimiz kerak. Bu uzunligi *N*ga teng bo'lgan listda `index bo'yicha kirish`da o'rtacha xollarda `O(N)` oladi.

Misol uchun, yuqoridagi misolda, head node 23. 3-nodega kirishning birgina yo'li 2-node(6)ga o'tish uchun `head nodening next maydoni`dan foydalanish. Keyin biz 6 qiymatli nodening next maydoni bilan 3-nodega kira olamiz.

Siz index bo'yicha kirishda yomon performancega ega bo'lgan (arrayga bilan solishtirganda) linked list nima uchun foydali ekanligi haqida o'ylayotgan bo'lishingiz mumkin. Keyingi maqolalarda biz `insert` va `delete` operatsiyalarini tanishtiramiz va siz linked listning foydalarini tushunib yetasiz.

Shundan so'ng, biz o'zingizning linked listingizni ishlab chiqish uchun mashq taqdim etamiz.

© Leetcode [link](https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1287/)