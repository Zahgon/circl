package group

import (
	_ "crypto/sha512"
	"io"
	"math/big"

	r255 "github.com/bwesterb/go-ristretto"
	"golang.org/x/crypto/cryptobyte"
)

// Ristretto255 is a quotient group generated from the edwards25519 curve.
var Ristretto255 Group = ristrettoGroup{}

type ristrettoGroup struct{}

func (g ristrettoGroup) String() string { _ = "STUB: not implemented"; return "" }

func (g ristrettoGroup) Params() *Params { _ = "STUB: not implemented"; return nil }

type ristrettoElement struct {
	p r255.Point
}

type ristrettoScalar struct {
	s r255.Scalar
}

func (g ristrettoGroup) NewElement() Element { _ = "STUB: not implemented"; return *new(Element) }

func (g ristrettoGroup) NewScalar() Scalar { _ = "STUB: not implemented"; return *new(Scalar) }

func (g ristrettoGroup) Identity() Element { _ = "STUB: not implemented"; return *new(Element) }

func (g ristrettoGroup) Generator() Element { _ = "STUB: not implemented"; return *new(Element) }

func (g ristrettoGroup) RandomElement(r io.Reader) Element {
	_ = "STUB: not implemented"
	return *new(Element)
}

func (g ristrettoGroup) RandomScalar(io.Reader) Scalar {
	_ = "STUB: not implemented"
	return *new(Scalar)
}

func (g ristrettoGroup) RandomNonZeroScalar(io.Reader) Scalar {
	_ = "STUB: not implemented"
	return *new(Scalar)
}

func (g ristrettoGroup) HashToElementNonUniform(b, dst []byte) Element {
	_ = "STUB: not implemented"
	return *new(Element)
}

func (g ristrettoGroup) HashToElement(msg, dst []byte) Element {
	_ = "STUB: not implemented"
	// Compliant with draft-irtf-cfrg-hash-to-curve.
	// Appendix B - Hashing to ristretto255
	// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-hash-to-curve-14#appendix-B
	// SuiteID: ristretto255_XMD:SHA-512_R255MAP_RO_
	return *new(Element)
}

func (g ristrettoGroup) HashToScalar(msg, dst []byte) Scalar {
	_ = "STUB: not implemented"
	// Adapted to be compliant with draft-irtf-cfrg-voprf
	// Section 4.1.1 - OPRF(ristretto255, SHA-512)
	// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-voprf-09#section-4.1.1
	return *new(Scalar)
}

func (e *ristrettoElement) Group() Group { _ = "STUB: not implemented"; return *new(Group) }

func (e *ristrettoElement) String() string { _ = "STUB: not implemented"; return "" }

func (e *ristrettoElement) IsIdentity() bool { _ = "STUB: not implemented"; return false }

func (e *ristrettoElement) IsEqual(x Element) bool { _ = "STUB: not implemented"; return false }

func (e *ristrettoElement) Set(x Element) Element { _ = "STUB: not implemented"; return *new(Element) }

func (e *ristrettoElement) Copy() Element { _ = "STUB: not implemented"; return *new(Element) }

func (e *ristrettoElement) CMov(v int, x Element) Element {
	_ = "STUB: not implemented"
	return *new(Element)
}

func (e *ristrettoElement) CSelect(v int, x Element, y Element) Element {
	_ = "STUB: not implemented"
	return *new(Element)
}

func (e *ristrettoElement) Add(x Element, y Element) Element {
	_ = "STUB: not implemented"
	return *new(Element)
}

func (e *ristrettoElement) Dbl(x Element) Element { _ = "STUB: not implemented"; return *new(Element) }

func (e *ristrettoElement) Neg(x Element) Element { _ = "STUB: not implemented"; return *new(Element) }

func (e *ristrettoElement) Mul(x Element, y Scalar) Element {
	_ = "STUB: not implemented"
	return *new(Element)
}

func (e *ristrettoElement) MulGen(x Scalar) Element {
	_ = "STUB: not implemented"
	return *new(Element)
}

func (e *ristrettoElement) MarshalBinaryCompress() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ristrettoElement) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ristrettoElement) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ristrettoScalar) Group() Group   { _ = "STUB: not implemented"; return *new(Group) }
func (s *ristrettoScalar) String() string { _ = "STUB: not implemented"; return "" }
func (s *ristrettoScalar) SetUint64(n uint64) Scalar {
	_ = "STUB: not implemented"
	return *new(Scalar)
}
func (s *ristrettoScalar) SetBigInt(x *big.Int) Scalar {
	_ = "STUB: not implemented"
	return *new(Scalar)
}
func (s *ristrettoScalar) IsZero() bool          { _ = "STUB: not implemented"; return false }
func (s *ristrettoScalar) IsEqual(x Scalar) bool { _ = "STUB: not implemented"; return false }

func (s *ristrettoScalar) Set(x Scalar) Scalar { _ = "STUB: not implemented"; return *new(Scalar) }

func (s *ristrettoScalar) Copy() Scalar { _ = "STUB: not implemented"; return *new(Scalar) }

func (s *ristrettoScalar) CMov(v int, x Scalar) Scalar {
	_ = "STUB: not implemented"
	return *new(Scalar)
}

func (s *ristrettoScalar) CSelect(v int, x Scalar, y Scalar) Scalar {
	_ = "STUB: not implemented"
	return *new(Scalar)
}

func (s *ristrettoScalar) Add(x Scalar, y Scalar) Scalar {
	_ = "STUB: not implemented"
	return *new(Scalar)
}

func (s *ristrettoScalar) Sub(x Scalar, y Scalar) Scalar {
	_ = "STUB: not implemented"
	return *new(Scalar)
}

func (s *ristrettoScalar) Mul(x Scalar, y Scalar) Scalar {
	_ = "STUB: not implemented"
	return *new(Scalar)
}

func (s *ristrettoScalar) Neg(x Scalar) Scalar { _ = "STUB: not implemented"; return *new(Scalar) }

func (s *ristrettoScalar) Inv(x Scalar) Scalar { _ = "STUB: not implemented"; return *new(Scalar) }

func (s *ristrettoScalar) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ristrettoScalar) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s *ristrettoScalar) Marshal(b *cryptobyte.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ristrettoScalar) Unmarshal(str *cryptobyte.String) bool {
	_ = "STUB: not implemented"
	return false
}
