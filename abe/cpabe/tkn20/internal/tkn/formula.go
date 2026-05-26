package tkn

import (
	"io"
)

const (
	Andgate = iota
	Orgate
)

// Gate is a Gate in a monotone boolean formula.
type Gate struct {
	Class int // either Andgate or Orgate
	In0   int // numbering of wires
	In1   int
	Out   int
}

func (g Gate) operator() string { _ = "STUB: not implemented"; return "" }

// Formula represents a monotone boolean circuit with Inputs not
// repeated.  The representation is as follows: for n Gates there n+1
// input wires, 1 output Wire, and n-1 intermediate wires.  That's
// because there are 2n Inputs to all Gates and n outputs since every
// Gate is 2:1.
//
// The wires are conceptually in an array. Wires 0 through n are
// the input wires, and Wire 2n is the output Wire. If there are wires
// between n and 2n they are intermediate wires.
//
// All intermediate and input wires must be used exactly once as Inputs.
type Formula struct {
	Gates []Gate
}

func (g Gate) Equal(g2 Gate) bool { _ = "STUB: not implemented"; return false }

func (f *Formula) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Formula) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (f *Formula) wellformed() error {
	_ = "STUB: not implemented"
	// Check every Wire used once
	return nil
}

// n+1 already, n-1 intermediates

// Sort the Gates so that Inputs are set before outputs.
func (f *Formula) toposort() error { _ = "STUB: not implemented"; return nil }

// Intermediate wires are indexed after subtracting n+1
// the Gate that sets this Wire
// the Gate that uses this intermediate Wire.
// the number of Inputs no yet output

// No Gate uses the output as input

// Given a set of possible Inputs (not necessarily in order!)
// return a subset that satisfy the formula with no extras.
func (f *Formula) satisfaction(available []match) ([]match, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// share distributes an input into shares for a secret sharing system
// for the formula: the original vector can be recovered from shares
// that satisfy the formula, by adding them all up.
func (f *Formula) share(rand io.Reader, k *matrixZp) ([]*matrixZp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reverse order: we want to set the share of the output ahead of the Inputs

// insertAnd adds and Gate for a new input
func (f *Formula) insertAnd() Formula {
	_ = "STUB: not implemented"
	// Let n=3
	// The old Inputs are 0, 1, 2, 3.
	// Old intermediates 4, 5,
	// Old output 6.
	// The old Inputs are Inputs 0,1,2,3 and new input 4
	// Intermediates are all shifted up by 1: 5, 6
	// Old output is also shifted up but is the intermediate 7
	// New output 8.
	return *new(Formula)
}

// if there were zero gates, then In0 = 0, In1 = 1, Out = 2

func (f *Formula) Equal(g Formula) bool { _ = "STUB: not implemented"; return false }
