# Delete Operation - Singly Linked List

Agar linked listda mavjud bo'lgan `cur` nodeni o'chirmoqchi bo'lsak, buni ikki bosqichqa qilishimiz mumkin.

1. `cur` nodedan oldingi `prev` node va keyingi `next` nodelarni topamiz.

![prev and next nodes](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/27/screen-shot-2018-04-26-at-203558.png)

2. `prev` nodeni `cur` nodedan keyingi nodega bog'laymiz.

![link prev to next](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/26/screen-shot-2018-04-26-at-203640.png)

Birinchi qadamda, `prev` va `next` nodelarni topishimiz kerak. `next` ni `cur`ning havolasidan foydalanib topish oson. Lekin biz `prev` ni topish uchun headdan boshlab o'tib chiqishimiz kerak, uzunligi *N* ga teng bo'lgan linked listda o'rtacha `O(N)` vaqt oladi. Shu sababli nodeni o'chirish vaqt murakkabligi `O(N)` bo'ladi.

## Misol

![an example](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/12/screen-shot-2018-04-12-at-152754.png)

Keling yuqoridagi singly linked listdan node-6 ni o'chirishga harakat qilib ko'ramiz.

1. head node dan boshlab oldingi `prev` node node-23 ni topganimizcha linked listni o'tib chiqamiz.

2. `prev` nodeni `next` nodega bog'laymiz.

![link prev to next](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/12/screen-shot-2018-04-12-at-154821.png)

node-6 endi bizning singly linked listimizda emas.

## Birinchi Nodeni o'chirish

Agar birinchi nodeni o'chirmoqchi bo'lsak bu biroz boshqacha bo'ladi.

Avval ta'kidlaganimizdek, linked listni ifodalash uchun bosh node sifatida `head node`dan foydalanamiz, Bosh nodeimiz pastda misoldagi qora node-23.

![black head node](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/19/screen-shot-2018-04-19-at-130024.png)

Agar birinchi nodeni o'chirmoqchi bo'lsak, shunchaki keyingi nodeni `head node`ga biriktiramiz. Ya'ni o'chirishdan keyin node-6 head node bo'ladi.

![delete head node](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/19/screen-shot-2018-04-19-at-130031.png)

Linked list head nodedan boshlanad, shuning uchun endi node-23 yo'q.

> Note
>
> Oxirgi tugunni o'chirish haqida nima deyish mumkin? Shunga oʻxshash strategiyadan foydalanishimiz mumkinmi?


© Leetcode [link](https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1289/)