package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeek(t *testing.T) {
	s := NewStack()
	s.Push("y")
	assert.Equal(t, "y", s.Peek(), "Expecting y")
}

func TestPop(t *testing.T) {

	s := NewStack()
	s.Push("z")
	s.Push("David")
	assert.Equal(t, "David", s.Pop(), "Expecting David")
	assert.Equal(t, "z", s.Pop(), "Expecting z")
}

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push("y")
	assert.Equal(t, "y", s.Peek(), "Expecting y")

	s.Push("Hec")
	assert.Equal(t, "Hec", s.Peek(), "Expecting Hec")
}

func TestCompareOp(t *testing.T) {
	p := "+"
	s := "-"
	d := "/"
	m := "*"
	e := "^"

	/* Tests to see if the precedence are equal */
	assert.Equal(t, 0, compareOp(e, e), "Expecting 0")
	assert.Equal(t, 0, compareOp(p, p), "Expecting 0")
	assert.Equal(t, 0, compareOp(s, s), "Expecting 0")
	assert.Equal(t, 0, compareOp(m, m), "Expecting 0")
	assert.Equal(t, 0, compareOp(d, d), "Expecting 0")
	assert.Equal(t, 0, compareOp(p, s), "Expecting 0")
	assert.Equal(t, 0, compareOp(d, m), "Expecting 0")
	assert.Equal(t, 0, compareOp(m, d), "Expecting 0")
	assert.Equal(t, 0, compareOp(s, p), "Expecting 0")
	/* Tests to see if the current operator has higher precedence than the operator in the stack */
	assert.Equal(t, 1, compareOp(m, s), "Expecting 1")
	assert.Equal(t, 1, compareOp(d, s), "Expecting 1")
	assert.Equal(t, 1, compareOp(d, p), "Expecting 1")
	assert.Equal(t, 1, compareOp(m, p), "Expecting 1")
	assert.Equal(t, 1, compareOp(e, p), "Expecting 1")
	assert.Equal(t, 1, compareOp(e, m), "Expecting 1")
	assert.Equal(t, 1, compareOp(e, s), "Expecting 1")
	assert.Equal(t, 1, compareOp(e, d), "Expecting 1")
	/* Tests to see if the current operator has lower precedence than the operator in the stack */
	assert.Equal(t, -1, compareOp(s, m), "Expecting -1")
	assert.Equal(t, -1, compareOp(s, d), "Expecting -1")
	assert.Equal(t, -1, compareOp(p, m), "Expecting -1")
	assert.Equal(t, -1, compareOp(p, d), "Expecting -1")
	assert.Equal(t, -1, compareOp(p, e), "Expecting -1")
	assert.Equal(t, -1, compareOp(s, e), "Expecting -1")
	assert.Equal(t, -1, compareOp(m, e), "Expecting -1")
	assert.Equal(t, -1, compareOp(d, e), "Expecting -1")

}

func TestInfixToPostfix(t *testing.T) {
	s := "a+b*c-d"
	s1 := "a+b*(c^d-e)^(f+g*h)-i"
	s2 := "a+b*c"
	s3 := "a*(b+c)"

	assert.Equal(t, "abc*+d-", InfixToPostfix(s), "Expecting abc*+d-")
	assert.Equal(t, "abcd^e-fgh*+^*+i-", InfixToPostfix(s1), "Expecting abcd^e-fgh*+^*+i-")
	assert.Equal(t, "abc*+", InfixToPostfix(s2), "Expecting abc*+")
	assert.Equal(t, "abc+*", InfixToPostfix(s3), "Expecting abc+*")

}

func TestOperate(t *testing.T) {
	d1 := 3
	d2 := 2

	assert.Equal(t, 9, operate(d1, d2, "^"), "Expecting 9")
	assert.Equal(t, 8, operate(d2, d1, "^"), "Expecting 8")
	assert.Equal(t, 1, operate(d1, d2, "/"), "Expecting 1")
	assert.Equal(t, 6, operate(d1, d2, "*"), "Expecting 6")
	assert.Equal(t, 1, operate(d1, d2, "-"), "Expecting 1")
	assert.Equal(t, -1, operate(d2, d1, "-"), "Expecting -1")
	assert.Equal(t, 5, operate(d1, d2, "+"), "Expecting 5")

}

func TestEvalPostfix(t *testing.T) {
	s := "231*+9-"
	s1 := "568^+"

	assert.Equal(t, -4, EvalPostfix(s), "Expecting -4")
	assert.Equal(t, 1679621, EvalPostfix(s1), "Expecting 1679621")

}

func TestCheckEquiv(t *testing.T) {
	s1 := "("
	s2 := ")"
	s3 := "{"
	s4 := "}"
	s5 := "["
	s6 := "]"

	assert.True(t, checkEquiv(s1, s2), "Expected to be true")
	assert.True(t, checkEquiv(s3, s4), "Expected to be true")
	assert.True(t, checkEquiv(s5, s6), "Expected to be true")
	assert.False(t, checkEquiv(s2, s1), "Expected to be false")
	assert.False(t, checkEquiv(s4, s3), "Expected to be false")
	assert.False(t, checkEquiv(s6, s5), "Expected to be false")
	assert.False(t, checkEquiv(s1, s1), "Expected to be false")
	assert.False(t, checkEquiv(s2, s3), "Expected to be false")

}
func TestCheckBalance(t *testing.T) {
	s := "[()]{}{[()()]()}"
	s1 := "[(])"
	s2 := "}{()"
	s3 := "{}{"
	s4 := "[{}{}](){}[]({})([][])"

	assert.True(t, CheckBalance(s), "Expected to be true")
	assert.True(t, CheckBalance(s4), "Expected to be true")
	assert.False(t, CheckBalance(s1), "Expected to be false")
	assert.False(t, CheckBalance(s2), "Expected to be false")
	assert.False(t, CheckBalance(s3), "Expected to be false")
}
