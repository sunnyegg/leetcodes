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

		if string(s[i]) == "I" {
			if string(s[i+1]) == "V" {
				output += m[string(s[i+1])] - m[string(s[i])]
				i++
				continue
			}

			if string(s[i+1]) == "X" {
				output += m[string(s[i+1])] - m[string(s[i])]
				i++
				continue
			}
		}

		if string(s[i]) == "X" {
			if string(s[i+1]) == "L" {
				output += m[string(s[i+1])] - m[string(s[i])]
				i++
				continue
			}

			if string(s[i+1]) == "C" {
				output += m[string(s[i+1])] - m[string(s[i])]
				i++
				continue
			}
		}

		if string(s[i]) == "C" {
			if string(s[i+1]) == "D" {
				output += m[string(s[i+1])] - m[string(s[i])]
				i++
				continue
			}

			if string(s[i+1]) == "M" {
				output += m[string(s[i+1])] - m[string(s[i])]
				i++
				continue
			}
		}

		output += m[string(s[i])]
	}

	return output
}
