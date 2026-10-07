type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result string

	for _, str :=  range strs {
		for _, c := range str {
			result += string(c+256)
		}
		result += "a"
	}

	return result

}

func (s *Solution) Decode(encoded string) []string {
	var result []string

	current := ""
	for _, v := range encoded {
		if v == 'a' {
			result = append(result, current)
			current = ""
		} else {
			current += string(v - 256)
		}
	}

	return result


}
