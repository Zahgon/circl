package hpke

import (
	"crypto/ecdh"

	"github.com/cloudflare/circl/kem"
)

type shortKEM struct {
	dhKemBase
	ecdh.Curve
}

func (s shortKEM) PrivateKeySize() int        { _ = "STUB: not implemented"; return 0 }
func (s shortKEM) SeedSize() int              { _ = "STUB: not implemented"; return 0 }
func (s shortKEM) CiphertextSize() int        { _ = "STUB: not implemented"; return 0 }
func (s shortKEM) PublicKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (s shortKEM) EncapsulationSeedSize() int { _ = "STUB: not implemented"; return 0 }

func (s shortKEM) byteSize() int { _ = "STUB: not implemented"; return 0 }

func (s shortKEM) sizeDH() int { _ = "STUB: not implemented"; return 0 }
func (s shortKEM) calcDH(dh []byte, sk kem.PrivateKey, pk kem.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Deterministically derives a keypair from a seed. If you're unsure,
// you're better off using GenerateKey().
//
// Panics if seed is not of length SeedSize().
func (s shortKEM) DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey) {
	_ = "STUB: not implemented"
	// Implementation based on
	// https://www.ietf.org/archive/id/draft-irtf-cfrg-hpke-07.html#name-derivekeypair
	return *new(kem.PublicKey), *new(kem.PrivateKey)
}

func (s shortKEM) GenerateKeyPair() (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (s shortKEM) UnmarshalBinaryPrivateKey(data []byte) (kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PrivateKey), nil
}

func (s shortKEM) UnmarshalBinaryPublicKey(data []byte) (kem.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), nil
}

type shortKEMPubKey struct {
	scheme shortKEM
	pub    ecdh.PublicKey
}

func (k *shortKEMPubKey) String() string     { _ = "STUB: not implemented"; return "" }
func (k *shortKEMPubKey) Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }
func (k *shortKEMPubKey) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *shortKEMPubKey) Equal(pk kem.PublicKey) bool { _ = "STUB: not implemented"; return false }

type shortKEMPrivKey struct {
	scheme shortKEM
	priv   *ecdh.PrivateKey
}

func (k *shortKEMPrivKey) String() string     { _ = "STUB: not implemented"; return "" }
func (k *shortKEMPrivKey) Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }
func (k *shortKEMPrivKey) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *shortKEMPrivKey) Equal(pk kem.PrivateKey) bool { _ = "STUB: not implemented"; return false }

func (k *shortKEMPrivKey) Public() kem.PublicKey {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey)
}
