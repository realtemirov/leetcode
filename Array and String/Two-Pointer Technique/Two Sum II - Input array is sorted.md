# Two Sum II - Input array is sorted

***Kamaymaydigan tartibda saralangan*** **1-indekslangan** butun sonlar qatori `numbers` berilgan, ikki raqamni toping, shunda ular maxsus `target` raqamga qo'shiladi.

Allaqachon kamaymaydigan tartibda tartiblangan 1 -indekslangan butun sonlar massivini hisobga olib , ikkita raqamni topingki, ular ma'lum bir raqamga qo'shiladi . Bu ikki raqam va qaerda bo'lsin .numberstargetnumbers[index1]numbers[index2]1 <= index1 < index2 < numbers.length

*Uzunligi 2 bo'lgan butun son massivi* `[index1, index2]` *sifatida **bitta tomonidan** qo'shilgan*, <code>indeks<sub>1</sub></code> va <code>indeks<sub>2</sub></code> *indekslarini qaytaring*.

Sinovlar shunday yaratilganki, **aynan bitta yechim mavjud** . Siz bir xil elementni ikki marta **ishlata olmaysiz** .

Sizning yechimingiz faqat doimiy qo'shimcha joydan foydalanishi kerak.


#### Example 1:

```
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: 2 va 7 yig'indisi 9. Shuning uchun indeks1 = 1, indeks2 = 2. Biz [1, 2] ni qaytaramiz.
```

#### Example 2:

```
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: 2 va 4 yig'indisi 6. Shuning uchun indeks1 = 1, indeks2 = 3. Biz [1, 3] ni qaytaramiz.
```

#### Example 3:

```
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: -1 va 0 yig'indisi -1 ga teng. Shuning uchun indeks1 = 1, indeks2 = 2. Biz [1, 2] ni qaytaramiz.
```
 

#### Cheklovlar:

* <code>2 <= numbers.length <= 3 * 10<sup>4</sup></code>
* `-1000 <= numbers[i] <= 1000`
* `numbers` **kamaymaydigan tartibda** tartiblangan.
* `-1000 <= target <= 1000`
* Sinovlar shunday yaratilganki, **aynan bitta yechim** mavjud.

```go
func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}
```

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1153/)