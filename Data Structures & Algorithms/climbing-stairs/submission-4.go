func climbStairs(n int) int {
	current, prev := 1, 1

	for range n - 1 {
		current, prev = current + prev, current
	}

	return current
}
