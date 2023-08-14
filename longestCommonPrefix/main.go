package longestcommonprefix

func Calc(strs []string) string {
	var prevString string

	for i := 0; i < len(strs); i++ {
		var currentString string

		if i == 0 {
			prevString = strs[i]
			continue
		}

		for j := 0; j < len(prevString); j++ {
			if j >= len(strs[i]) {
				break
			}

			if string(prevString[j]) == string(strs[i][j]) {
				currentString += string(prevString[j])
			} else {
				break
			}
		}

		prevString = currentString
	}

	return prevString
}
