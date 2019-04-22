package word

// IsPalindrome用于检查一个字符串是否从前向后和从后向前读都是一样的
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
