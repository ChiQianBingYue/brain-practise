package P557ReverseWords

import "strings"

func _reverse(words []string) []string {
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return words
}

func _filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// ReverseWords 反转字符串
func ReverseWords(s string) string {
	wordArr := strings.Split(strings.Trim(s, " "), " ")
	wordArr = _filter(wordArr, func(v string) bool {
		return v != ""
	})
	return strings.Join(_reverse(wordArr), " ")
}
