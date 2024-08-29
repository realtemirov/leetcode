# Summary - Two-Pointer in Linked List

Bu erda biz bog'langan ro'yxatdagi ikki nuqtali muammoni hal qilish uchun shablonni taqdim etamiz.

```c++
// Initialize slow & fast pointers
ListNode* slow = head;
ListNode* fast = head;
/**
 * Change this condition to fit specific problem.
 * Attention: remember to avoid null-pointer error
 **/
while (slow && fast && fast->next) {
    slow = slow->next;          // move slow pointer one step each time
    fast = fast->next->next;    // move fast pointer two steps each time
    if (slow == fast) {         // change this condition to fit specific problem
        return true;
    }
}
return false;   // change return value to fit specific problem
```

## Maslahatlar
Bu biz massivda o'rgangan narsamizga o'xshaydi. Ammo bu qiyinroq va xatolarga moyil bo'lishi mumkin. E'tibor berishingiz kerak bo'lgan bir nechta narsalar mavjud:

**1. Hardoim keyingi fieldni chaqirishdan oldin har doim tugun `null` ekanligini tekshiring.**

`Null` tugunning keyingi tugunini olish `null koʻrsatkich xatosi`ga olib keladi. Misol uchun, biz `fast = fast.next.next` ishga tushirishdan oldin, biz `fast` va `fast.next` null emasligini tekshirishimiz kerak.

**2. Tsiklingizni yakuniy shartlarini diqqat bilan aniqlang.**
Yakuniy shartlaringiz cheksiz tsiklga olib kelmasligiga ishonch hosil qilish uchun bir nechta misollarni bajaring. Yakuniy shartlarni belgilashda birinchi maslahatimizni hisobga olishingiz kerak.

## Complexity Analysis
`Space complexity`ni tahlil qilish oson. Agar siz boshqa qo'shimcha bo'sh joysiz faqat ko'rsatgichlardan foydalansangiz, spacening murakkabligi `O(1)` bo'ladi. Biroq, vaqt murakkabligini tahlil qilish qiyinroq. Javobni olish uchun biz `tsiklimizni necha marta bajarishi`mizni tahlil qilishimiz kerak.

Oldingi `tsiklni topish` misolida, biz `tezroq ko'rsatgich`ni har safar 2 qadam va `sekinroq ko'rsatkich`ni har safar 1 qadam harakatlantiramiz deb faraz qilaylik.

1. Agar sikl boʻlmasa, tezkor koʻrsatkich bogʻlangan roʻyxatning oxiriga yetib borish uchun `N/2` marta oladi, bu yerda N bogʻlangan roʻyxatning uzunligi
2. Agar tsikl mavjud bo'lsa, tez ko'rsatkich sekinroq ko'rsatkichni ushlash uchun `M` marta kerak bo'ladi, bu erda M - ro'yxatdagi tsiklning uzunligi.

Shubhasiz, `M <= N`. Shunday qilib, biz tsiklni `N martagacha` bajaramiz. Va har bir loop uchun bizga faqat doimiy vaqt kerak. Demak, bu algoritmning vaqt murakkabligi jami `O(N)` ga teng.

Tahlil qilish mahoratingizni oshirish uchun boshqa muammolarni o'zingiz tahlil qiling. Turli xil shartlarni hisobga olishni unutmang. Agar barcha vaziyatlarni tahlil qilish qiyin bo'lsa, eng yomonini ko'rib chiqing.