package medium

func LengthOfLongestSubstring(s string) int {
	max := 0
	m := make(map[string]int)
	var sl []string

	for i := 0; i < len(s); i++ {
		_, ok := m[string(s[i])]
		if ok == true {
			var j int
			for _, ss := range sl {
				j++
				if ss == string(s[i]) {
					break
				}
			}
			sl = sl[j:]
			m = make(map[string]int)
			for _, ss := range sl {
				m[ss] = 0
			}
		}

		m[string(s[i])] = 0
		sl = append(sl, string(s[i]))

		if len(sl) > max {
			max = len(sl)
		}
	}

	return max
}
