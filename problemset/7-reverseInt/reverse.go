package reverseint

import (
	"bytes"
	"strconv"
)

func reverse(x int) int {
	var negative bool
	if x < 0 {
		x = -x
		negative = true
	}
	xstr := strconv.FormatInt(int64(x), 10)
	resbytes := bytes.NewBufferString("")
	if negative {
		resbytes.WriteString("-")
	}
	for i := len(xstr) - 1; i >= 0; i-- {
		resbytes.WriteByte(xstr[i])
	}
	if res, err := strconv.ParseInt(resbytes.String(), 10, 32); err == nil {
		return int(res)
	} else {
		return 0
	}
}
