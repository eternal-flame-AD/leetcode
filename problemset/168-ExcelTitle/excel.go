package excel

import (
	"bytes"
)

func constructCol(n int) byte {
	if n == 0 {
		return byte(64 + 26)
	}
	return byte(64 + n)
}

func convertToTitle(n int) string {
	res := make([]byte, 0)
	for n != 0 {
		res = append(res, constructCol(n%26))
		n = (n - 1) / 26
	}
	resstr := bytes.NewBufferString("")
	for i := len(res) - 1; i >= 0; i-- {
		resstr.WriteByte(res[i])
	}
	return resstr.String()
}
