func isValid(s string) bool {
   var stack []rune
   for _, v := range s {
        if v == '(' || v == '[' || v == '{' {
            stack = append(stack, v)
        } else {
            if len(stack) == 0 {
                return false
            }
            v2 := stack[len(stack)-1]
            comparison := '('
            switch v {
                case ']':
                    comparison = '['
                case '}':
                    comparison = '{'
                default:
                    comparison = '('
            }
            if v2 != comparison {
                return false
            }
            stack = stack[:len(stack)-1]
        }
        
   } 
   return len(stack) == 0
}
