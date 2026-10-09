func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := []rune{}

	opening := map[rune]bool{
		'(':true,
		'[':true,
		'{':true,
	}

	matching := map[rune]rune{
		'(':')',
		'[':']',
		'{':'}',
	}

	for _, v := range s{
		if opening[v] {
			stack = append(stack, v)
			continue
		}
		if v == matching[peak(stack)] {
			stack = pop(stack)
		} else {
			return false
		}
	}

	return len(stack) == 0
    
}

func peak(stack []rune) rune {
	if len(stack) == 0 {
		return '1'
	}
	
	return stack[len(stack)-1]
}

func pop(stack []rune) []rune {
	if stack == nil {
		return nil
	}

	return stack[0:len(stack)-1]
}
