package paren

func checkValidString(s string) bool {
	l, r := 0, 0
	for _, char := range s {
		switch char {
		case '(':
			l++
			r++
		case ')':
			l--
			r--
		case '*':
			l--
			r++
		}
		if l < 0 {
			l = 0
		}
		if r < 0 {
			return false
		}
	}
	return l == 0
}
