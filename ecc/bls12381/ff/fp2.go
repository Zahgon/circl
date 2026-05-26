package ff

// Fp2Size is the length in bytes of an Fp2 element.
const Fp2Size = 2 * FpSize

type Fp2 [2]Fp

func (z Fp2) String() string { _ = "STUB: not implemented"; return "" }
func (z *Fp2) SetOne()       { _ = "STUB: not implemented"; return }

// IsNegative returns 1 if z is lexicographically larger than -z; otherwise returns 0.
func (z Fp2) IsNegative() int    { _ = "STUB: not implemented"; return 0 }
func (z Fp2) IsZero() int        { _ = "STUB: not implemented"; return 0 }
func (z Fp2) IsEqual(x *Fp2) int { _ = "STUB: not implemented"; return 0 }
func (z *Fp2) MulBeta()          { _ = "STUB: not implemented"; return }
func (z *Fp2) Frob(x *Fp2)       { _ = "STUB: not implemented"; return }
func (z *Fp2) Cjg()              { _ = "STUB: not implemented"; return }
func (z *Fp2) Neg()              { _ = "STUB: not implemented"; return }
func (z *Fp2) Add(x, y *Fp2)     { _ = "STUB: not implemented"; return }
func (z *Fp2) Sub(x, y *Fp2)     { _ = "STUB: not implemented"; return }
func (z *Fp2) Mul(x, y *Fp2)     { _ = "STUB: not implemented"; return }

func (z *Fp2) Sqr(x *Fp2) { _ = "STUB: not implemented"; return }

func (z *Fp2) Inv(x *Fp2) { _ = "STUB: not implemented"; return }

func (z Fp2) Sgn0() int { _ = "STUB: not implemented"; return 0 }

func (z *Fp2) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (z Fp2) MarshalBinary() (b []byte, e error) { _ = "STUB: not implemented"; return nil, nil }

// SetString reconstructs a Fp2 element as s0+s1*i, where s0 and s1 are numeric
// strings from 0 to FpOrder-1.
func (z *Fp2) SetString(s0, s1 string) (err error) { _ = "STUB: not implemented"; return nil }

func (z *Fp2) CMov(x, y *Fp2, b int) { _ = "STUB: not implemented"; return }

// ExpVarTime calculates z=x^n, where n is the exponent in big-endian order.
func (z *Fp2) ExpVarTime(x *Fp2, n []byte) { _ = "STUB: not implemented"; return }

// Sqrt returns 1 and sets z=sqrt(x) only if x is a quadratic-residue; otherwise, returns 0 and z is unmodified.
func (z *Fp2) Sqrt(x *Fp2) int {
	_ = "STUB: not implemented"
	// "Square-root for q = p^2 = 9 (mod 16)" Appendix I.3 of Hashing to elliptic curves.
	// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-hash-to-curve-11#appendix-I.3
	return 0
}

var fp2SqrtConst = struct {
	// "Square-root for q = p^2 = 9 (mod 16)" Appendix I.3 of Hashing to elliptic curves.
	// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-hash-to-curve-11#appendix-I.3
	c1 Fp2      // c1 = sqrt( -1) = u
	c2 Fp2      // c2 = sqrt( c1)
	c3 Fp2      // c3 = sqrt(-c1)
	c4 [95]byte // c4 = (p^2 + 7) / 16 (big-endian)
}{
	c1: Fp2{ // (little-endian)
		Fp{fpMont{}},
		Fp{fpMont{0x760900000002fffd, 0xebf4000bc40c0002, 0x5f48985753c758ba, 0x77ce585370525745, 0x5c071a97a256ec6d, 0x15f65ec3fa80e493}},
	},
	c2: Fp2{ // (little-endian)
		Fp{fpMont{0x3e2f585da55c9ad1, 0x4294213d86c18183, 0x382844c88b623732, 0x92ad2afd19103e18, 0x1d794e4fac7cf0b9, 0x0bd592fc7d825ec8}},
		Fp{fpMont{0x7bcfa7a25aa30fda, 0xdc17dec12a927e7c, 0x2f088dd86b4ebef1, 0xd1ca2087da74d4a7, 0x2da2596696cebc1d, 0x0e2b7eedbbfd87d2}},
	},
	c3: Fp2{ // (little-endian)
		Fp{fpMont{0x7bcfa7a25aa30fda, 0xdc17dec12a927e7c, 0x2f088dd86b4ebef1, 0xd1ca2087da74d4a7, 0x2da2596696cebc1d, 0x0e2b7eedbbfd87d2}},
		Fp{fpMont{0x7bcfa7a25aa30fda, 0xdc17dec12a927e7c, 0x2f088dd86b4ebef1, 0xd1ca2087da74d4a7, 0x2da2596696cebc1d, 0x0e2b7eedbbfd87d2}},
	},
	c4: [95]byte{ // (big-endian)
		0x2a, 0x43, 0x7a, 0x4b, 0x8c, 0x35, 0xfc, 0x74, 0xbd, 0x27, 0x8e, 0xaa,
		0x22, 0xf2, 0x5e, 0x9e, 0x2d, 0xc9, 0x0e, 0x50, 0xe7, 0x04, 0x6b, 0x46,
		0x6e, 0x59, 0xe4, 0x93, 0x49, 0xe8, 0xbd, 0x05, 0x0a, 0x62, 0xcf, 0xd1,
		0x6d, 0xdc, 0xa6, 0xef, 0x53, 0x14, 0x93, 0x30, 0x97, 0x8e, 0xf0, 0x11,
		0xd6, 0x86, 0x19, 0xc8, 0x61, 0x85, 0xc7, 0xb2, 0x92, 0xe8, 0x5a, 0x87,
		0x09, 0x1a, 0x04, 0x96, 0x6b, 0xf9, 0x1e, 0xd3, 0xe7, 0x1b, 0x74, 0x31,
		0x62, 0xc3, 0x38, 0x36, 0x21, 0x13, 0xcf, 0xd7, 0xce, 0xd6, 0xb1, 0xd7,
		0x63, 0x82, 0xea, 0xb2, 0x6a, 0xa0, 0x00, 0x01, 0xc7, 0x18, 0xe4,
	},
}
