func minCostClimbingStairs(cost []int) int {
	left, mid := cost[0], cost[1]

	for i:=2; i<len(cost); i++ {
		right := min(left, mid) + cost[i]
        left = mid
        mid = right
	}

	return min(left, mid)
}
