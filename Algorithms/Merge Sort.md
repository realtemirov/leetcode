# Merge Sort

## What is Merge Sort?
`Merge Sort` massivni kichikroq kichik massivlarga bo'lish, har bir kichik massivni saralash va so'ngra saralangan pastki massivlarni oxirgi tartiblangan massivni hosil qilish uchun qayta birlashtirish orqali ishlaydigan tartiblash algoritmi sifatida aniqlanadi.

Oddiy qilib aytganda, birlashtirish tartiblash jarayoni massivni ikki yarmiga bo'lish, har bir yarmini tartiblash va keyin tartiblangan yarmini yana birlashtirishdir. Bu jarayon butun massiv saralanmaguncha takrorlanadi.

## How does Merge Sort work?
> **Merge Sort** - massivni keyingi boʻlinish boʻlmaguncha doimiy ravishda yarmiga boʻladigan rekursiv algoritmdir, yaʼni massivda faqat bitta element qoladi (bitta elementli massiv har doim tartiblanadi). Keyin tartiblangan pastki massivlar bitta tartiblangan massivga birlashtiriladi.

Massivni ko'rib chiqaylik []arr = {38, 27, 43, 10} \
Merge Sortni tushunish uchun quyidagi rasmlarga qarang.

* Dastlab massivni ikkita teng yarmiga bo'ling:

![step 1](https://media.geeksforgeeks.org/wp-content/uploads/20230530153635/img1drawio.png)

* Bu pastki massivlar yana ikkiga bo'linadi. Endi ular endi bo'linib bo'lmaydigan birlik uzunlikdagi massivga aylanadi va birlik uzunlikdagi massiv har doim tartiblanadi.

![step 2](https://media.geeksforgeeks.org/wp-content/uploads/20230530153654/img2drawio.png)

* Ushbu tartiblangan pastki qatorlar birlashtiriladi va biz kattaroq tartiblangan pastki qatorlarni olamiz.

![step 3](https://media.geeksforgeeks.org/wp-content/uploads/20230530153714/img3drawio.png)

* Bu birlashtirish jarayoni kichikroq pastki massivlardan tartiblangan massiv qurilguncha davom ettiriladi. 

![step 4](https://media.geeksforgeeks.org/wp-content/uploads/20230530153747/img4drawio.png)


## Merge Sort Algorithm
* `mergeSort(arr []int) []int`  funksiya yaratamiz, bu funksiya massivni saralash uchun ishlatiladi. 
* Massiv uzunligi 1 dan kichik yoki teng bo'lsa, to'liq massivni qaytaring, chunki massivdagi bitta elementni saralash mumkin emas. Keyin berilgan massivning o'rta nuqtasini toping.
* Massivning chap yarmida `mergeSort()` funktsiyasini rekursiv chaqiring va natijani `left` deb nomlangan o'zgaruvchiga saqlang.
* Shunga o'xshab, massivning o'ng yarmida `mergeSort()` funktsiyasini qayta chaqiring va natijani `right` deb nomlangan o'zgaruvchiga saqlang.
* Chap va o'ng massivlar bilan `merge()` funksiyasini kirish sifatida chaqiring va natijani qaytaring.
* `merge(left, right []int) []int` funktsiyasi chap va o'ng massivlarni argument sifatida qabul qiladi va ularni `for` sikllari va `if` shartlaridan foydalanib bitta massivda birlashtiradi.
* `merge()` funksiyasida `left` va `right` uzunligi teng bo'lgan `result` massivi yaratiladi. 
* `left` va `right` massivlarining boshlang'ich indekslari `i=0` va `j=0` sifatida olinadi. Ular `counter` sifatida ishlatiladi. Agar `left[i] > right[j]` bo'lsa `left[i]` `result`ga joylashtiring va `i` ni 1 ga oshiring. Aks holda `right[j]` `result`ga joylashtiring va `j` ni 1 ga oshiring.
* `left` va `right` massivlaridan biri tugaganda, qolgan massivning barcha elementlarini `result`ga joylashtiring.
* `result` massivini qaytaring.

#### Quyida Merge Sort kodini amalga oshirish ko'rsatilgan.
```go
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    mid := len(arr) / 2

    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])

    return merge(left, right)
}

func merge(left,right []int) []int {
    
    result := make([]int, len(left)+len(right))

    i, j := 0, 0

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result[i+j] = left[i]
            i++
        } else {
            result[i+j] = right[j]
            j++
        }
    }

    for i < len(left) {
        result[i+j] = left[i]
        i++
    }

    for j < len(right) {
        result[i+j] = right[j]
        j++
    }

    return result
}
```
Kodni ishlatib ko'rish uchun playgrounddan foydalaning.

* **Time Complexity** `O(N log(N))`, bu yerda **N** - massiv uzunligi.
* **Space Complexity**: `O(N)`, bu yerda **N** - massiv uzunligi. Merge Sortda barcha elementlar yordamchi massivga ko'chiriladi. Shuning uchun N yordamchi bo'sh joy talab qilinadi.

## Advantage of Merge Sort:
* **Barqarorlik:** Merge Sort barqaror tartiblash algoritmidir, ya'ni u kirish massividagi teng elementlarning nisbiy tartibini saqlaydi.
* **Kafolatlangan eng yomon ish unumdorligi:** Merge Sort eng yomon vaqt murakkabligi O(N logN) ga ega, ya'ni u hatto katta ma'lumotlar to'plamlarida ham yaxshi ishlaydi.
* **Parallellash mumkin:** Merge Sort tabiiy ravishda parallellashtiriladigan algoritmdir, ya'ni uni bir nechta protsessor yoki threadlardan foydalanish uchun osongina parallellashtirish mumkin.

## Disadvantage of Merge Sort:
* **Bo'shliqning murakkabligi:** Merge Sort saralash jarayonida birlashtirilgan kichik massivlarni saqlash uchun qo'shimcha xotira talab qiladi.
* Joyda emas: Merge Sort [In Place algoritmi](https://en.wikipedia.org/wiki/In-place_algorithm) emas, ya'ni tartiblangan ma'lumotlarni saqlash uchun qo'shimcha xotira talab qilinadi. Bu xotiradan foydalanish tashvish tug'diradigan ilovalarda kamchilik bo'lishi mumkin
* **Kichik ma'lumotlar to'plamlari uchun har doim ham maqbul emas:** Kichik ma'lumotlar to'plamlari uchun Merge Sort Insertion Sort kabi ba'zi boshqa tartiblash algoritmlariga qaraganda yuqori vaqt murakkabligiga ega. Bu juda kichik ma'lumotlar to'plamlari uchun sekinroq ishlashga olib kelishi mumkin.

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/merge-sort/)