package consts

type Expression interface {
	Eval() float64
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
