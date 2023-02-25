package medium

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	palindrome := string(s[0])

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			cur := s[i : j+1]
			if len(cur) <= len(palindrome) {
				break
			}
			if isPalindrome(cur) && len(cur) > len(palindrome) {
				palindrome = cur
			}
		}
	}

	return palindrome
}

func isPalindrome(characters string) bool {
	for i := 0; i <= len(characters)/2; i++ {
		if characters[i] != characters[len(characters)-1-i] {
			return false
		}
	}

	return true
}
