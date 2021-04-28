func searchInsert(nums []int, target int) int {
	for index, value := range nums {
		if target == value {
			return index
		}
		if target < value {
			return index
		}
	}
	return len(nums)
}

func binInsert(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		median := (left + right) / 2
		if target > nums[median] {
			left = median + 1
		} else {
			right = median
		}
	}
	return left
}