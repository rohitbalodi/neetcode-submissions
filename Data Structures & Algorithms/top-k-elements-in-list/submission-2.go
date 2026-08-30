func topKFrequent(nums []int, k int) []int {
    // Let's count the frequency of each element
	mymap := make(map[int]int)
	for _, v := range nums {
		mymap[v] += 1
	}

	// {3:2, 4:1, 1:4}

	// We will create a frequency bucket
	// But since more than 1 element could have same frequency
	// We should have a 2D array

	bucket := make([][]int, len(nums)+1)
	for num, count := range mymap {
		bucket[count] = append(bucket[count], num)
	}

	result := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0; i-- {
		result = append(result, bucket[i]...)
		if len(result) == k {
			break
		}
	}
	return result

}
