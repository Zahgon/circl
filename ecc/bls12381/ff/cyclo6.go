package ff

// Cyclo6 represents an element of the 6th cyclotomic group.
type Cyclo6 Fp12

func (z Cyclo6) String() string           { _ = "STUB: not implemented"; return "" }
func (z Cyclo6) IsEqual(x *Cyclo6) int    { _ = "STUB: not implemented"; return 0 }
func (z Cyclo6) IsIdentity() int          { _ = "STUB: not implemented"; return 0 }
func (z *Cyclo6) Frob(x *Cyclo6)          { _ = "STUB: not implemented"; return }
func (z *Cyclo6) Mul(x, y *Cyclo6)        { _ = "STUB: not implemented"; return }
func (z *Cyclo6) Inv(x *Cyclo6)           { _ = "STUB: not implemented"; return }
func (z *Cyclo6) exp(x *Cyclo6, n []byte) { _ = "STUB: not implemented"; return }
func (z *Cyclo6) Sqr(x *Cyclo6) {
	_ = "STUB: not implemented"
	// Method of Granger-Scott.
	// Page 7 of "Faster Squaring in the Cyclotomic Subgroup of Sixth Degree Extensions"
	// https://www.iacr.org/archive/pkc2010/60560212/60560212.pdf
	return
}

// PowToX computes z = x^paramX, where paramX is the parameter of the BLS curve.
func (z *Cyclo6) PowToX(x *Cyclo6) { _ = "STUB: not implemented"; return }

// paramX is -2 ^ 63 - 2 ^ 62 - 2 ^ 60 - 2 ^ 57 - 2 ^ 48 - 2 ^ 16

// EasyExponentiation calculates g = f^(p^6-1)(p^2+1), where g becomes an
// element of the 6-th cyclotomic group.
func EasyExponentiation(g *Cyclo6, f *Fp12) { _ = "STUB: not implemented"; return }

// p = f^(p)
// p = f^(p^2)
// t0 = f^(p^2 + 1)
// t1 = f^(p^2 + 1)*(p)
// t1 = f^(p^2 + 1)*(p^2)
// t1 = f^(p^2 + 1)*(p^3)
// t1 = f^(p^2 + 1)*(p^4)
// t1 = f^(p^2 + 1)*(p^5)
// t1 = f^(p^2 + 1)*(p^6)
// t0 = f^-(p^2 + 1)
// t0 = f^(p^2 + 1)*(p^6 - 1)

// HardExponentiation calculates u = g^(Cy_6(p)/r), where u is a root of unity.
func HardExponentiation(u *URoot, g *Cyclo6) { _ = "STUB: not implemented"; return }

// _g = g^-1
// g3 = g^2
// g3 = g^3
// t0 = g^x
// t0 = g^(x-1)
// t1 = g^(x-1)*x
// t0 = g^-(x-1)
// a3 = g^(x-1)*(x-1)
// a2 = a3*p
// a1 = a2*p = a3*p^2
// t0 = -a3
// a1 = a3*p^2-a3
// a0 = a3*p^3-a3*p
// a0 = a3*p^3-a3*p+3

// c = g^(a3*x)
// c = g^(a3*x+a2)
// c = g^(a3*x+a2)*x = g^(a3*x^2+a2*x)
// c = g^(a3*x^2+a2*x+a1)
// c = g^(a3*x^2+a2*x+a1)*x = g^(a3*x^3+a2*x^2+a1*x)
// c = g^(a3*x^3+a2*x^2+a1*x+a0)
