# Two-Pointer in Linked List
Klassik muammodan boshlaylik:

> Bog'langan ro'yxat berilgan bo'lsa, unda tsikl bor yoki yo'qligini aniqlang.

`Xesh-jadval` yordamida yechim topib kelgan bo'lishingiz mumkin. Ammo `ikki nuqtali texnikadan` foydalangan holda samaraliroq yechim mavjud. Qolgan tarkibni o'qishdan oldin o'zingiz o'ylab ko'ring. 

Tasavvur qiling, u erda turli xil tezlikka ega ikkita yuguruvchi bor. Agar ular toʻgʻri yoʻlda yugurayotgan boʻlsa, tez yuguruvchi avval belgilangan manzilga yetib boradi. Biroq, agar ular aylana yo'lda yugurayotgan bo'lsa, tez yuguruvchi yugurishda davom etsa, sekin yuguruvchiga yetib oladi.

Bog'langan ro'yxatda biz har xil tezlikdagi ikkita ko'rsatgichdan foydalanganda aynan shu narsaga duch kelamiz:

1. Agar tsikl bo'lmasa, tezkor ko'rsatkich bog'langan ro'yxat oxirida to'xtaydi.
2. Agar tsikl mavjud bo'lsa, tezkor ko'rsatkich oxir-oqibat sekin ko'rsatkich bilan uchrashadi.

Shunday qilib, qolgan yagona muammo:

> Ikki ko'rsatkich uchun to'g'ri tezlik qanday bo'lishi kerak?

Tez ko'rsatkichni bir vaqtning o'zida ikki qadam siljitishda sekin ko'rsatkichni bir qadam siljitish xavfsiz  tanlovdir. Har bir iteratsiya uchun tezkor koʻrsatkich qoʻshimcha bir qadam harakat qiladi. Agar sikl uzunligi M boʻlsa, M marta takrorlangandan soʻng, tezkor koʻrsatgich albatta yana bir sikl harakatlanadi va sekin koʻrsatkichga yetib oladi.

> Boshqa tanlovlar haqida nima deyish mumkin? Ular ishlaydimi? Ular samaraliroq bo'ladimi?

© Leetcode [link](https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1211/)