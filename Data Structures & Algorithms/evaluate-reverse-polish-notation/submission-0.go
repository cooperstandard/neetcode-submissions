
func evalRPN(tokens []string) int {
  var evalStack []int = nil 
  // result will be the top of the stack

  var a, b int
  for _, v := range tokens {
    switch v {
      case "+":
        // do addition
        a, evalStack = pop(evalStack)
        b, evalStack = pop(evalStack)
        evalStack = append(evalStack, a+b)
      case "-":
        // do subtraction
        a, evalStack = pop(evalStack)
        b, evalStack = pop(evalStack)
        evalStack = append(evalStack, b-a)
      case "*":
        // do multiplication
        a, evalStack = pop(evalStack)
        b, evalStack = pop(evalStack)
        evalStack = append(evalStack, a*b)
      case "/":
        a, evalStack = pop(evalStack)
        b, evalStack = pop(evalStack)
        evalStack = append(evalStack, b/a)
        // do division
      default:
        // push new number to the stack
        val, err := strconv.Atoi(v)
        if err != nil {
          return 0
        }
        evalStack = append(evalStack, val)
    }

    
  }

  top, evalStack := pop(evalStack)
  return top
}

func pop(s []int) (int, []int) {
  if len(s) == 0 {
    // should really be checking this before calling but since this is leetcode it is gauranteed not to be called unless the stack is not empty
    return 0, nil
  }
  return s[len(s)-1], s[:len(s)-1]
}