func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	var res, l, maxf int

	for r := range len(s) {
		count[s[r]]++ 
		maxf = max(count[s[r]], maxf)
		
		for (r - l +1) -maxf > k { // length of the current window - matches in that window can't be more than k because everything that doesn't match in the window would need to be replaced
			count[s[l]]--
			l++
		}

		res = max(res, r-l+1)
	}
	return res
}


