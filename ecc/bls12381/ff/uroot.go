package ff

// URootSize is the length in bytes of a root of unit.
const URootSize = Fp12Size

// URoot represents an n-th root of unit, that is an element x in Cyclo6 such
// that x^n=1, where n = ScalarOrder().
type URoot Cyclo6

func (z URoot) String() string                  { _ = "STUB: not implemented"; return "" }
func (z *URoot) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }
func (z URoot) MarshalBinary() ([]byte, error)  { _ = "STUB: not implemented"; return nil, nil }
func (z *URoot) SetIdentity()                   { _ = "STUB: not implemented"; return }
func (z URoot) IsEqual(x *URoot) int            { _ = "STUB: not implemented"; return 0 }
func (z URoot) IsIdentity() int                 { _ = "STUB: not implemented"; return 0 }
func (z *URoot) Exp(x *URoot, n []byte)         { _ = "STUB: not implemented"; return }
func (z *URoot) Mul(x, y *URoot)                { _ = "STUB: not implemented"; return }
func (z *URoot) Sqr(x *URoot)                   { _ = "STUB: not implemented"; return }
func (z *URoot) Inv(x *URoot)                   { _ = "STUB: not implemented"; return }
