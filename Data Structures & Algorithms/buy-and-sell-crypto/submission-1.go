func maxProfit(prices []int) int {
	least, best := prices[0], 0

	for _, p := range prices {
		least = min(least, p)
		best = max(best, p - least) 
	}

	return best

}
