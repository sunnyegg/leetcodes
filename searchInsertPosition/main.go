package searchinsertposition

// sorted
func SearchInsert(nums []int, target int) int {
	// jika yg terakhir masih jg lebih kecil target, maka target taroh di belakang
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}

	return -1
}
