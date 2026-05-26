package ff

// Fp12Size is the length in bytes of an Fp12 element.
const Fp12Size = 2 * Fp6Size

// Fp12 represents an element of the field Fp12 = Fp6[w]/(w^2-v)., where v in Fp6.
type Fp12 [2]Fp6

func (z Fp12) String() string      { _ = "STUB: not implemented"; return "" }
func (z *Fp12) SetOne()            { _ = "STUB: not implemented"; return }
func (z Fp12) IsZero() int         { _ = "STUB: not implemented"; return 0 }
func (z Fp12) IsEqual(x *Fp12) int { _ = "STUB: not implemented"; return 0 }
func (z *Fp12) MulBeta()           { _ = "STUB: not implemented"; return }
func (z *Fp12) Frob(x *Fp12)       { _ = "STUB: not implemented"; return }
func (z *Fp12) Cjg()               { _ = "STUB: not implemented"; return }
func (z *Fp12) Neg()               { _ = "STUB: not implemented"; return }
func (z *Fp12) Add(x, y *Fp12)     { _ = "STUB: not implemented"; return }
func (z *Fp12) Sub(x, y *Fp12)     { _ = "STUB: not implemented"; return }
func (z *Fp12) Mul(x, y *Fp12)     { _ = "STUB: not implemented"; return }

func (z *Fp12) Sqr(x *Fp12) { _ = "STUB: not implemented"; return }

func (z *Fp12) Inv(x *Fp12) { _ = "STUB: not implemented"; return }

func (z *Fp12) CMov(x, y *Fp12, b int) { _ = "STUB: not implemented"; return }

// Exp calculates z=x^n, where n is the exponent in big-endian order.
func (z *Fp12) Exp(x *Fp12, n []byte) { _ = "STUB: not implemented"; return }

func (z *Fp12) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (z Fp12) MarshalBinary() (b []byte, e error) { _ = "STUB: not implemented"; return nil, nil }

// frob12W1 is Fp2 = [toMont(frob12W1_0), toMont(frob12W1_1) ], where
//
//	frob12W1_0 = 0x1904d3bf02bb0667c231beb4202c0d1f0fd603fd3cbd5f4f7b2443d784bab9c4f67ea53d63e7813d8d0775ed92235fb8
//	frob12W1_1 = 0xfc3e2b36c4e03288e9e902231f9fb854a14787b6c7b36fec0c8ec971f63c5f282d5ac14d6c7ec22cf78a126ddc4af3
var frob12W1 = Fp2{
	Fp{fpMont{ // (little-endian)
		0x07089552b319d465, 0xc6695f92b50a8313, 0x97e83cccd117228f,
		0xa35baecab2dc29ee, 0x1ce393ea5daace4d, 0x08f2220fb0fb66eb,
	}},
	Fp{fpMont{ // (little-endian)
		0xb2f66aad4ce5d646, 0x5842a06bfc497cec, 0xcf4895d42599d394,
		0xc11b9cba40a8e8d0, 0x2e3813cbe5a0de89, 0x110eefda88847faf,
	}},
}
