func longestConsecutive(nums []int) int {
    // the starting point could only be those numbers whose before value
    // does not exist, ex: arr = [1,2,3,4,5]
    // 2 cannot be starting point since 1 is already present
    // 1 can be the starting point since 0 is not present
    mymap := make(map[int]bool)

    for _,value := range(nums) {
        mymap[value] = true
    }

    best := 0

    for _, value := range(nums) {
        // this could be the starting point
        if mymap[value-1] == false {
            currEle := value
            curr := 0
            for mymap[currEle]==true {
                curr+=1
                best = max(best, curr)
                currEle+=1
            }
        }
    }
    return best

}
