func threeSum(nums []int) [][]int {
	res := [][]int{}
	resMap := make(map[[3]int]bool)
	sort.Ints(nums)

	for i, a := range nums {
		if a > 0 {
			break
		}

		if i > 0 && a == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums) -1

		for l < r {
			threeSum := a + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				resMap[[3]int{a, nums[l], nums[r]}] = true
				l++
				r--
			}
		}
	}	


	for three := range resMap {
		res = append(res, []int{three[0], three[1], three[2]})
	}
	return res 


}