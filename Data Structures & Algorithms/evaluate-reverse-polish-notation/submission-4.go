func evaluate(operation string, pop1 string, pop2 string) int {
	number1, _ := strconv.Atoi(pop1)
	number2, _ := strconv.Atoi(pop2)
	if operation == "+" {
		return number2 + number1
	} else if operation == "-" {
		return number2 - number1
	} else if operation == "*" {
		return number2 * number1
	} else {
		return number2 / number1
	}
}

func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/"{
			pop1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pop2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result := evaluate(token, pop1, pop2)
			stack = append(stack, strconv.Itoa(result))
		} else {
			stack = append(stack, token)
		}
	}
	output, _ := strconv.Atoi(stack[0])
	return output

}
