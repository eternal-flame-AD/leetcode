package roman

import (
	"bytes"
)

func calculateDigit(onechar, fivechar, tenchar byte, num int) string {
	res := bytes.NewBufferString("")
	switch num {
	case 4:
		res.WriteByte(onechar)
		res.WriteByte(fivechar)
	case 9:
		res.WriteByte(onechar)
		res.WriteByte(tenchar)
	default:
		remaining := num
		if remaining >= 5 {
			res.WriteByte(fivechar)
			remaining -= 5
		}
		for i := 0; i < remaining; i++ {
			res.WriteByte(onechar)
		}
	}
	return res.String()
}

func intToRoman(num int) string {
	res := ""
	res += calculateDigit('M', 'A', 'A', num/1000)
	res += calculateDigit('C', 'D', 'M', num/100%10)
	res += calculateDigit('X', 'L', 'C', num/10%10)
	res += calculateDigit('I', 'V', 'X', num%10)
	return res
}
