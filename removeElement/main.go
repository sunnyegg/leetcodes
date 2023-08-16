package removeelement

func RemoveElement(nums []int, val int) int {
	var output int

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = -1
		} else {
			output++
		}
	}

	for i := 0; i < len(nums); i++ {
		pivot := i

		for j := i + 1; j < len(nums); j++ {
			if nums[pivot] == -1 {
				pivot = j
			}
		}

		if pivot == i {
			continue
		}

		// swap
		nums[i], nums[pivot] = nums[pivot], nums[i]
	}

	return output
}
