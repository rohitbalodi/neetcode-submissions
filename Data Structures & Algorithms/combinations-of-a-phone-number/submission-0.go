func letterCombinations(digits string) []string {
	if len(digits)==0 {
		return []string{}
	}
	output := []string {}
	myMap := map[byte] string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	input := []string{}
	for i:=0; i<len(digits); i++ {
		input = append(input, myMap[digits[i]])
	}

	var backtrack func(curr string, i int, input []string)
	backtrack = func(curr string, i int, input []string) {
		if len(curr)==len(digits) {
			output = append(output, curr)
			return
		}

		for j:=0; j<len(input[i]); j++ {
			backtrack(curr+string(input[i][j]), i+1, input)
		}

	}

	backtrack("",0,input)
	return output
	
}
