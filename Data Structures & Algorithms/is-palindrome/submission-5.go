func isPalindrome(s string) bool {
	s = CleanString(s)
	fmt.Println(s)
	i, j := 0, len(s) -1

	for i <= j {
		if i >= j {
			return true
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

var alphaNumericRegex = regexp.MustCompile("[^a-zA-Z0-9]+")

func CleanString(str string) string {
    return strings.ToLower(alphaNumericRegex.ReplaceAllString(str, ""))
}