func longestConsecutive(nums []int) int {
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
