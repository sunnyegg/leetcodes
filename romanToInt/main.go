package romantoint

func Calc(s string) int {
	var output int

	m := make(map[string]int)

	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	for i := 0; i < len(s); i++ {
		if i+1 >= len(s) {
			output += m[string(s[i])]
			break
		}

		if m[string(s[i])] < m[string(s[i+1])] {
			output -= m[string(s[i])]
		} else {
			output += m[string(s[i])]
		}
	}

	return output
}
