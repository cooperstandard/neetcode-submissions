func hasDuplicate(nums []int) bool {
    seen := make(map[int]int)

    for _,v := range nums {
        seen[v] += 1
    }

    for _, timesSeen := range seen {
        if timesSeen > 1 {
            return true
        }
    }
    return false
}
