# Two-pointer Technique - Scenario II
Ba'zan biz muammolarni hal qilish uchun `turli bosqichlarga ega ikkita ko'rsatgichdan` foydalanishimiz mumkin.

## Misol
Yana bir klassik muammodan boshlaylik:
> Massiv va qiymat berilgan boʻlsa, oʻsha qiymatning barcha nusxalarini joyida olib tashlang va yangi uzunlikni qaytaring.

Agar bizda kosmik murakkablik chegarasi bo'lmasa, bu osonroq bo'ladi. Javobni saqlash uchun yangi massivni ishga tushirishimiz mumkin. Asl massivni takrorlang va agar element berilgan maqsad qiymatiga teng bo'lmasa, elementni yangi massivga qo'shing.

Aslida, bu ikkita ko'rsatkichdan foydalanishga teng, biri asl massivni takrorlash uchun ishlatiladi, ikkinchisi esa har doim yangi massivning oxirgi holatiga ishora qiladi.

## Kosmik cheklovni qayta ko'rib chiqing
Endi bo'sh joy chegarasini qayta ko'rib chiqaylik.

Biz shunga o'xshash strategiyadan foydalanishimiz mumkin. Biz hali ham ikkita ko'rsatgichdan foydalanamiz: `biri hali ham iteratsiya uchun ishlatiladi, ikkinchisi esa har doim keyingi qo'shilish uchun pozitsiyani ko'rsatadi`.

Malumot uchun kod:
```cpp
int removeElement(vector<int>& nums, int val) {
    int k = 0;
    for (int i = 0; i < nums.size(); ++i) {
        if (nums[i] != val) {
            nums[k] = nums[i];
            ++k;
        }
    }
    return k;
}
```

Yuqoridagi misolda ikkita ko'rsatkichdan foydalanamiz, biri tezroq yuguruvchi `i` va yana biri sekinroq yuguruvchi `k`. `i` har safar bir qadam harakat qiladi, `k` esa bir qadam faqat yangi kerakli qiymat qo'shilsa.

## Xulosa
Bu sizga kerak bo'lganda ikki nuqtali texnikadan foydalanishning juda keng tarqalgan stsenariysi:

> Bir vaqtning o'zida bitta sekin yuguruvchi va bitta tez yuguruvchi.

Bunday muammolarni hal qilishning kaliti bu

> Ikkala ko'rsatkich uchun harakat strategiyasini aniqlang.

Oldingi stsenariyga o'xshab, ba'zan ikki ko'rsatkichli texnikani ishlatishdan oldin massivni `saralashingiz` kerak bo'lishi mumkin. Va harakat strategiyangizni aniqlash uchun sizga `ochko'z` fikr kerak bo'lishi mumkin.