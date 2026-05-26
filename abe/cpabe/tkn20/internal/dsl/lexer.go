package dsl

var keywords = map[string]string{
	"and": And,
	"or":  Or,
	"not": Not,
}

type Lexer struct {
	source   string
	tokens   []Token
	start    int
	curr     int
	line     int
	hadError bool
}

func newLexer(source string) Lexer { _ = "STUB: not implemented"; return *new(Lexer) }

func (l *Lexer) scanTokens() error { _ = "STUB: not implemented"; return nil }

func (l *Lexer) addToken(tokenType string) { _ = "STUB: not implemented"; return }

func (l *Lexer) identifier() { _ = "STUB: not implemented"; return }
