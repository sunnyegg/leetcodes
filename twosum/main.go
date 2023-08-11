package twosum

func Calc(nums []int, target int) []int {
	var output []int

	for i := range nums {
		if len(output) == 2 {
			break
		}

		for j := range nums {
			if i != j {
				if nums[i]+nums[j] == target {
					output = append(output, i, j)
					break
				}
			}
		}
	}

	return output
}
