package findindexinstring

func StrStr(haystack string, needle string) int {
	output := -1

	// sadbutsad = 9
	// sad = 3
	// 9-3 = 6
	// i <= 6 = 0123456
	// output = i
	// if haystack[0:3+0] == needle = return output
	// haystack[0:3+0] = sad
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			output = i
			break
		}
	}

	return output
}
