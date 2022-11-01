package simple

func LongestCommonPrefix(strings []string) string {
	counter := 0
	prefix := strings[0]
	for i := 0; i < len(strings); i++ {
		if len(prefix) <= len(strings[i]) && prefix == strings[i][:len(prefix)] {
			counter++
		} else {
			i = -1
			counter = 0
			prefix = prefix[:len(prefix)-1]
		}

		if counter == len(strings) {
			break
		}
	}

	return prefix
}
