# Insertion Sort

## What is Insertion Sort?
Insertion Sort - bu oddiy tartiblash algoritmi bo'lib, u sizning qo'lingizda o'yin kartalarini saralash usuliga o'xshash ishlaydi. Massiv deyarli tartiblangan va tartiblanmagan qismga bo'lingan. Saralanmagan qismdan qiymatlar tanlanadi va tartiblangan qismning to'g'ri joyiga joylashtiriladi.

## Insertion Sort Algorithm
N o'lchamli massivni o'sish tartibida saralash uchun massiv bo'ylab takrorlang va joriy elementni (kalit) o'zidan oldingi bilan solishtiring, agar asosiy element avvalgisidan kichik bo'lsa, uni oldingi elementlar bilan solishtiring. Almashtirilgan element uchun joy ochish uchun kattaroq elementlarni bir pozitsiya yuqoriga siljiting. 

## How does Insertion Sort Algorithm work?
* Consider an example: arr[]: `{12, 11, 13, 5, 6}`

* Dastlab, massivning dastlabki ikki elementi kiritish tartibida solishtiriladi.

#### Step-1
> | `12`  | `11` | 13 | 5  | 6 |
> |-------|------|----|----|---|

* Bu yyerda 12 11 dan katta, shuning uchun ular o'sish tartibida emas va 12 o'zining to'g'ri joyida emas. Shunday qilib, 11 va 12 ni almashtiring.
* Shunday qilib, hozircha 11 tartiblangan pastki qatorda saqlanadi.

> | 11  | 12 | 13 | 5  | 6 |
> |-----|----|----|----|---|

#### Step-2
* Endi keyingi ikkita elementga o'ting va ularni solishtiring

> | 11  | `12` | `13` | 5  | 6 |
> |-----|------|------|----|---|

* Bu yerda 13 12 dan katta, shuning uchun ikkala element ham o'sish tartibida ko'rinadi, shuning uchun almashtirish sodir bo'lmaydi. 12, shuningdek, 11 bilan birga tartiblangan pastki qatorda saqlanadi

#### Step-3
* Endi tartiblangan kichik massivda ikkita element mavjud, ular 11 va 12
* Keyingi ikkita elementga o'tish - 13 va 5

> | 11  | 12 | `13` | `5`  | 6 |
> |-----|----|------|------|---|

* 5 va 13 ikkalasi ham o'z joyida yo'q, shuning uchun ularni almashtiring

> | 11  | 12 | `5` | `13` | 6 |
> |-----|----|-----|------|---|

* Almashtirilgandan so'ng, 12 va 5 elementlar tartiblanmaydi, shuning uchun yana almashtiriladi

> | 11  | `5` | `12` | 13 | 6 |
> |-----|-----|------|----|---|

* Bu yerda yana 11 va 5 tartiblashtirilmaydi, shuning uchun yana almashtiring

> | `5` | `11` | 12 | 13 | 6 |
> |-----|------|----|----|---|

* Bu yerda 5 to'g'ri holatda

#### Step-4
* Endi tartiblangan kichik massivda mavjud bo'lgan elementlar 5, 11 va 12 dir
* Keyingi ikkita elementga o'tish 13 va 6

> | 5 | 11 | 12 | `13` | `6` |
> |---|----|----|------|-----|

* Shubhasiz, ular tartiblanmagan, shuning uchun ikkalasini almashtirishni amalga oshiring

> | 5 | 11 | 12 | `6` | `13` |
> |---|----|----|-----|------|

* Endi 6 12 dan kichik, shuning uchun yana almashtiring

> | 5  | 11 | `6` | `12` | 13 |
> |----|----|-----|------|----|

* Bu erda, shuningdek, almashtirish 11 va 6 ni tartibsiz qiladi, shuning uchun yana almashtiring

> | 5 | `6` | `11` | 12 | 13 |
> |---|-----|------|----|----|

* Nihoyat, massiv to'liq tartiblangan.

![image](https://media.geeksforgeeks.org/wp-content/uploads/insertionsort.png)

## Implementation of Insertion Sort Algorithm

```go
// Insertion Sort - massivni o'sish tartibida saralash uchun algoritm
// arr - saralash kerak bo'lgan massiv
func insertionSort(arr []int) []int {
    

    for i:=1; i<len(arr); i++{

        // key - saralash kerak bo'lgan element
        key := arr[i]

        // j - key elementdan oldingi element
        j := i - 1

        // key elementdan oldingi elementdan katta bo'lsa
        // va j >= 0 bo'lsa, arr[j] elementni key elementga almashtiriladi
        for j >= 0 && arr[j] > key {

            // key elementdan oldingi elementni key elementga almashtirish
            arr[j+1] = arr[j]
            j -= 1
        }

        // key elementni key elementdan oldingi elementga almashtirish
        arr[j+1] = key
    }

    return arr
}
```

Kodni ishlatib ko'rish uchun playgrounddan foydalaning.

## Complexity Analysis of Insertion Sort

#### Time Complexity
* **Insertion Sort**ning **eng yomon** `vaqt` murakkabligi `O(N^2)`
* **Insertion Sort**ning **o'rtacha** `vaqt` murakkabligi `O(N^2)`
* **Eng yaxshi** holatning `vaqt` murakkabligi `O(N)` dir.

#### Space Complexity
* Insertion Sortning yordamchi `space` murakkabligi O(1) dir.

## Characteristics of Insertion Sort
* Bu algoritm oddiy amalga oshirish bilan eng oddiy algoritmlardan biridir
* Asosan, Insertion Sort kichik ma'lumotlar qiymatlari uchun samarali
* Insertion Sort moslashuvchan, ya'ni qisman saralangan ma'lumotlar to'plamlari uchun mos keladi.

## When to use Insertion Sort?
* Elementlar soni kichik bo'lsa, **Insertion Sort** qo'llaniladi. Bundan tashqari, kirish massivi deyarli tartiblangan va faqat bir nechta elementlar to'liq katta massivda noto'g'ri joylashtirilganda ham foydali bo'lishi mumkin.
* Insertion Sort - ([In-Place](https://en.wikipedia.org/wiki/In-place_algorithm)) algoritm, shuning uchun, qo'shimcha saqlash uchun joy talab qilmaydi.

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/insertion-sort/)