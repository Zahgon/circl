package xwing

import (
	"github.com/cloudflare/circl/kem"
)

// This file contains the boilerplate code to connect X-Wing to the
// generic KEM API.

// Returns the generic KEM interface for  X-Wing PQ/T hybrid KEM.
func Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }

type scheme struct{}

func (scheme) Name() string               { _ = "STUB: not implemented"; return "" }
func (scheme) PublicKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (scheme) PrivateKeySize() int        { _ = "STUB: not implemented"; return 0 }
func (scheme) SeedSize() int              { _ = "STUB: not implemented"; return 0 }
func (scheme) EncapsulationSeedSize() int { _ = "STUB: not implemented"; return 0 }
func (scheme) SharedKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (scheme) CiphertextSize() int        { _ = "STUB: not implemented"; return 0 }
func (*PrivateKey) Scheme() kem.Scheme    { _ = "STUB: not implemented"; return *new(kem.Scheme) }
func (*PublicKey) Scheme() kem.Scheme     { _ = "STUB: not implemented"; return *new(kem.Scheme) }

func (sch scheme) Encapsulate(pk kem.PublicKey) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (scheme) EncapsulateDeterministically(
	pk kem.PublicKey, seed []byte,
) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (scheme) UnmarshalBinaryPublicKey(buf []byte) (kem.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), nil
}

func (scheme) UnmarshalBinaryPrivateKey(buf []byte) (kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PrivateKey), nil
}

func (sk *PrivateKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (sk *PrivateKey) Equal(other kem.PrivateKey) bool { _ = "STUB: not implemented"; return false }

func (sk *PrivateKey) Public() kem.PublicKey { _ = "STUB: not implemented"; return *new(kem.PublicKey) }

func (pk *PublicKey) Equal(other kem.PublicKey) bool { _ = "STUB: not implemented"; return false }

func (pk *PublicKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (scheme) DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey)
}

func (scheme) GenerateKeyPair() (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (scheme) Decapsulate(sk kem.PrivateKey, ct []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
