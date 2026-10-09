func missingNumber(nums []int) int {
	//sum of 0..n = (n^2)/2
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return (((len(nums)*(len(nums)+1))/2))- sum

}
