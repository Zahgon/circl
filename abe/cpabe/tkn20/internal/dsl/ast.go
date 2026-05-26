package dsl

import (
	"github.com/cloudflare/circl/abe/cpabe/tkn20/internal/tkn"
)

var operators = map[string]int{
	"and": tkn.Andgate,
	"or":  tkn.Orgate,
}

type attrValue struct {
	value    string
	positive bool
}

type attr struct {
	key string
	id  int
}

type gate struct {
	op  string
	in1 attr
	in2 attr
	out attr
}

type Ast struct {
	wires map[attr]attrValue
	gates []gate
}

func (t *Ast) RunPasses() (*tkn.Policy, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Ast) hashAttrValues() ([]tkn.Wire, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Ast) transformGates() ([]tkn.Gate, error) { _ = "STUB: not implemented"; return nil, nil }
