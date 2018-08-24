package regexmatch

import (
	"testing"
)

func TestMed(t *testing.T) {
	if isMatch("aa", "a") {
		t.FailNow()
	}
	if !isMatch("a", "ab*") {
		t.FailNow()
	}
	if !isMatch("aab", "c*a*b") {
		t.FailNow()
	}
	if isMatch("mississippi", "mis*is*p*.") {
		t.FailNow()
	}
}
