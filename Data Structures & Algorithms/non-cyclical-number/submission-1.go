func isHappy(n int) bool {
	mymap := make(map[int]bool)
	for n>0 {
		if mymap[n]==false {
			mymap[n]=true
		} else {
			return false
		}
		n = getSquareOfDigits(n)
		if n==1 {
			return true
		}
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
