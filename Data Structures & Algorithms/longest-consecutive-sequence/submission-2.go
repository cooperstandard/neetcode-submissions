func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	
	current, longest := 0, 0
	next := nums[0]

	for _, v := range nums {
		if v == next - 1 {
			continue
		}
		if v == next {
			next++
			current++
		} else {
			longest = max(longest, current)
			current = 1
			next = v + 1
		}
	}

	return max(longest, current)

	//count runs
}

