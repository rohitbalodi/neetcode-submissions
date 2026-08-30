func groupAnagrams(strs []string) [][]string {
	mymaps := make(map[[26]int][]int)

	for in, val := range strs {
		key := anagramtoarray(val)
		mymaps[key] = append(mymaps[key], in)
	}

	output := make([][]string, 0)
	for i, val := range mymaps {
		mymaps[i] = val
		sub := make([]string, 0)
		for _, index := range val {
			sub = append(sub, strs[index])
		}
		output = append(output, sub)
	}
	return output
}

func anagramtoarray(strs string) [26]int {
	var myarr = [26]int{}
	for i := 0; i < len(strs); i++ {
		myarr[strs[i]-'a'] += 1
	}
	return myarr
}