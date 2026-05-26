package ff

// Fp4Size is the size of an Fp4 element
const Fp4Size = 4 * FpSize

// Fp4 is obtained by adjoining t, the square root of u+1 to Fp2
type Fp4 [2]Fp2

func (z Fp4) String() string { _ = "STUB: not implemented"; return "" }

func (z *Fp4) SetOne() { _ = "STUB: not implemented"; return }

func (z *Fp4) IsZero() int { _ = "STUB: not implemented"; return 0 }

func (z *Fp4) IsEqual(x *Fp4) int { _ = "STUB: not implemented"; return 0 }

func (z *Fp4) Cjg() { _ = "STUB: not implemented"; return }

func (z *Fp4) Neg() { _ = "STUB: not implemented"; return }

func (z *Fp4) Add(x *Fp4, y *Fp4) { _ = "STUB: not implemented"; return }

func (z *Fp4) Sub(x *Fp4, y *Fp4) { _ = "STUB: not implemented"; return }

func (z *Fp4) Mul(x *Fp4, y *Fp4) { _ = "STUB: not implemented"; return }

// k is x0y1+x1y0 computed as (x0+x1)(y0+y1)-x0y0-x1y1

// Multiply x1y1 by u+1

func (z *Fp4) Sqr(x *Fp4) { _ = "STUB: not implemented"; return }

// Multiplying x1s by u+1

func (z *Fp4) Inv(x *Fp4) {
	_ = "STUB: not implemented"
	// Compute the inverse via conjugation
	return
}

func (z *Fp4) mulSubfield(x *Fp4, y *Fp2) { _ = "STUB: not implemented"; return }

func (z *Fp4) mulT(x *Fp4) { _ = "STUB: not implemented"; return }
