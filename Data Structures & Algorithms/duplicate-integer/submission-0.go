func hasDuplicate(nums []int) bool {
	mymaps := make(map[int]int)
	for i, v := range nums {
		if _, ok := mymaps[v]; ok {
			return true
		} else {
			mymaps[v] = i
		}
	}
	return false

}