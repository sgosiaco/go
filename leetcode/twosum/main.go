func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		element, ok := m[target-val]
		if ok {
			return []int{element, index}
		} else {
			m[val] = index
		}
	}

	return []int{-1, -1}
}