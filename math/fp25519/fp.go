// Package fp25519 provides prime field arithmetic over GF(2^255-19).
package fp25519

// Size in bytes of an element.
const Size = 32

// Elt is a prime field element.
type Elt [Size]byte

func (e Elt) String() string { _ = "STUB: not implemented"; return "" }

// p is the prime modulus 2^255-19.
var p = Elt{
	0xed, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f,
}

// P returns the prime modulus 2^255-19.
func P() Elt {
	_ = "STUB: not implemented"

	// ToBytes stores in b the little-endian byte representation of x.
	return *new(Elt)
}

func ToBytes(b []byte, x *Elt) error { _ = "STUB: not implemented"; return nil }

// IsZero returns true if x is equal to 0.
func IsZero(x *Elt) bool { _ = "STUB: not implemented"; return false }

// SetOne assigns x=1.
func SetOne(x *Elt) { _ = "STUB: not implemented"; return }

// Neg calculates z = -x.
func Neg(z, x *Elt) {
	_ = "STUB: not implemented"

	// InvSqrt calculates z = sqrt(x/y) iff x/y is a quadratic-residue, which is
	// indicated by returning isQR = true. Otherwise, when x/y is a quadratic
	// non-residue, z will have an undetermined value and isQR = false.
	return
}

func InvSqrt(z, x, y *Elt) (isQR bool) { _ = "STUB: not implemented"; return false }

// t0 = u*v
// t1 = v^2
// t2 = u*v^3
// t0 = v^4
// t1 = u*v^7

// z = xy^(p+3)/8 = xy^3*(xy^7)^(p-5)/8
// Checking whether y z^2 == x
// t0 = z^2
// t0 = yz^2
// t1 = t0-u
// t2 = t0+u

// z = z*sqrt(-1)

// Inv calculates z = 1/x mod p.
func Inv(z, x *Elt) { _ = "STUB: not implemented"; return }

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

func Sqr(z, x *Elt) {
	_ = "STUB: not implemented"

	// Modp ensures that z is between [0,p-1].
	return
}

func Modp(z *Elt) { _ = "STUB: not implemented"; return }
