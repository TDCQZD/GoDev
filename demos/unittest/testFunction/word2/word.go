package word2

import "unicode"

// IsPalindrome用于检查一个字符串是否从前向后和从后向前读都是一样的
// 修复word1中字符不能识别的bug
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
