func productExceptSelf(nums []int) []int {
    // array [1,2,4,6]
    // prefix [1,1,2,8]
    // postfix [48,24,6,1]
    // output [48,24,12,8]

	prefix, postfix := make([]int, len(nums)), make([]int, len(nums))
	prefix[0] = 1
	postfix[len(postfix)-1] = 1

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		postfix[i] = postfix[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		postfix[i] = postfix[i] * prefix[i]
	}

	return postfix
}
