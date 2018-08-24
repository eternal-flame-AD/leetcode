package regexmatch

type regexOp struct {
	char        byte
	hasAsterisk bool
}

func byteMatch(s, p byte) bool {
	return p == '.' || s == p
}

func doSubMatch(s *string, spos int, opindex int, ops *[]*regexOp) bool {
	if spos == len(*s) && opindex == len(*ops) {
		return true
	}
	if opindex == len(*ops) {
		return false
	}
	if (*ops)[opindex].hasAsterisk {
		for matchcount := 0; matchcount <= len(*s)-spos; matchcount++ {
			if matchcount > 0 && !byteMatch((*s)[spos+matchcount-1], (*ops)[opindex].char) {
				break
			}
			if doSubMatch(s, spos+matchcount, opindex+1, ops) {
				return true
			}
		}
	} else {
		if spos != len(*s) && byteMatch((*s)[spos], (*ops)[opindex].char) {
			return doSubMatch(s, spos+1, opindex+1, ops)
		}
	}
	return false
}

func isMatch(s string, p string) bool {
	ops := make([]*regexOp, 0)
	for index := range p {
		if p[index] != '*' {
			ops = append(ops, &regexOp{
				char:        p[index],
				hasAsterisk: !(index == len(p)-1 || p[index+1] != '*'),
			})
		}
	}
	return doSubMatch(&s, 0, 0, &ops)
}
