package lps

func detectPalindromeFromIndex(s *string, l int, r int) int {
	delta := 0
	for l-delta >= 0 && r+delta < len(*s) && (*s)[l-delta] == (*s)[r+delta] {
		delta++
	}
	return delta - 1
}

func calcLength(l, r, delta int) int {
	if l < 0 || r < 0 {
		return -1
	}
	return (r + delta) - (l - delta) + 1
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	l := len(s) / 2
	r := (len(s) - 1) / 2
	maxmedpointl := -1
	maxmedpointr := -1
	maxresdelta := 0
	var medpoint int
	for delta := 0; delta <= l; delta++ {
		//search for l
		medpoint = l - delta
		if medpoint < maxresdelta || len(s)-medpoint-1 < maxresdelta {
		} else {
			resdelta := detectPalindromeFromIndex(&s, medpoint, medpoint)
			if calcLength(medpoint, medpoint, resdelta) > calcLength(maxmedpointl, maxmedpointr, maxresdelta) {
				maxresdelta = resdelta
				maxmedpointl = medpoint
				maxmedpointr = medpoint
			}
			resdelta = detectPalindromeFromIndex(&s, medpoint, medpoint+1)
			if calcLength(medpoint, medpoint+1, resdelta) > calcLength(maxmedpointl, maxmedpointr, maxresdelta) {
				maxresdelta = resdelta
				maxmedpointl = medpoint
				maxmedpointr = medpoint + 1
			}
		}

		//search for r
		medpoint = r + delta
		if medpoint < maxresdelta || len(s)-medpoint-1 < maxresdelta {
		} else {
			resdelta := detectPalindromeFromIndex(&s, medpoint, medpoint)
			if calcLength(medpoint, medpoint, resdelta) > calcLength(maxmedpointl, maxmedpointr, maxresdelta) {
				maxresdelta = resdelta
				maxmedpointl = medpoint
				maxmedpointr = medpoint
			}
			resdelta = detectPalindromeFromIndex(&s, medpoint-1, medpoint)
			if calcLength(medpoint-1, medpoint, resdelta) > calcLength(maxmedpointl, maxmedpointr, maxresdelta) {
				maxresdelta = resdelta
				maxmedpointl = medpoint - 1
				maxmedpointr = medpoint
			}
		}
	}
	if maxmedpointl == -1 {
		return ""
	}
	return s[maxmedpointl-maxresdelta : maxmedpointr+maxresdelta+1]
}
