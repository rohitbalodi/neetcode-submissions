func twoSum(nums []int, target int) []int {
    mymap := make(map[int]int)
    for index, value := range(nums) {
        _, ok := mymap[target-value]
        if ok {
            first := min(index, mymap[target-value])
            second := max(index, mymap[target-value])
            return []int{first, second}
        } else {
            mymap[value] = index
        }
    }
    
    return []int{}
    
}
