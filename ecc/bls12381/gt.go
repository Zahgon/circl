package bls12381

import "github.com/cloudflare/circl/ecc/bls12381/ff"

// GtSize is the length in bytes of an element in Gt.
const GtSize = ff.URootSize

// Gt represents an element of the output (multiplicative) group of a pairing.
type Gt struct{ i ff.URoot }

func (z Gt) String() string                  { _ = "STUB: not implemented"; return "" }
func (z *Gt) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }
func (z Gt) MarshalBinary() ([]byte, error)  { _ = "STUB: not implemented"; return nil, nil }
func (z *Gt) SetIdentity()                   { _ = "STUB: not implemented"; return }
func (z Gt) IsEqual(x *Gt) bool              { _ = "STUB: not implemented"; return false }
func (z Gt) IsIdentity() bool                { _ = "STUB: not implemented"; return false }
func (z *Gt) Mul(x, y *Gt)                   { _ = "STUB: not implemented"; return }
func (z *Gt) Sqr(x *Gt)                      { _ = "STUB: not implemented"; return }
func (z *Gt) Inv(x *Gt)                      { _ = "STUB: not implemented"; return }

// Exp calculates z=x^n, where n is the exponent in big-endian order.
func (z *Gt) Exp(x *Gt, n *Scalar) { _ = "STUB: not implemented"; return }
