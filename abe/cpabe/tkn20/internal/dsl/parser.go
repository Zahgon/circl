package dsl

type Parser struct {
	tokens   []Token
	curr     int
	wires    map[attr]attrValue
	gates    []gate
	negative bool
}

func newParser(tokens []Token) Parser { _ = "STUB: not implemented"; return *new(Parser) }

func (p *Parser) parse() (Ast, error) { _ = "STUB: not implemented"; return *new(Ast), nil }

func (p *Parser) expression() (Expr, error) { _ = "STUB: not implemented"; return *new(Expr), nil }

func (p *Parser) or() (Expr, error) { _ = "STUB: not implemented"; return *new(Expr), nil }

func (p *Parser) and() (Expr, error) { _ = "STUB: not implemented"; return *new(Expr), nil }

func (p *Parser) not() (Expr, error) { _ = "STUB: not implemented"; return *new(Expr), nil }

func (p *Parser) primary() (Expr, error) { _ = "STUB: not implemented"; return *new(Expr), nil }

func extractAttr(expr Expr) attr { _ = "STUB: not implemented"; return *new(attr) }
