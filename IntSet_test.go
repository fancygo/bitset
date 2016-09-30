package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	var intSet IntSet
	intSet.Add(12)
	intSet.Add(2)
	intSet.Add(255)
	intSet.Add(20)
	fmt.Println(intSet.Has(255))
	fmt.Println(intSet.Has(3))
	fmt.Println(intSet.String())
}
