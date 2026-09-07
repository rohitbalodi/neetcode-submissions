func carFleet(target int, position []int, speed []int) int {
	stack := []float64{}
	carsWithIndex := make([][2]int, len(position))
	for i:=0; i<len(position); i++ {
		carsWithIndex[i][0] = position[i] 
		carsWithIndex[i][1] = speed[i] 
	}

	sort.Slice(carsWithIndex, func(i, j int) bool {
		return carsWithIndex[i][0] > carsWithIndex[j][0]
	})



	for _, val := range(carsWithIndex) {
		time := float64(target - val[0])/float64(val[1])
		stack = append(stack, time)
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)

}
