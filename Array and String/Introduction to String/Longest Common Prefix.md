# Longest Common Prefix

Satrlar massivi orasidan eng uzun umumiy prefiks qatorini topish funksiyasini yozing.

Agar umumiy prefiks bo'lmasa, bo'sh `""` qatorini qaytaring.

#### Example 1:

```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

#### Example 2:

```
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: Kirish satrlari orasida umumiy prefiks yo'q.
```
 

#### Cheklovlar:

* `1 <= strs.length <= 200`
* `0 <= strs[i].length <= 200`
* `strs[i]` faqat inglizcha kichik harflardan iborat.


```go
func longestCommonPrefix(strs []string) string {

	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}
```
© Leetcode [link](https://leetcode.com/explore/learn/card/array-and-string/203/introduction-to-string/1162/)