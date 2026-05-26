package hpke

// This file implements a hybrid KEM for HPKE using a simple concatenation
// combiner.
//
// WARNING It is not safe to combine arbitrary KEMs using this combiner.
// See the draft specification for more details:
// https://bwesterb.github.io/draft-westerbaan-cfrg-hpke-xyber768d00/draft-westerbaan-cfrg-hpke-xyber768d00.html#name-security-considerations

import (
	"github.com/cloudflare/circl/kem"
)

type hybridKEM struct {
	kemBase
	kemA kem.Scheme
	kemB kem.Scheme
}

func (h hybridKEM) PrivateKeySize() int        { _ = "STUB: not implemented"; return 0 }
func (h hybridKEM) SeedSize() int              { _ = "STUB: not implemented"; return 0 }
func (h hybridKEM) CiphertextSize() int        { _ = "STUB: not implemented"; return 0 }
func (h hybridKEM) PublicKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (h hybridKEM) EncapsulationSeedSize() int { _ = "STUB: not implemented"; return 0 }

func (h hybridKEM) SharedKeySize() int { _ = "STUB: not implemented"; return 0 }
func (h hybridKEM) Name() string       { _ = "STUB: not implemented"; return "" }

func (h hybridKEM) AuthDecapsulate(skR kem.PrivateKey,
	ct []byte,
	pkS kem.PublicKey,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h hybridKEM) AuthEncapsulate(pkr kem.PublicKey, sks kem.PrivateKey) (
	ct []byte, ss []byte, err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (h hybridKEM) AuthEncapsulateDeterministically(pkr kem.PublicKey, sks kem.PrivateKey, seed []byte) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (h hybridKEM) Encapsulate(pkr kem.PublicKey) (
	ct []byte, ss []byte, err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (h hybridKEM) Decapsulate(skr kem.PrivateKey, ct []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h hybridKEM) EncapsulateDeterministically(
	pkr kem.PublicKey, seed []byte,
) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type hybridKEMPrivKey struct {
	scheme kem.Scheme
	privA  kem.PrivateKey
	privB  kem.PrivateKey
}

func (k *hybridKEMPrivKey) Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }

func (k *hybridKEMPrivKey) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *hybridKEMPrivKey) Equal(sk kem.PrivateKey) bool { _ = "STUB: not implemented"; return false }

func (k *hybridKEMPrivKey) Public() kem.PublicKey {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey)
}

type hybridKEMPubKey struct {
	scheme kem.Scheme
	pubA   kem.PublicKey
	pubB   kem.PublicKey
}

func (k *hybridKEMPubKey) Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }

func (k hybridKEMPubKey) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *hybridKEMPubKey) Equal(pk kem.PublicKey) bool { _ = "STUB: not implemented"; return false }

// Deterministically derives a keypair from a seed. If you're unsure,
// you're better off using GenerateKey().
//
// Panics if seed is not of length SeedSize().
func (h hybridKEM) DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey) {
	_ = "STUB: not implemented"
	// Implementation based on
	// https://www.ietf.org/archive/id/draft-irtf-cfrg-hpke-07.html#name-derivekeypair
	return *new(kem.PublicKey), *new(kem.PrivateKey)
}

func (h hybridKEM) GenerateKeyPair() (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (h hybridKEM) UnmarshalBinaryPrivateKey(data []byte) (kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PrivateKey), nil
}

func (h hybridKEM) UnmarshalBinaryPublicKey(data []byte) (kem.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), nil
}
