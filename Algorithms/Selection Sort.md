# Selection Sort

`Selection sort` - oddiy va samarali tartiblash algoritmi boʻlib, roʻyxatning tartiblanmagan qismidan eng kichik (yoki eng katta) elementni qayta-qayta tanlab, uni roʻyxatning tartiblangan qismiga koʻchirish orqali ishlaydi.

Algoritm ro'yxatning tartiblanmagan qismidan eng kichik (yoki eng katta) elementni qayta-qayta tanlaydi va uni tartiblanmagan qismning birinchi elementi bilan almashtiradi. Bu jarayon qolgan saralanmagan qism uchun butun ro'yxat saralanmaguncha takrorlanadi.

## How does Selection Sort Algorithm work?

Misol sifatida quyidagi massivni ko'rib chiqamiz: `arr[] = {64, 25, 12, 22, 11}`

Birinchi o'tish:
* Saralangan massivdagi birinchi o'rin uchun butun massiv 0 dan 4 gacha bo'lgan indeksdan ketma-ket o'tkaziladi. Hozirda `64` saqlanadigan birinchi pozitsiya, butun massivni aylanib o'tgandan so'ng, `11` eng past qiymat ekanligi ayon bo'ladi.

* Shunday qilib, 64 ni 11 bilan almashtiring. Bir marta takrorlangandan so'ng, massivdagi eng kichik qiymat bo'lgan 11 tartiblangan ro'yxatning birinchi pozitsiyasida paydo bo'ladi.

![image](https://media.geeksforgeeks.org/wp-content/uploads/20230524115038/1.webp)

Ikkinchi o'tish:
* `25` mavjud bo'lgan ikkinchi pozitsiya uchun massivning qolgan qismini yana ketma-ketlikda aylantiring.

* Ketishdan so'ng biz `12` massivdagi ikkinchi eng past qiymat ekanligini va u massivda ikkinchi o'rinda paydo bo'lishi kerakligini aniqladik, shuning uchun bu qiymatlarni almashtiring.

![image](https://media.geeksforgeeks.org/wp-content/uploads/20230526165135/2.webp)

Uchinchi o'tish:
* Endi, uchinchi o'rin uchun, `25` mavjud bo'lgan joyda, yana massivning qolgan qismini kesib o'ting va massivdagi uchinchi eng kam qiymatni toping.
* Ketish paytida `22` uchinchi eng kam qiymat bo'lib chiqdi va u massivda uchinchi o'rinda paydo bo'lishi kerak, shuning uchun `22` ni uchinchi o'rindagi element bilan almashtiring.

![image](https://media.geeksforgeeks.org/wp-content/uploads/20230526165200/3.webp)

To'rtinchi o'tish:
* Xuddi shunday, to'rtinchi pozitsiya uchun massivning qolgan qismini kesib o'ting va massivdagi to'rtinchi eng kichik elementni toping.
* `25` 4-eng past qiymat bo'lgani uchun u to'rtinchi o'rinni egallaydi.

![image](https://media.geeksforgeeks.org/wp-content/uploads/20230526165244/4.webp)

Beshinchi o'tish:
* Nihoyat, massivda mavjud bo'lgan eng katta qiymat avtomatik ravishda massivning oxirgi pozitsiyasiga joylashtiriladi
* Olingan massiv tartiblangan massivdir.

![image](https://media.geeksforgeeks.org/wp-content/uploads/20230526165320/5.webp)

## Implementation of Selection Sort
```go
// Selection sort - massivni o'sish tartibida saralash uchun algoritm
// arr - saralash kerak bo'lgan massiv
func selectionSort(arr []int) []int {

    // Massivdan eng kiçik elementni topib, uni massivning pastki o'riniga almashtirish
    for i:=0; i<len(arr); i++{

        minIndex := i

        // Eng kichik elementi indexini topish
        for j:=i+1; j<len(arr); j++{
            if arr[j] < arr[minIndex]{
                minIndex = j
            }
        }

        // Eng kichik elementni almashtirish
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
    return arr
}
```

## Complexity Analysis of Selection Sort

#### Time Complexity
Selection sort vaqt murakkabligi `O(N^2)` ga teng, chunki ikkita ichki oʻrnatilgan tsikl mavjud:
* Massiv elementini birma-bir tanlash uchun bitta tsikl = `O(N)`
* Ushbu elementni boshqa massiv elementi bilan solishtirish uchun yana bir tsikl = `O(N)`
* Shuning uchun umumiy murakkablik =` O(N) * O(N) = O(N*N) = O(N^2)`

#### Space Complexity
* `O(1)` qo'shimcha xotira sifatida faqat massivdagi ikkita qiymatni almashtirishda vaqtinchalik o'zgaruvchilar uchun ishlatiladi. Selection sort hech qachon `O(N)` almashinuvidan ko'proq narsani qilmaydi va xotirani yozish qimmat bo'lganda foydali bo'lishi mumkin.

## Advantages of Selection Sort Algorithm
* Oddiy va tushunish oson. 
* Kichik ma'lumotlar to'plamlari bilan yaxshi ishlaydi.

## Disadvantages of the Selection Sort Algorithm
* Seelction sort eng yomon va o'rtacha holatda `O(n^2)` vaqt murakkabligiga ega.
* Katta ma'lumotlar to'plamlarida yaxshi ishlamaydi.
* Teng kalitli elementlarning nisbiy tartibini saqlamaydi, bu barqaror emasligini anglatadi.

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/selection-sort/)