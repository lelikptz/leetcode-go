package problems

import (
	"fmt"
	"strconv"
)

func AddBinary(a string, b string) string {
	var large, small string
	if len(a) > len(b) {
		large, small = a, b
	} else {
		large, small = b, a
	}

	carry := 0
	for i := 0; i < len(small); i++ {
		aa, _ := strconv.Atoi(string(large[len(large)-1-i]))
		bb, _ := strconv.Atoi(string(small[len(small)-1-i]))
		ccc := aa + bb + carry
		if ccc > 1 {
			ccc = 0
			carry++
		}
		large = large[0:len(large)-1-i] + strconv.Itoa(ccc) + large[len(large)-i:]
		fmt.Println(large)
		return ""
	}

	return large
}
