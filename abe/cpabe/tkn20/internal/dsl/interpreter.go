package dsl

type Interpreter struct {
	Literal
}

func (i *Interpreter) Evaluate(expr Expr) Literal { _ = "STUB: not implemented"; return *new(Literal) }

func (i *Interpreter) visitBinary(b Binary) { _ = "STUB: not implemented"; return }

func (i *Interpreter) visitUnary(u Unary) { _ = "STUB: not implemented"; return }

func (i *Interpreter) visitGrouping(g Grouping) { _ = "STUB: not implemented"; return }

func (i *Interpreter) visitLiteral(at Literal) { _ = "STUB: not implemented"; return }
