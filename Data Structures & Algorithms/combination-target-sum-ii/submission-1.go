func combinationSum2(candidates []int, target int) [][]int {
	output := [][]int{}
	sort.Ints(candidates)

	var backtrack func(index int, curr []int, remaining int)
	backtrack = func(index int, curr []int, remaining int) {
		if remaining==0 {
			temp := append([]int{}, curr...)
			output = append(output, temp)
			return
		}

		for i:=index; i<len(candidates); i++ {
			if candidates[i]>remaining {
				break
			}

			if i>index && candidates[i]==candidates[i-1] {
				continue
			}
			curr = append(curr, candidates[i])
			backtrack(i+1, curr, remaining - candidates[i])
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, []int{}, target)
	return output

}
