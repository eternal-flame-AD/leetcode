package eliminate

func lastRemaining(n int) int {
	res := 1
	diff := 1
	lfirst := true
	for n != 1 {
		if lfirst || n%2 == 1 {
			res += diff
		}

		n /= 2
		diff *= 2
		lfirst = !lfirst
	}
	return res
}
