package zeros

func zero(x int) (res int) {
	for x > 0 {
		res += x / 5
		x /= 5
	}
	return
}

func preimageSizeFZF(K int) int {
	l := 0
	r := (K + 1) * 5
	for l <= r {
		m := l + (r-l)/2
		zeros := zero(m)
		switch {
		case zeros < K:
			l = m + 1
		case zeros > K:
			r = m - 1
		default:
			return 5
		}
	}
	return 0
}
