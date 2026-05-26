package hpke

import (
	"github.com/cloudflare/circl/kem"
)

type xKEM struct {
	dhKemBase
	size int
}

func (x xKEM) PrivateKeySize() int        { _ = "STUB: not implemented"; return 0 }
func (x xKEM) SeedSize() int              { _ = "STUB: not implemented"; return 0 }
func (x xKEM) CiphertextSize() int        { _ = "STUB: not implemented"; return 0 }
func (x xKEM) PublicKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (x xKEM) EncapsulationSeedSize() int { _ = "STUB: not implemented"; return 0 }

func (x xKEM) sizeDH() int { _ = "STUB: not implemented"; return 0 }
func (x xKEM) calcDH(dh []byte, sk kem.PrivateKey, pk kem.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Deterministically derives a keypair from a seed. If you're unsure,
// you're better off using GenerateKey().
//
// Panics if seed is not of length SeedSize().
func (x xKEM) DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey) {
	_ = "STUB: not implemented"
	// Implementation based on
	// https://www.ietf.org/archive/id/draft-irtf-cfrg-hpke-07.html#name-derivekeypair
	return *new(kem.PublicKey), *new(kem.PrivateKey)
}

func (x xKEM) GenerateKeyPair() (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (x xKEM) UnmarshalBinaryPrivateKey(data []byte) (kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PrivateKey), nil
}

func (x xKEM) UnmarshalBinaryPublicKey(data []byte) (kem.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), nil
}

type xKEMPubKey struct {
	scheme xKEM
	pub    []byte
}

func (k *xKEMPubKey) String() string                 { _ = "STUB: not implemented"; return "" }
func (k *xKEMPubKey) Scheme() kem.Scheme             { _ = "STUB: not implemented"; return *new(kem.Scheme) }
func (k *xKEMPubKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *xKEMPubKey) Equal(pk kem.PublicKey) bool { _ = "STUB: not implemented"; return false }

func (k *xKEMPubKey) validate() bool { _ = "STUB: not implemented"; return false }

type xKEMPrivKey struct {
	scheme xKEM
	priv   []byte
	pub    *xKEMPubKey
}

func (k *xKEMPrivKey) String() string                 { _ = "STUB: not implemented"; return "" }
func (k *xKEMPrivKey) Scheme() kem.Scheme             { _ = "STUB: not implemented"; return *new(kem.Scheme) }
func (k *xKEMPrivKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *xKEMPrivKey) Equal(pk kem.PrivateKey) bool { _ = "STUB: not implemented"; return false }

func (k *xKEMPrivKey) Public() kem.PublicKey { _ = "STUB: not implemented"; return *new(kem.PublicKey) }

func (k *xKEMPrivKey) validate() bool { _ = "STUB: not implemented"; return false }
