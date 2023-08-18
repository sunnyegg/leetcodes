package lengthoflastword

func LengthOfLastWord(s string) int {
	var output int
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			continue
		} else {
			output++
			if i-1 > 0 && string(s[i-1]) == " " {
				break
			}
		}
	}
	return output
}
