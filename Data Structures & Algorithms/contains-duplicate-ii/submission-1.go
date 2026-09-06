func containsNearbyDuplicate(nums []int, k int) bool {
	left, right := 0, k
	myMap := map[int] bool{}
	for i:=0; i<k && i<len(nums); i++ {
		if myMap[nums[i]] {
			return true
		}
		myMap[nums[i]] = true
	}

	for i:=right; i<len(nums); i++ {
		if myMap[nums[i]] {
			return true
		}
		myMap[nums[i]] = true
		myMap[nums[left]] = false
		left++
	}

	return false

}
