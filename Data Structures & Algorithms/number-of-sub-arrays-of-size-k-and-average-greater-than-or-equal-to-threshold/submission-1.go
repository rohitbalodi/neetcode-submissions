func numOfSubarrays(arr []int, k int, threshold int) int {
	prefix := make([]int, len(arr)+1)
	prefix[0] = 0

	// create a prefix sum array
	for i:=1; i<=len(arr); i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	// Leave the left element and include the right element
	// by subtracting prefxi[right] - prefix[left]
	left, right, output := 0, k, 0
	for right<len(prefix) {
		currAvg := (prefix[right] - prefix[left])/k
		if currAvg >= threshold {
			output++
		}
		left++
		right++
	}

	return output

}
