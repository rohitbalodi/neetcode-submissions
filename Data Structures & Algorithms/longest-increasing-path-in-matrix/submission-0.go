func longestIncreasingPath(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i:= range dp {
		dp[i] = make([]int, cols)
	}

	var dfs func(r, c, prevVal int) int
	dfs = func(r, c, prevVal int) int {
		if r<0 || c<0 || r>=rows || c>=cols || matrix[r][c] <=prevVal {
			return 0
		}

		// If a element in matrix is non-zero, it means
		// LIS Path is calculated for that one already
		if dp[r][c] !=0 {
			return dp[r][c]
		}

		// If no LIS is present, still LIS will be 1
		// because of that element
		res := 1
		res = max(res, 1+ dfs(r+1, c, matrix[r][c]))
		res = max(res, 1+ dfs(r-1, c, matrix[r][c]))
		res = max(res, 1+ dfs(r, c+1, matrix[r][c]))
		res = max(res, 1+ dfs(r, c-1, matrix[r][c]))
		dp[r][c] = res
		return res
	}

	maxEle := -1
	for i:=0; i<len(dp); i++ {
		for j:=0; j<len(dp[i]); j++ {
			maxEle = max(maxEle, dfs(i,j,-1))
		}
	}

	return maxEle
}
