# Two-pointer Technique - Scenario I
Oldingi bobda biz massivni takrorlash orqali ba'zi muammolarni hal qildik. Odatda, biz takrorlash uchun birinchi elementdan boshlanib, oxirgi elementda tugaydigan faqat bitta ko'rsatgichdan foydalanamiz. Biroq, ba'zida iteratsiyani bajarish uchun `bir vaqtning o'zida ikkita ko'rsatkich`dan foydalanishimiz kerak bo'lishi mumkin.

## Misol
Klassik muammodan boshlaylik:

> Massivdagi elementlarni teskari aylantiring.

Maqsad birinchi elementni oxiri bilan almashtirish, keyingi elementga o'tish va o'rta holatga kelguncha qayta-qayta almashishdir.

Takrorlashni amalga oshirish uchun bir vaqtning o'zida ikkita ko'rsatgichdan foydalanishimiz mumkin: `biri birinchi elementdan, ikkinchisi esa oxirgi elementdan boshlanadi`. Ikkala ko'rsatgich bir-biriga to'qnashguncha elementlarni almashtirishni davom eting.

Malumot uchun kod:
```cpp
void reverse(int *v, int N) {
    int i = 0;
    int j = N - 1;
    while (i < j) {
        swap(v[i], v[j]);
        i++;
        j--;
    }
}
```

## Xulosa
Xulosa qilib aytadigan bo'lsak, ikki nuqtali texnikadan foydalanishning odatiy stsenariylaridan biri bu siz xohlashingizdir

> Iterate the array from two ends to the middle.

Shunday qilib, ikki koʻrsatkichli texnikadan foydalanishingiz mumkin:

> Bir ko'rsatgich boshidan boshlanadi, ikkinchisi esa oxiridan boshlanadi.

Va shuni ta'kidlash kerakki, ushbu texnika ko'pincha `tartiblangan` massivda qo'llaniladi.

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1156/)