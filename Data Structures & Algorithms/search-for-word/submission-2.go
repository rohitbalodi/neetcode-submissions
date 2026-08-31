func exist(board [][]byte, word string) bool {
	var backtrack func(i int, j int, currLen int)(bool)
	backtrack = func(i int, j int, currLen int) (bool) {
		if board[i][j] == '+' {
			return false
		}

		if board[i][j] != word[currLen] {
			return false
		}

		if currLen == len(word)-1 {
			return true
		}

		temp := board[i][j]
		board[i][j] = 43
		var a,b,c,d bool
		if i>0 {
			a= backtrack(i-1, j, currLen+1)
		}
		if j>0 {
			b= backtrack(i, j-1, currLen+1)
		}
		if i<len(board)-1 {
			c= backtrack(i+1, j, currLen+1)
		}
		if j<len(board[0])-1 {
			d= backtrack(i, j+1, currLen+1)
		}

		board[i][j] = temp

		return a||b||c||d
	}

	for i:=0; i<len(board); i++ {
		for j:=0; j<len(board[0]); j++ {
			if backtrack(i,j,0) {
				return true
			}
		}
	}

	return false

}
