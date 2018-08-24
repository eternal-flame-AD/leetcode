package medoftwosortedarray

import (
	"fmt"
	"testing"
)

func TestMed(t *testing.T) {
	n1 := []int{}
	n2 := []int{1}
	fmt.Println(findMedianSortedArrays(n1, n2))
}
