package consts

import "math"

type Expression interface {
	Eval() float64
}
type NumberExpression struct {
	Value float64
}
type PlusExpression struct {
	Lhs Expression `json:"plus-left"`
	Rhs Expression `json:"plus-right"`
}
type MinusExpression struct {
	Lhs Expression `json:"minus-left"`
	Rhs Expression `json:"minus-right"`
}
type MultiplicationExpression struct {
	Lhs Expression `json:"mul-left"`
	Rhs Expression `json:"mul-right"`
}
type DivisionExpression struct {
	Lhs Expression `json:"div-left"`
	Rhs Expression `json:"div-right"`
}
type FactorExpression struct {
	Lhs Expression `json:"fac-left"`
	Rhs Expression `json:"fac-right"`
}
type EofExpression struct{}

func (e EofExpression) Eval() any {
	return nil
}
func (e NumberExpression) Eval() float64 {
	return e.Value
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
func (e FactorExpression) Eval() float64 {
	return math.Pow(e.Lhs.Eval(), e.Rhs.Eval())
}
