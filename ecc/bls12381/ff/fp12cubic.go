package ff

// Fp12Cubic represents elements of Fp4[w]/w^3-t
type Fp12Cubic [3]Fp4

// LineValue a represents a[0]+a[1]*w^2+a[2]*w^3, with all values in Fp2.
// This lets us shave off a number of Fp2 multiplications.
type LineValue [3]Fp2

func (z Fp12Cubic) String() string { _ = "STUB: not implemented"; return "" }

func (z Fp12Cubic) IsEqual(x *Fp12Cubic) int { _ = "STUB: not implemented"; return 0 }

func (z *Fp12Cubic) FromFp12(x *Fp12) {
	_ = "STUB: not implemented"
	// To understand this function, write everything in Fp2[w]/(w^6-(u+1)).
	// v = w^2
	// t = w^3
	return
}

// w^0
// w^1
// w^2
// w^3
// w^4
// w^5

func (z *Fp12) FromFp12Cubic(x *Fp12Cubic) {
	_ = "STUB: not implemented"
	// w^0
	return
}

// w^1
// w^2
// w^3
// w^4
// w^5

func (z *Fp12Cubic) Add(x *Fp12Cubic, y *Fp12Cubic) { _ = "STUB: not implemented"; return }

func (z *Fp12Cubic) SetOne() { _ = "STUB: not implemented"; return }

func (z *Fp12Cubic) Mul(x *Fp12Cubic, y *Fp12Cubic) {
	_ = "STUB: not implemented"
	// This is a Karatsuba like technique: compute x_iy_i, then
	// subtract the terms from expressions like (x0+x1)(y0+y1).
	// See Multiplication and Squaring in Pairing Friendly Fields
	// for more.
	return
}

func (z *Fp12Cubic) Sqr(x *Fp12Cubic) {
	_ = "STUB: not implemented"
	// The Chung-Hasan asymmetric squaring formula.
	// We keep the same notation as in Multiplication
	// and Squaring on Pairing-Friendly Fields to make
	// it easier to compare.
	return
}

// x0^2

// 2x0x1

// (x0-x1+x2)^2

// 2x1x2
// x2^2

func (z *Fp12Cubic) MulLine(x *Fp12Cubic, y *LineValue) {
	_ = "STUB: not implemented"
	// Values produced by evaluating a line function
	// do not have a linear term, and the quadratic term is
	// in Fp2.
	return
}

// We copy the Mul method, but remove any multiplies by y1,
// and strength reduce multiplies by y2.

func (z *LineValue) IsZero() int { _ = "STUB: not implemented"; return 0 }

func (z *LineValue) SetOne() { _ = "STUB: not implemented"; return }
