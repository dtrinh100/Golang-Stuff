package stack

import (
	"log"
	"math"
	"strconv"
	"unicode"
)

type Stack struct {
	Arr []interface{}
	len int
}

func NewStack() Stack {
	var s Stack
	return s
}

func (s Stack) Peek() interface{} {
	if len(s.Arr) < 1 {
		panic("Stack Underflow")
	}
	return s.Arr[len(s.Arr)-1]
}

func (s *Stack) Push(e interface{}) {
	s.Arr = append(s.Arr, e)
}

func (s *Stack) Pop() interface{} {
	if len(s.Arr) < 1 {
		panic("Stack underflow")
	}
	prevEle := s.Arr[len(s.Arr)-1]
	s.Arr = append(s.Arr[:len(s.Arr)-1], s.Arr[len(s.Arr):]...)
	return prevEle
}

func (s Stack) IsEmpty() bool {
	return len(s.Arr) == 0
}

// compareOp returns 0, if operators are equal, -1 if current operator is less than the operator in the stack, and
// 1 if the current operator is greater than the operator in the stack
func compareOp(s1, s2 string) int {
	val := 100000
	if s1 == "^" && s2 == "^" {
		val = 0
	}
	if s1 == "/" && (s2 == "*" || s2 == "/") {
		val = 0
	}
	if s1 == "*" && (s2 == "*" || s2 == "/") {
		val = 0
	}
	if s1 == "+" && (s2 == "+" || s2 == "-") {
		val = 0
	}
	if s1 == "-" && (s2 == "+" || s2 == "-") {
		val = 0
	}
	if s1 == "/" && (s2 == "+" || s2 == "-") {
		val = 1
	}
	if s1 == "*" && (s2 == "+" || s2 == "-") {
		val = 1
	}
	if s1 == "^" && (s2 == "+" || s2 == "-" || s2 == "*" || s2 == "/") {
		val = 1
	}
	if (s1 == "+" || s1 == "-" || s1 == "*" || s1 == "/") && s2 == "^" {
		val = -1
	}
	if s1 == "+" && (s2 == "*" || s2 == "/") {
		val = -1
	}
	if s1 == "-" && (s2 == "*" || s2 == "/") {
		val = -1
	}
	return val

}

// operate takes 2 digits and an operand and returns the computation of the two digits based on the operand
func operate(d1, d2 int, operand string) int {
	val := 0

	switch operand {
	case "^":
		val = int(math.Pow(float64(d1), float64(d2)))
	case "/":
		val = d1 / d2
	case "*":
		val = d1 * d2
	case "+":
		val = d1 + d2
	case "-":
		val = d1 - d2
	}
	return val
}

// checkEquiv takes two string and returns a boolean that tells if the two strings are balanced expressions
// Ex: { and } are balanced, but { and { are not
func checkEquiv(s1, s2 string) bool {
	isEquiv := false
	if s1 == "(" && s2 == ")" {
		isEquiv = true
	}
	if s1 == "{" && s2 == "}" {
		isEquiv = true
	}
	if s1 == "[" && s2 == "]" {
		isEquiv = true
	}

	return isEquiv

}

// InfixToPostfix is a function used to convert any given infix string that has operators in it into a postfix string
// Ex: a+b*c-d becomes abc*+d-
func InfixToPostfix(s string) string {
	postfix := ""
	st := NewStack()

	// Scans each character in the string
	for i := 0; i < len(s); i++ {

		if s[i:i+1] == "+" || s[i:i+1] == "-" || s[i:i+1] == "*" || s[i:i+1] == "/" || s[i:i+1] == "^" {
			if st.IsEmpty() || compareOp(s[i:i+1], st.Peek().(string)) == 1 { // If the stack is empty or the current operator is greater than the one on the stack, the push
				st.Push(s[i : i+1])

			} else {
				// While the current operator is less than or equal to the one in the stack, then pop the operator from the stack and add it to the postfix string
				for compareOp(s[i:i+1], st.Peek().(string)) == -1 || compareOp(s[i:i+1], st.Peek().(string)) == 0 {

					postfix += st.Pop().(string)

					if st.IsEmpty() { // If the stack is empty then we break out of the loop
						break
					}

				}

				st.Push(s[i : i+1]) // Adds the current operator onto the stack
			}
		} else if s[i:i+1] != "(" && s[i:i+1] != ")" { // This means that the char is a operand
			postfix += s[i : i+1]
		}

		if s[i:i+1] == "(" {
			st.Push(s[i : i+1])
		}
		if s[i:i+1] == ")" {
			for st.Peek().(string) != "(" { // Outputs all the content of the stack until it hits an opening parenthesis
				postfix += st.Pop().(string)

			}
			st.Pop() // Gets rid of the opening parenthesis
		}

	}
	// Emptys the rest of the stack
	for !st.IsEmpty() {

		postfix += st.Pop().(string)

	}
	return postfix

}

// EvalPostfix evaulates a postfix expression and returns the integer value of the computation.
func EvalPostfix(pf string) int {
	result := 0
	sLen := len(pf)
	st := NewStack()

	for i := 0; i < sLen; i++ {
		if unicode.IsDigit(rune(pf[i])) { // If the character is a digit, then just push it onto the stack
			tmp, err := strconv.Atoi(pf[i : i+1])
			if err != nil {
				log.Println(err)
			}
			st.Push(tmp)

		} else { // Takes the 1st operator and 2nd operator and operates on 2nd operator operand 1st operator. Ex. d2 * d1
			d1 := st.Pop().(int)
			d2 := st.Pop().(int)
			total := operate(d2, d1, pf[i:i+1])
			st.Push(total) // Result is pushed onto stack

		}

	}
	result = st.Pop().(int) // There is only one value left and that value is the result
	return result
}

// Check balance checks for balanced parentheses in a given expression
func CheckBalance(s string) bool {
	isValid := true
	sLen := len(s)
	st := NewStack()

	for i := 0; i < sLen; i++ {
		if s[i:i+1] == "(" || s[i:i+1] == "[" || s[i:i+1] == "{" { // Push all opening parentheses to the stack
			st.Push(s[i : i+1])
		} else if s[i:i+1] == ")" || s[i:i+1] == "]" || s[i:i+1] == "}" {
			if st.IsEmpty() { // If a expression starts with }, ], or ), then set it to false right away and break out of the loop
				isValid = false
				break
			}
			if !checkEquiv(st.Pop().(string), s[i:i+1]) { // If the string pop from the stack does not match, then it is invalid
				isValid = false
				break
			}

		}
	}
	if !st.IsEmpty() { // If an expression ends with an open parentheses, then there will be one left on the stack, so set isValid to false
		isValid = false
	}
	return isValid

}
