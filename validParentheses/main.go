package validparentheses

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var output bool
	var parentheses []string
	m := make(map[string]string)
	m["("] = ")"
	m["{"] = "}"
	m["["] = "]"

	for _, t := range s {
		if string(t) == "(" || string(t) == "{" || string(t) == "[" {
			parentheses = append(parentheses, string(t))
		}

		if len(parentheses) == 0 {
			output = false
			break
		}

		if string(t) == ")" || string(t) == "}" || string(t) == "]" {
			checkPair := m[parentheses[len(parentheses)-1]] == string(t)
			if checkPair {
				parentheses = parentheses[:len(parentheses)-1]
				output = true
			} else {
				output = false
				break
			}
		}
	}

	if len(parentheses) > 0 {
		output = false
	}

	return output
}
