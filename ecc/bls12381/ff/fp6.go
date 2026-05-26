package ff

// Fp6Size is the length in bytes of an Fp6 element.
const Fp6Size = 3 * Fp2Size

type Fp6 [3]Fp2

func (z Fp6) String() string     { _ = "STUB: not implemented"; return "" }
func (z *Fp6) SetOne()           { _ = "STUB: not implemented"; return }
func (z Fp6) IsZero() int        { _ = "STUB: not implemented"; return 0 }
func (z Fp6) IsEqual(x *Fp6) int { _ = "STUB: not implemented"; return 0 }

func (z *Fp6) Neg()          { _ = "STUB: not implemented"; return }
func (z *Fp6) Add(x, y *Fp6) { _ = "STUB: not implemented"; return }
func (z *Fp6) Sub(x, y *Fp6) { _ = "STUB: not implemented"; return }
func (z *Fp6) MulBeta()      { _ = "STUB: not implemented"; return }

func (z *Fp6) Mul(x, y *Fp6) {
	_ = "STUB: not implemented"
	// https://ia.cr/2006/224 (Sec3.1)
	//
	//	z = x*y mod (v^3-B)
	//
	// | v^4 | v^3 ||  v^2  |  v^1  |  v^0  |
	// |-----|-----||-------|-------|-------|
	// |     |     ||  -c2  |  -c1  |  +c0  |
	// |     | -c2 ||  +c1  |  -c0  |       |
	// | +c2 | -c1 ||  -c0  |       |       |
	// |     | +c5 ||  +c4  |  +c3  |       |
	// |-----|-----||-------|-------|-------|
	// |     |     ||       | B(+c2)| B(-c2)|
	// |     |     ||       |       | B(-c1)|
	// |     |     ||       |       | B(+c5)|
	return
}

// c4+c1
// c4+c1-c0
// z2 = c4+c1-c0-c2
// Bc2
// Bc2-c0
// c3-c1
// z1 = Bc2-c0+c3-c1
// c5-c1
// B(c5-c1)
// z0 = B(c5-c1)-Bc2+c0 = B(c5-c1-c2)+c0

func (z *Fp6) Sqr(x *Fp6) {
	_ = "STUB: not implemented"
	//	z = x^2 mod (v^3-B)
	//
	// z0 = B(2x1*x2) + x0^2
	// z1 = B(x2^2) + 2x0*x1
	// z2 = 2x0*x2 + x1^2
	return
}

// 2c5
// 2c3
// 2c3+2c5
// 2c3+2c5+c4
// 2c3+2c5+c4-c0
// z2 = 2c3+2c5+c4-c0-c2
// B(2c5)
// z0 = B(2c5)+c0
// B(c2)
// z1 = B(c2)+2c3

func (z *Fp6) Inv(x *Fp6) { _ = "STUB: not implemented"; return }

// c0 = aL^2 - B(aM*AH)

// c1 = B(aH^2) - aL*AM
// c1 = aM^2 - aL*AH

// den = B(aL*c2 + aM*c1) + aLc0
// z0 = c0/den
// z1 = c1/den
// z2 = c2/den

func (z *Fp6) Frob(x *Fp6) { _ = "STUB: not implemented"; return }

func (z *Fp6) CMov(x, y *Fp6, b int) { _ = "STUB: not implemented"; return }

func (z Fp6) MarshalBinary() (b []byte, e error) { _ = "STUB: not implemented"; return nil, nil }

func (z *Fp6) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

var (
	// frob6V1 is toMont(v) = 2**384 * v mod fpPrime, where
	// v = 0x1a0111ea397fe699ec02408663d4de85aa0d857d89759ad4897d29650fb85f9b409427eb4f49fffd8bfd00000000aaac
	frob6V1 = Fp{fpMont{
		0xcd03c9e48671f071, 0x5dab22461fcda5d2, 0x587042afd3851b95,
		0x8eb60ebe01bacb9e, 0x03f97d6e83d050d2, 0x18f0206554638741,
	}}

	// frob6V2 is toMont(v) = 2**384 * v mod fpPrime, where
	// v = 0x1a0111ea397fe699ec02408663d4de85aa0d857d89759ad4897d29650fb85f9b409427eb4f49fffd8bfd00000000aaad
	frob6V2 = Fp{fpMont{
		0x890dc9e4867545c3, 0x2af322533285a5d5, 0x50880866309b7e2c,
		0xa20d1b8c7e881024, 0x14e4f04fe2db9068, 0x14e56d3f1564853a,
	}}
)
