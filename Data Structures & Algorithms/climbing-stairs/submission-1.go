func climbStairs(n int) int {
    // bruteforce: compute all the ways
	// dynamic programming: iteratively build up the number of ways to get to any of the stairs
	// at every step after the first 2 you just need to add the 2 previous steps

	memory := make([]int, n)


	for i, _ := range memory {
		if i == 0 {
			memory[0] = 1
			continue
		}
		if i == 1 {
			memory[1] = 2
			continue
		}

		memory[i] = memory[i-1] + memory[i-2]


	}
	return memory[len(memory)-1]
}
