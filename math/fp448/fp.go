// Package fp448 provides prime field arithmetic over GF(2^448-2^224-1).
package fp448

// Size in bytes of an element.
const Size = 56

// Elt is a prime field element.
type Elt [Size]byte

func (e Elt) String() string { _ = "STUB: not implemented"; return "" }

// p is the prime modulus 2^448-2^224-1.
var p = Elt{
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
}

// P returns the prime modulus 2^448-2^224-1.
func P() Elt {
	_ = "STUB: not implemented"

	// ToBytes stores in b the little-endian byte representation of x.
	return *new(Elt)
}

func ToBytes(b []byte, x *Elt) error { _ = "STUB: not implemented"; return nil }

// IsZero returns true if x is equal to 0.
func IsZero(x *Elt) bool { _ = "STUB: not implemented"; return false }

// IsOne returns true if x is equal to 1.
func IsOne(x *Elt) bool { _ = "STUB: not implemented"; return false }

// SetOne assigns x=1.
func SetOne(x *Elt) {
	_ = "STUB: not implemented"

	// One returns the 1 element.
	return
}

func One() (x Elt) {
	_ = "STUB: not implemented"

	// Neg calculates z = -x.
	return *new(Elt)
}

func Neg(z, x *Elt) {
	_ = "STUB: not implemented"

	// Modp ensures that z is between [0,p-1].
	return
}

func Modp(z *Elt) {
	_ = "STUB: not implemented"

	// InvSqrt calculates z = sqrt(x/y) iff x/y is a quadratic-residue. If so,
	// isQR = true; otherwise, isQR = false, since x/y is a quadratic non-residue,
	// and z = sqrt(-x/y).
	return
}

func InvSqrt(z, x, y *Elt) (isQR bool) {
	_ = "STUB: not implemented"
	// First note that x^(2(k+1)) = x^(p-1)/2 * x = legendre(x) * x
	// so that's x if x is a quadratic residue and -x otherwise.
	// Next, y^(6k+3) = y^(4k+2) * y^(2k+1) = y^(p-1) * y^((p-1)/2) = legendre(y).
	// So the z we compute satisfies z^2 y = x^(2(k+1)) y^(6k+3) = legendre(x)*legendre(y).
	// Thus if x and y are quadratic residues, then z is indeed sqrt(x/y).
	return false
}

// x*y
// y^2
// x*y^3
// (x*y^3)^k
// z = x*y*(x*y^3)^k = x^(k+1) * y^(3k+1)

// Check if x/y is a quadratic residue
// z^2
// y*z^2
// y*z^2-x

// Inv calculates z = 1/x mod p.
func Inv(z, x *Elt) {
	_ = "STUB: not implemented"
	// Calculates z = x^(4k+1) = x^(p-3+1) = x^(p-2) = x^-1, where k = (p-3)/4.
	return
}

// t = x^k
// t = x^2k
// t = x^4k
// z = x^(4k+1)

// powPminus3div4 calculates z = x^k mod p, where k = (p-3)/4.
func powPminus3div4(z, x *Elt) { _ = "STUB: not implemented"; return }

// Cmov assigns y to x if n is 1.
func Cmov(x, y *Elt, n uint) {
	_ = "STUB: not implemented"

	// Cswap interchanges x and y if n is 1.
	return
}

func Cswap(x, y *Elt, n uint) {
	_ = "STUB: not implemented"

	// Add calculates z = x+y mod p.
	return
}

func Add(z, x, y *Elt) {
	_ = "STUB: not implemented"

	// Sub calculates z = x-y mod p.
	return
}

func Sub(z, x, y *Elt) {
	_ = "STUB: not implemented"

	// AddSub calculates (x,y) = (x+y mod p, x-y mod p).
	return
}

func AddSub(x, y *Elt) {
	_ = "STUB: not implemented"

	// Mul calculates z = x*y mod p.
	return
}

func Mul(z, x, y *Elt) {
	_ = "STUB: not implemented"

	// Sqr calculates z = x^2 mod p.
	return
}

func Sqr(z, x *Elt) { _ = "STUB: not implemented"; return }
