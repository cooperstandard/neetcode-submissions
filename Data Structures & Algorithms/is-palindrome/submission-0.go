

func isPalindrome(s string) bool {
    reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	// Replace all matched non-alphanumeric characters with an empty string.
	s = strings.ToUpper(reg.ReplaceAllString(s, ""))
    if len(s) == 0 {
        return true
    }

    i, j := 0, len(s) - 1

    for i < j {
        if !(s[i] == s[j]) {
            return false
        }
        i += 1
        j -= 1
    }
    return true
}
