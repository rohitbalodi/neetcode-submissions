func lengthOfLongestSubstring(s string) int {
	left, right := 0,0
	myMap := map[byte] int {}
	best := 0

	for right<len(s) {
		// already has this character, we should move the
		// window from left to remove this character
		if myMap[s[right]]==1 {
			for myMap[s[right]]==1 {
				myMap[s[left]]-=1
				left+=1
			}
			myMap[s[right]]=1
			right+=1
		} else {
			myMap[s[right]]+=1
			right++
		}
		best = max(best, right-left)

	}

	return best

}
