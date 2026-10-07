func twoSum(nums []int, target int) []int {
	partner := make(map[int]int) 
	// stores complement:index

	for i, num := range nums {
		partner[target - num] = i
	}

	for i, num := range nums {
		if j, ok := partner[num]; ok {
			if i == j {
				continue
			}
			return []int{min(i,j), max(i,j)}
		}
	}
	return []int{}
}
