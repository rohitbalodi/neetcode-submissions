func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	stack :=[][]int{{intervals[0][0], intervals[0][1]}}

	output := 0
	for i:=1; i<len(intervals); i++ {
		top := stack[len(stack)-1]
		if intervals[i][0]<top[1] {
			output++
			// prefer a range which has less end
			// 1,8 vs 2,3
			if top[1]>intervals[i][1] {
				top[0] = intervals[i][0]
				top[1] = intervals[i][1]
			}
		} else {
			stack = append(stack, intervals[i])
		}
	}

	return output
}