package simple

import (
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
	for i := 0; i < len(large); i++ {
		aa, _ := strconv.Atoi(string(large[len(large)-1-i]))
		bb := 0
		if i < len(small) {
			bb, _ = strconv.Atoi(string(small[len(small)-1-i]))
		}
		ccc := aa + bb + carry
		if ccc >= 2 {
			ccc = ccc - 2
			carry = 1
		} else {
			carry = 0
		}
		large = large[0:len(large)-1-i] + strconv.Itoa(ccc) + large[len(large)-i:]
	}
	if carry != 0 {
		large = strconv.Itoa(carry) + large
	}

	return large
}
