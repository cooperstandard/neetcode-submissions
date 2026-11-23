func minWindow(s string, t string) string {
	tContents := contents(t)
	sContents := contents("")
	var best string

	var l, r int
	matches := func(sContents map[byte]int) bool {
		for k,v := range tContents {
			if v > sContents[k] {
				return false
			}
		}
		return true
	}

	for ; r < len(s); r++ {
		sContents[s[r]]++
		if matches(sContents) {
			for matches(sContents) {
				sContents[s[l]]--
				l++
			}
			l--
			sContents[s[l]]++
			if len(best) > r - l + 1 || best == ""{
				best = s[l:r+1]
			}
		}
	}
	return best
    
}

func contents(s string) map[byte]int {
	ret := make(map[byte]int)
	for _, v := range s {
		ret[byte(v)]++
	}

	return ret
}

func matches(tContents, sContents map[rune]int) bool {
	for k,v := range tContents {
		if v > sContents[k] {
			return false
		}
	}
	return true
}
