package utils

func SubStringFormHead(str string, n int) string {
	rs := []rune(str)
	return string(rs[n:])
}

func SubStringFormEnd(str string, n int) string {
	rs := []rune(str)
	return string(rs[:n])
}
