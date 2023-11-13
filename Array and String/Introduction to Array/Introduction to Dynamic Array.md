# Introduction to Dynamic Array

Oldingi maqolada aytib o'tganimizdek, array `o'zgarmas sig'imga` ega va biz uni ishga tushirishda massiv hajmini ko'rsatishimiz kerak. Ba'zan bu biroz noqulay va behuda bo'ladi.

Shu sababli, koʻpgina dasturlash tillari oʻrnatilgan `dynamic array`ni taklif qiladi, bu esa baribir tasodifiy kirish roʻyxati maʼlumotlar tuzilmasi boʻlib, lekin `oʻzgaruvchan oʻlchamga` ega. Masalan, bizda C++ da `vektor` va Golang da `slice` mavjud.

## Operations in Dynamic Array

Keling, dinamik massivdan foydalanishni ko'rib chiqaylik:

```cpp
#include <iostream>

int main() {
    // 1. initialize
    vector<int> v0;
    vector<int> v1(5, 0);
    // 2. make a copy
    vector<int> v2(v1.begin(), v1.end());
    vector<int> v3(v2);
    // 3. cast an array to a vector
    int a[5] = {0, 1, 2, 3, 4};
    vector<int> v4(a, *(&a + 1));
    // 4. get length
    cout << "The size of v4 is: " << v4.size() << endl;
    // 5. access element
    cout << "The first element in v4 is: " << v4[0] << endl;
    // 6. iterate the vector
    cout << "[Version 1] The contents of v4 are:";
    for (int i = 0; i < v4.size(); ++i) {
        cout << " " << v4[i];
    }
    cout << endl;
    cout << "[Version 2] The contents of v4 are:";
    for (int& item : v4) {
        cout << " " << item;
    }
    cout << endl;
    cout << "[Version 3] The contents of v4 are:";
    for (auto item = v4.begin(); item != v4.end(); ++item) {
        cout << " " << *item;
    }
    cout << endl;
    // 7. modify element
    v4[0] = 5;
    // 8. sort
    sort(v4.begin(), v4.end());
    // 9. add new element at the end of the vector
    v4.push_back(-1);
    // 10. delete the last element
    v4.pop_back();
}
```

[C++ Playground](https://leetcode.com/playground/4Rk4iKhu)

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. initialize
	v0 := []int{}
	fmt.Println("v0", v0)

	v1 := make([]int, 5)
	fmt.Println("v1", v1)

	// 2. make a copy
	v2 := append([]int(nil), v1...)
	fmt.Println("v2", v2)

	v3 := append([]int(nil), v2...)
	fmt.Println("v3", v3)

	// 3. cast an array to a slice
	a := [5]int{0, 1, 2, 3, 4}
	v4 := a[:]
	fmt.Println("v4", v4)

	// 4. get length
	fmt.Println("The size of v4 is:", len(v4))

	// 5. access element
	fmt.Println("The first element in v4 is:", v4[0])

	// 6. iterate the slice
	fmt.Println("[Version 1] The contents of v4 are:")
	for i := 0; i < len(v4); i++ {
		fmt.Print(" ", v4[i])
	}
	fmt.Println()
	fmt.Println("[Version 2] The contents of v4 are:")
	for _, item := range v4 {
		fmt.Print(" ", item)
	}
	fmt.Println()

	// 7. modify element
	v4[0] = 5

	// 8. sort
	sort.Ints(v4)

	// 9. add new element at the end of the slice
	v4 = append(v4, -1)

	// 10. delete the last element
	v4 = v4[:len(v4)-1]
}
```
[Go Playground](https://go.dev/play/p/3n1VlWy7ZOa)

© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/201/introduction-to-array/1146/)
