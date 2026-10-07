func productExceptSelf(nums []int) []int {
	prefixes := make([]int, len(nums))
	postfixes := make([]int, len(nums))

	prefixes[0] = 1

	for i, _ := range nums {
		if i == 0 {
			continue
		}

		prefixes[i] = prefixes[i-1] * nums[i-1]
	}

	postfixes[len(nums)-1] = 1

	for i := len(nums)-2; i >= 0; i-- {
		postfixes[i] = nums[i+1] * postfixes[i+1]
	}

	result := []int{}

	for i, v := range prefixes {
		result = append(result, v * postfixes[i])
	}

	return result


}
