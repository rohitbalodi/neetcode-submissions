func generateParenthesis(n int) []string {
	output := []string {}

	var backtrack func(open int, close int, curr string)
	backtrack = func(open int, close int, curr string) {
		if len(curr)==2*n {
			output = append(output, curr)
			return
		}

		if open<n {
			backtrack(open+1, close, curr+"(")
		}

		if close<open {
			backtrack(open, close+1, curr+")")
		}

	}

	backtrack(0,0,"")
	return output

}
