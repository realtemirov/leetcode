# Introduction to String

Satr aslida `unicode belgilar` massividir. Massivda biz ishlatgan deyarli barcha amallarni bajarishingiz mumkin. Siz buni o'zingiz sinab ko'rishingiz mumkin.

Biroq, ba'zi farqlar mavjud. Ushbu maqolada biz string bilan ishlashda bilishingiz kerak bo'lgan ba'zilarini ko'rib chiqamiz. Bu xususiyatlar bir tildan boshqasiga juda farq qilishi mumkin.

## Compare Function
String o'zining `solishtirish funksiyasi`ga ega (sizga quyidagi kodda solishtirish funksiyasidan foydalanishni ko'rsatamiz).
Biroq, muammo bor:

> Ikki qatorni solishtirish uchun "==" dan foydalana olamizmi?

Bu savolga javobga bog'liq:
> Til operatorni haddan tashqari yuklashni qo'llab-quvvatlaydimi?

* Agar javobingiz `ha` boʻlsa (masalan, C++), ikkita satrni solishtirish uchun "==" dan foydalanishimiz `mumkin`.
* Agar javob `yo'q` bo'lsa (Java kabi) bo'lsa, biz ikkita satrni taqqoslash uchun "==" dan `foydalanmasligimiz mumkin`. Biz "==" dan foydalansak, u bu ikki obyekt bir xil ob'ekt yoki yo'qligini solishtiradi.

Keling, quyidagi misolni bajaramiz va natijalarni solishtiramiz:

```cpp
#include <iostream>

int main() {
    string s1 = "Hello World";
    cout << "s1 is \"Hello World\"" << endl;
    string s2 = s1;
    cout << "s2 is initialized by s1" << endl;
    string s3(s1);
    cout << "s3 is initialized by s1" << endl;
    // compare by '=='
    cout << "Compared by '==':" << endl;
    cout << "s1 and \"Hello World\": " << (s1 == "Hello World") << endl;
    cout << "s1 and s2: " << (s1 == s2) << endl;
    cout << "s1 and s3: " << (s1 == s3) << endl;
    // compare by 'compare'
    cout << "Compared by 'compare':" << endl;
    cout << "s1 and \"Hello World\": " << !s1.compare("Hello World") << endl;
    cout << "s1 and s2: " << !s1.compare(s2) << endl;
    cout << "s1 and s3: " << !s1.compare(s3) << endl;
}
```
[C++ Playground](https://leetcode.com/playground/g4kbNuNZ)

## Immutable or Mutable

O'zgarmas(`immutable`) satrni ishga tushirgandan so'ng uning mazmunini o'zgartira olmasligingizni anglatadi.

1. Ba'zi tillarda (masalan, C++) satr `o'zgaruvchan`. Ya'ni, siz qatorni xuddi massivda qilganingiz kabi o'zgartirishingiz mumkin.

2. Ba'zi boshqa tillarda (Java kabi) string `o'zgarmasdir`. Bu xususiyat bir nechta muammolarni keltirib chiqaradi. Muammolar va echimlarni keyingi maqolada tasvirlab beramiz.

`O'zgartirish operatsiyasini sinab ko'rish` orqali sevimli tilingizdagi satr o'zgarmas yoki o'zgaruvchanligini aniqlashingiz mumkin. Mana bir misol:

```cpp
#include <iostream>

int main() {
    string s1 = "Hello World";
    s1[5] = ',';
    cout << s1 << endl;
}
```

## Qo'shimcha operatsiyalar

Massiv bilan solishtiring, biz satrda bajarishimiz mumkin bo'lgan qo'shimcha operatsiyalar mavjud. Mana bir nechta misollar:

```cpp
#include <iostream>

int main() {
    string s1 = "Hello World";
    // 1. concatenate
    s1 += "!";
    cout << s1 << endl;
    // 2. find
    cout << "The position of first 'o' is: " << s1.find('o') << endl;
    cout << "The position of last 'o' is: " << s1.rfind('o') << endl;
    // 3. get substr
    cout << s1.substr(6, 5) << endl;
}
```
[C++ Playground](https://leetcode.com/playground/bx9nUVsD)

Ushbu o'rnatilgan operatsiyalarning `vaqt murakkabligi`dan xabardor bo'lishingiz kerak.

Masalan, agar satr uzunligi `N` boʻlsa, qidiruv va pastki satr operatsiyasining vaqt murakkabligi `O(N)` ga teng.

Shuningdek, satr oʻzgarmas tillarda birlashtirish operatsiyasidan ehtiyot boʻlishingiz kerak (buni keyingi maqolada ham tushuntiramiz).

Yechimingiz uchun vaqt murakkabligini hisoblashda o'rnatilgan operatsiyalarning vaqt murakkabligini hisobga olishni hech qachon unutmang.

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/203/introduction-to-string/1158/)