func isHappy(n int) bool {
	mymap := make(map[int]bool)
	for n>0 {
		if n==1 {
			return true
		}
		if mymap[n]==true {
			return false
		} else {
			mymap[n]=true
		}
		n = getSquareOfDigits(n)
	}
	return false
    
}

func getSquareOfDigits(num int) int {
	output := 0
	for num>0 {
		remainder := num%10
		output = output + (remainder*remainder)
		num = num/10
	}
	return output
}
