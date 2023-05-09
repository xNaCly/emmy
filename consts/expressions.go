package consts

import "fmt"

type Expression interface {
	Eval() float64
	String() string
}
type PlusExpression struct {
	Lhs Expression
	Rhs Expression
}
type MinusExpression struct {
	Lhs Expression
	Rhs Expression
}
type MultiplicationExpression struct {
	Lhs Expression
	Rhs Expression
}
type DivisionExpression struct {
	Lhs Expression
	Rhs Expression
}

func (e PlusExpression) Eval() float64 {
	return e.Lhs.Eval() + e.Rhs.Eval()
}
func (e MinusExpression) Eval() float64 {
	return e.Lhs.Eval() - e.Rhs.Eval()
}
func (e MultiplicationExpression) Eval() float64 {
	return e.Lhs.Eval() * e.Rhs.Eval()
}
func (e DivisionExpression) Eval() float64 {
	return e.Lhs.Eval() / e.Rhs.Eval()
}
func (e PlusExpression) String() string {
	return fmt.Sprintf("[+; left=%s, right=%s]\n", e.Lhs.String(), e.Rhs.String())
}
func (e MinusExpression) String() string {
	return fmt.Sprintf("[-; left=%s, right=%s]\n", e.Lhs.String(), e.Rhs.String())
}
func (e MultiplicationExpression) String() string {
	return fmt.Sprintf("[*; left=%s, right=%s]\n", e.Lhs.String(), e.Rhs.String())
}
func (e DivisionExpression) String() string {
	return fmt.Sprintf("[/; left=%s, right=%s]\n", e.Lhs.String(), e.Rhs.String())
}
