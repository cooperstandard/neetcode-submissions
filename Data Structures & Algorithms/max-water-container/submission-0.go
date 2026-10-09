func maxArea(heights []int) int {
	biggest := 0
	i, j := 0, len(heights)-1

	for i < j {
		area := min(heights[i], heights[j]) * (j-i)
		biggest = max(biggest, area)

		if heights[i] < heights[j] {
			i ++
		} else {
			j--
		}
	}

	return biggest
}
