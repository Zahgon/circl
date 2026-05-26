package oprf

import (
	"io"

	"github.com/cloudflare/circl/group"
)

type PrivateKey struct {
	p   params
	k   group.Scalar
	pub *PublicKey
}

type PublicKey struct {
	p params
	e group.Element
}

func (k *PrivateKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (k *PublicKey) MarshalBinary() ([]byte, error)  { _ = "STUB: not implemented"; return nil, nil }

func (k *PrivateKey) UnmarshalBinary(s Suite, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *PublicKey) UnmarshalBinary(s Suite, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *PrivateKey) Public() *PublicKey { _ = "STUB: not implemented"; return nil }

// GenerateKey generates a private key compatible with the suite.
func GenerateKey(s Suite, rnd io.Reader) (*PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeriveKey generates a private key from a 32-byte seed and an optional info string.
func DeriveKey(s Suite, mode Mode, seed, info []byte) (*PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
