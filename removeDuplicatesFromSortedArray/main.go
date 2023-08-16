package removeduplicates

func RemoveDuplicates(nums []int) int {
	var output int

	for i := 0; i < len(nums); i++ {
		left := i

		for j := i + 1; j < len(nums); j++ {
			if nums[left] == nums[j] {
				nums[j] = -777
			}
		}

		if nums[i] != -777 {
			output++
		}
	}

	return output
}
