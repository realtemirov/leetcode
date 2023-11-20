# Immutable String - Problems & Solutions

Sevimli tilingizdagi satr o'zgarmas yoki o'zgarmasligini oldingi maqolada bilishingiz kerak. Agar satr o'zgarmas bo'lsa, u ba'zi muammolarni keltirib chiqaradi. Umid qilamizki, biz ham oxirida yechimni taqdim etamiz.

## O'zgartirish operatsiyasi

Shubhasiz, oʻzgarmas qatorni oʻzgartirib boʻlmaydi. Belgilardan faqat bittasini o'zgartirmoqchi bo'lsangiz, yangi qator yaratishingiz kerak.

### Java-da satrlarni birlashtirishdan ehtiyot bo'ling

`Satrlarni birlashtirishda` juda ehtiyot bo'lishingiz kerak. For tsiklida bir necha marta satrlarni birlashtirish misolini ko'rib chiqaylik:

```cpp
#include <iostream>

int main() {
    string s = "";
    int n = 10000;
    for (int i = 0; i < n; i++) {
        s += "hello";
    }
}
```
[C++ Playground](https://leetcode.com/playground/otMGzZHi)


```java
// "static void main" must be defined in a public class.
public class Main {
    public static void main(String[] args) {
        String s = "";
        int n = 10000;
        for (int i = 0; i < n; i++) {
            s += "hello";
        }
    }
}
```
[Java Playground](https://leetcode.com/playground/2q2Z2J6K)

Java uchun satrlarni birlashtirish qanchalik sekin ekanligiga e'tibor bering? Boshqa tomondan, C ++ da sezilarli ishlash ta'siri yo'q.

Java-da, satr `o'zgarmas` bo'lgani uchun, birlashtirish birinchi navbatda yangi satr uchun etarli joy ajratish, eski satrdan tarkibni nusxalash va yangi satrga qo'shish orqali ishlaydi.

Shunday qilib, umumiy vaqt murakkabligi quyidagicha bo'ladi:

> 5 + 5 × 2 + 5 × 3 + … + 5 × n \
> = 5 × (1 + 2 + 3 + … + n) \
> = 5 × n × (n + 1) / 2,

Bu <code>O(n<sup>2</sup>)</code>.   

### Solutions

Agar siz satringiz o'zgaruvchan bo'lishini istasangiz, ba'zi almashtirishlar mavjud:

### 1. Agar satringiz oʻzgaruvchan boʻlishini istasangiz, uni belgilar massiviga oʻzgartirishingiz mumkin.

```java
// "static void main" must be defined in a public class.
public class Main {
    public static void main(String[] args) {
        String s = "Hello World";
        char[] str = s.toCharArray();
        str[5] = ',';
        System.out.println(str);
    }
}
```
[Java Playground](https://leetcode.com/playground/kqZabUk6)

2. Agar satrlarni tez-tez birlashtirishga toʻgʻri kelsa, `StringBuilder` kabi boshqa maʼlumotlar tuzilmalaridan foydalangan maʼqulroq boʻladi. Quyidagi kod `O(n)` murakkabligida ishlaydi.

```java
// "static void main" must be defined in a public class.
public class Main {
    public static void main(String[] args) {
        int n = 10000;
        StringBuilder str = new StringBuilder();
        for (int i = 0; i < n; i++) {
            str.append("hello");
        }
        String s = str.toString();
    }
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/203/introduction-to-string/1184/)