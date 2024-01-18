# Counting Sort

## What is Counting Sort?
`Counting sort` bu taqqoslashga asoslanmagan saralash algoritmi bo'lib, u kirish qiymatlarining cheklangan diapazoni mavjud bo'lganda yaxshi ishlaydi. Bu ayniqsa, kirish qiymarlari diapazoni saralanadigan elementlar soniga nisbatan kichik bo'lsa samarali bo'ladi. `Counting sort`ni asosiy g'oyasi kirish massividagi har bir elementning chastotasini hisoblash va bu ma'lumotlardan elementlarni to'g'ri tartiblangan joylarga joylashtirish uchun foydalanishdir.

## How does Counting Sort Algorithm work?
#### Step 1:
* Berilgan massivdan maksimal elementni toping.

![step 1](https://media.geeksforgeeks.org/wp-content/uploads/20230920182425/1.png)

#### Step 2:
* Uzunligi **max+1** bo'lgan **countArray[]** ni barcha elementlari **0** sifatida ishga tushuring. Ushbu massiv kirish massivining elementlarining paydo bo'lishini saqlash uchun ishlatiladi.

![step 2](https://media.geeksforgeeks.org/wp-content/uploads/20230920182436/2.png)

#### Step 3:
* **countArray[]** da kirish massivining har bir noyob elementi sonini tegishli indekslarida saqlang.

* **Misol uchun**: Kirish massividagi `2` elementining soni `2` ga teng. Demak, **countArray[]** da **2**-indeksda **2** ni saqlang. Xuddi shunday, kirish massividagi `5` elementining soni `1` ga teng, shuning uchun **countArray[]** da **5**-indeksda **1** ni saqlang.

![step 3](https://media.geeksforgeeks.org/wp-content/uploads/20230922132754/3.png)

#### Step 4:
* Endi `countArray[i] = countArray[i - 1] + countArray[i]` ni bajarib, **countArray[]** **elementlarining yig'indisi** yoki **prefiks yig'indisi**ni saqlang. Bu kirish massivining elementlarini chiqish massividagi to'g'ri indeksga joylashtirishga yordam beradi.

![step 4](https://media.geeksforgeeks.org/wp-content/uploads/20230920182646/4.png)

#### Step 5:
* Kirish massivining oxiridan boshlab takrorlang, chunki kirish massivini oxiridan o'tish teng elementlar tartibini saqlaydi, natijada bu tartiblash algoritmini **barqaror** qiladi.

* `outputArray[countArray[inputArray[i]]-1] = inputArray[i]`, **outputArray**ni yangilang.

* Shuningdek, `countArray[inputArray[i]] = countArray[inputArray[i]]--` **countArray**ni ham yangilang.

![step 5](https://media.geeksforgeeks.org/wp-content/uploads/20230920182656/5.png)

#### Keling birma bir ko'rib chiqamiz.
#### Step 6. i = 6 uchun:
* `outputArray[countArray[inputArray[6]]-1] = inputArray[6]`, **outputArray**ni yangilang.\
Shuningdek, `countArray[inputArray[6]] = countArray[inputArray[6]]--` **countArray**ni ham yangilang.

![step 6. i = 6](https://media.geeksforgeeks.org/wp-content/uploads/20230920182724/6.png)

#### Step 7. i = 5 uchun:
* `outputArray[countArray[inputArray[5]]-1] = inputArray[5]`, **outputArray**ni yangilang.\
Shuningdek, `countArray[inputArray[5]] = countArray[inputArray[5]]--` **countArray**ni ham yangilang.

![step 7. i = 5](https://media.geeksforgeeks.org/wp-content/uploads/20230920182741/7.png)

#### Step 8. i = 4 uchun:
* `outputArray[countArray[inputArray[4]]-1] = inputArray[4]`, **outputArray**ni yangilang.\
Shuningdek, `countArray[inputArray[4]] = countArray[inputArray[4]]--` **countArray**ni ham yangilang.

![step 8. i = 4](https://media.geeksforgeeks.org/wp-content/uploads/20230920182752/8.png)

#### Step 9. i = 3 uchun:
* `outputArray[countArray[inputArray[3]]-1] = inputArray[3]`, **outputArray**ni yangilang.\
Shuningdek, `countArray[inputArray[3]] = countArray[inputArray[3]]--` **countArray**ni ham yangilang.

![step 9. i = 3](https://media.geeksforgeeks.org/wp-content/uploads/20230920182807/9.png)

#### Step 10. i = 2 uchun:
* `outputArray[countArray[inputArray[2]]-1] = inputArray[2]`, **outputArray**ni yangilang.\
Shuningdek, `countArray[inputArray[2]] = countArray[inputArray[2]]--` **countArray**ni ham yangilang.

![step 10. i = 2](https://media.geeksforgeeks.org/wp-content/uploads/20230920182827/10.png)

#### Step 11. i = 1 uchun:
* `outputArray[countArray[inputArray[1]]-1] = inputArray[1]`, **outputArray**ni yangilang.\
Shuningdek, `countArray[inputArray[1]] = countArray[inputArray[1]]--` **countArray**ni ham yangilang.

![step 11. i = 1](https://media.geeksforgeeks.org/wp-content/uploads/20230920182855/11.png)

#### Step 12. i = 0 uchun:
* `outputArray[countArray[inputArray[0]]-1] = inputArray[0]`, **outputArray**ni yangilang.\
Shuningdek, `countArray[inputArray[0]] = countArray[inputArray[0]]--` **countArray**ni ham yangilang.

![step 12. i = 0](https://media.geeksforgeeks.org/wp-content/uploads/20230920182910/12.png)

## Counting Sort Algorithm
* **max(inputArray[])+1** oʻlchamdagi **countArray[]** yordamchi massivini eʼlon qiling va uni *0* bilan ishga tushiring.
* **inputArray[]** massivini aylantiring va **inputArray[]** ning har bir elementini **countArray[]** massivi indeksi sifatida ko'rsating, ya'ni **0 <= i < N** uchun **countArray[inputArray[i]]++** ni bajaring.
* **inputArray[]** massivining har bir indeksidagi prefiks summasini hisoblang.
* **N** o'lchamdagi **outputArray[]** massivi yarating.
* **inputArray[]** qatorini oxiridan aylantiring va `outputArray[countArray[inputArra[i]]-1] = inputArray[i]` ni yangilang. Shuningdek, ``countArray[inputArray[i]] = countArray[inputArray[i]]--`` ni yangilang.

Quyida yuqoridagi algoritmni amalga oshirish ko'rsatilgan:

```go
func countingSort(nums []int) []int {

    // Massivdan maksimal elementni topish uchun dastlab birinchi elementni max deb olamiz
    max := nums[0]

    // Massivni aylantirib, massivning maksimal elementini topamiz
    for _, num := range nums {
        if num > max {
            max = num
        }
    }

    // Massiv uzunligi max+1 bo'lgan countArray yaratamiz
    countArray := make([]int, max+1)

    // Massivni aylantirib, countArray elementlarini hisoblaymiz
    for _, num := range nums {
        countArray[num]++
    }

    // countArray elementlarining prefiks summasini hisoblaymiz
    for i := 1; i < len(countArray); i++ {
        countArray[i] += countArray[i-1]
    }

    // N o'lchamdagi outputArray yaratamiz
    outputArray := make([]int, len(nums))

    // inputArray qatorini oxiridan aylantiramiz va outputArray elementlarini joylashtiramiz
    for i := len(nums) - 1; i >= 0; i-- {
        outputArray[countArray[nums[i]]-1] = nums[i]
        countArray[nums[i]]--
    }

    return outputArray  
}
```
Kodni ishlatib ko'rish uchun [playground](https://go.dev/play/p/_6rPxg1hw5x)dan foydalaning.

## Complexity Analysis of Counting Sort:
* **Time Complexity** `O(N+K)`, bu yerda **N** - massiv uzunligi, **K** - massivning maksimal elementi.
    * Best Case: `O(N+K)`
    * Average Case: `O(N+K)`
    * Worst Case: `O(N+K)`
* **Space Complexity**: `O(N+K)`, bu yerda **N** - massiv uzunligi, **K** - massivning maksimal elementi.

## Advantage of Counting Sort:
* Counting sort agar kiritish diapazoni kirish soniga teng bo'lsa, odatda **merge sort, quick sort** kabi barcha taqqoslashga asoslangan tartiblash algoritmlariga qaraganda tezroq ishlaydi.
* Counting sortni kodlash oson
* Counting sort barqaror algoritmdir.

## Disadvantage of Counting Sort:
* Counting sort **decimal** qiymatlarda ishlamaydi.
* Saralash kerak bo'lgan qiymatlar oralig'i juda katta bo'lsa, `counting sort` samarasiz bo'ladi.
* Counting sort joyida tartiblash([in-place](https://en.wikipedia.org/wiki/In-place_algorithm)) algoritmi emas, u massiv elementlarini saralash uchun qo'shimcha joydan foydalanadi.

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/counting-sort/)