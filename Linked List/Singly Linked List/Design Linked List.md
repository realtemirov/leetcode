# Design Linked List

O'zingizning linked listingizn implementatsiyasini ishlab chiqing. Singly yoki Doubly linked listni tanlashingiz mumkin.

Singly linked listdagi nodeda 2ta attribut bo'lishi kerak: `val` va `next`. `val` nodening qiymati va `next` keyingi nodega pointer/reference(havola).

Agar doubly linked listdan foydalanishni xohlasangiz, linked listdagi oldingi nodeni ko'rsatib turadigan `prev` attributeni qo'shishingiz kerak. Linked listdagi barcha node 0-indexdan boshlanadi deb faraz qiling.

`MyLinkedList` classini amalga oshiring:
- MyLinkedList() Initializes the MyLinkedList object.
- `int get(int index)` Linked listdagi index-nodeni qaytaradi. Agar index yaroqli bo'lmasa -1 qaytaradi.
- `void addAtHead(int val)` Linked listning birinchi elementidan oldin `val` qiymatli node qo'shadi. Qo'shgandan so'ng, yangi node linked listning birinchi nodei bo'ladi.
- `void addAtTail(int val)` Linked listning oxirgi elementi sifatida `val` qiymatli node qo'shadi.
- `void addAtIndex(int index, int val)` Linked listdagi `index-node`dan olding `val`qiymayli nodeni qo'shishadi. Agar index linked list uzunligiga teng bo'lsa, node linked list oxiridan qo'shiladi. Agar uzunligidan kattaroq bo'lsa node qo'shilmaydi.
- `void deleteAtIndex(int index)` Agar index yaroqli bo'lsa listdagi index-nodeni o'chiradi.

## Constraints:
- `0 <= index, val <= 1000`
- LinkedList qurib beruvchi kutubxonalardan foydalanmang,
- `get`, `addAtHead`, `addAtTail`, `addAtIndex` va `deleteAtIndex`larga ko'pi bilan `2000` ta qo'ng'iroq qilinadi.

## My Solution
```go
type Node struct {
    val int
    next *Node
}
type MyLinkedList struct {
    head *Node
    size int
}


func Constructor() MyLinkedList {
    return MyLinkedList{}
}


func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }

    curr := this.head 

    for i := 0; i < index; i++ {
        curr = curr.next
    }

    return curr.val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    node := &Node{val:val}

    if this.head != nil {
        node.next = this.head
        this.head = node
    } else {
        this.head = node
    }

    this.size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    node := &Node{val:val}

    var last *Node

    for i := this.head; i != nil; i = i.next {
        if i.next == nil {
            last = i
            break
        }
    }

    if last != nil {
        last.next = node
    } else {
        this.head = node
    }

    this.size++
}


func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index <0 || index > this.size {
        return
    }

    node := &Node{val:val}
    curr := this.head

    if index == 0 {
        node.next = this.head
        this.head = node
    } else {
        
        for i := 0; i < index-1; i++ {
            curr = curr.next
        }

        node.next = curr.next
        curr.next = node    
    }

    this.size++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.size {
        return
    }

    if index == 0 {
        this.head = this.head.next
    } else {
        curr := this.head

        for i := 0; i < index - 1; i++ {
            curr = curr.next
        }

        curr.next = curr.next.next
    }

    this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
```

© Leetcode [link](https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1290/)