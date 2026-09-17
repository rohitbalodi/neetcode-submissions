func longestIncreasingPath(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i:= range dp {
		dp[i] = make([]int, cols)
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		// If a element in matrix is non-zero, it means
		// LIS Path is calculated for that one already
		if dp[r][c] !=0 {
			return dp[r][c]
		}

		longest := 1
		directions := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
		for _, v := range directions {
			indexi := r+v[0]
			indexj := c+v[1]
			if indexi >=0 && indexj >=0 && indexi<rows && indexj<cols &&
			matrix[indexi][indexj] > matrix[r][c] {
				currLength := 1 + dfs(indexi, indexj)
				longest = max(longest, currLength)
			}
		}
		dp[r][c] = longest
		return longest
	}

	maxEle := -1
	for i:=0; i<len(dp); i++ {
		for j:=0; j<len(dp[i]); j++ {
			maxEle = max(maxEle, dfs(i,j))
		}
	}

	return maxEle
}
