package problem0049

import "sort"

func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, len(strs))

	for _, s := range strs {
		key := sortString(s)
		m[key] = append(m[key], s)
	}

	resp := make([][]string, 0, len(m))
	for _, v := range m {
		resp = append(resp, v)
	}

	return resp
}

func sortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
