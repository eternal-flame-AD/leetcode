package RomanToInteger

var charintmap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res := 0
	for index := range s {
		if index != len(s)-1 && charintmap[s[index+1]] > charintmap[s[index]] {
			res -= charintmap[s[index]]
		} else {
			res += charintmap[s[index]]
		}
	}
	return res
}
