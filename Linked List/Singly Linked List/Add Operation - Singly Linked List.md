# Add Operation - Singly Linked List

Agar biz berilgan `oldingi node`dan keyin yangi qiymat qo'oshmoqchi bo'lsak:

1. Berilgan qiymat bilan yangi `cur` nodeni initialize qilishimiz kerak.
![currently node](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/26/screen-shot-2018-04-25-at-163224.png)

2. `cur` nodening `next` fieldini `oldingi` nodedan keyingi nodega bog'lashimiz kerak.
![link nodes](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/26/screen-shot-2018-04-25-at-163234.png)

3. `prev` ning `next` fieldini `cur`ga bog'laymiz:
![link prev to cur](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/26/screen-shot-2018-04-25-at-163243.png)

Massivdan farqli ravishda, qo'shilgan elementdan keyin barcha elementlarni ko'chirishimiz shart emas. Shu sababli, agar sizda `prev`(oldingisi)ga bog'langan element bo'lsa,  `O(1)` vaqt murakkabligida linked listga yangi node qo'sha olasiz, bu juda samarali.

## Misol
![linked list example](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/12/screen-shot-2018-04-12-at-152754.png)

Keling, 6 qiymatli 2-nodedan keyin yangi 9 qiymatli node qo'shamiz.

Biz birinchi 9 qiymatli yangi node ishga tushiramiz. Keyin node-9 ni node-15 ga bog'laymiz. Va oxiri, node-6 ni node-9 ga bog'laymiz.

Qo'shganimizdan keyin bizning linked listimiz quyidagicha ko'rinadi.
![inserted linked list](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/12/screen-shot-2018-04-12-at-154238.png)

## Boshidan Node qo'shish

Bilamizki, biz butun listni ifodalash uchun bosh node `head node` dan foydalanamiz.

Demak, listni boshidan yangi qiymat qo'shganimizda `head` ni yangilash muhim.

1. Yangi `cur` nodeni ishga tushiramiz.
2. Yangi `cur` nodeimizni o'zimizning bosh node `head node`ga bog'laymiz.
3. `cur` nodeni head node ga biriktiramiz.

Misol uchun, yangi node-9 ni linked listni boshidan qo'shamiz.

1. 9 qiymatli yangi nodeni ishga tushiramiz va node-9 ni hozirgi head node node-23ga bog'laymiz.
![link cur to head node](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/19/screen-shot-2018-04-19-at-125118.png)
2. node-9 ni yangi head node qilib belgilaymiz.
![assign head node](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/19/screen-shot-2018-04-19-at-125350.png)

> Note
>
> Ro'yxat oxiriga yangi tugun qo'shish haqida nima deyish mumkin? Shunga o'xshash strategiyadan foydalanishimiz mumkinmi?

© Leetcode [link](https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1288/)