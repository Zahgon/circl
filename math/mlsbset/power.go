package mlsbset

// Power is a valid exponent produced by the MLSBSet encoding algorithm.
type Power struct {
	set Encoder // parameters of code.
	s   []int32 // set of signs.
	b   []int32 // set of digits.
	c   int     // carry is {0,1}.
}

// Exp is calculates x^k, where x is a predetermined element of a group G.
func (p *Power) Exp(G Group) EltG { _ = "STUB: not implemented"; return *new(EltG) }

// Digit returns the (v,e)-th digit and its sign.
func (p *Power) Digit(v, e uint) (sgn, dig int32) { _ = "STUB: not implemented"; return 0, 0 }

// bit returns the (w,v,e)-th bit of the code.
func (p *Power) bit(w, v, e uint) int32 { _ = "STUB: not implemented"; return 0 }

func (p *Power) String() string { _ = "STUB: not implemented"; return "" }
