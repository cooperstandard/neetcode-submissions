func isAnagram(s string, t string) bool {
	letters := make(map[rune]int)
	for _, letter := range s {
		letters[letter] += 1
	}

	for _, letter := range t {
		letters[letter] -= 1
	}

	for _, v := range letters {
		if v != 0 {
			return false
		}
	}

	return true

}
