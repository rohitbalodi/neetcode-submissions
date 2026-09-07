func numOfSubarrays(arr []int, k int, threshold int) int {
	prefix := make([]int, len(arr)+1)
	prefix[0] = 0

	for i:=1; i<=len(arr); i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	left, right, output := 0, k, 0
	for right<len(prefix) {
		var currAvg float64 = float64((prefix[right] - prefix[left])/k)
		fmt.Println(currAvg, threshold)
		if currAvg >= float64(threshold) {
			output++
		}
		left++
		right++
	}

	return output

}
