package atoi

import (
	"regexp"
)

const (
	MIN_INT     = -2147483648
	MIN_INT_STR = "2147483648"
	MAX_INT     = 2147483647
	MAX_INT_STR = "2147483647"
)

var regex = regexp.MustCompile(`^\s*([\+/\-]?)(0*)(\d+)`)

func myAtoi(str string) int {
	match := regex.FindAllStringSubmatch(str, 1)
	if len(match) == 0 {
		return 0
	}
	symbol := match[0][1]
	digits := match[0][3]
	if len(digits) > 10 || len(digits) == 10 && ((symbol == "-" && digits > MIN_INT_STR) || (symbol != "-" && digits > MAX_INT_STR)) {
		if symbol == "-" {
			return MIN_INT
		} else {
			return MAX_INT
		}
	}
	res := int32(0)
	for index := range digits {
		res *= 10
		res += int32(digits[index]) - int32('0')
	}
	if symbol == "-" {
		res = -res
	}
	return int(res)
}
