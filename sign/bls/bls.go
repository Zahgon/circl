// Package bls provides BLS signatures using the BLS12-381 pairing curve.
//
// This packages implements the IETF/CFRG draft for BLS signatures [1].
// Currently only the BASIC mode (one of the three modes specified
// in the draft) is supported. The pairing function is instantiated
// with the BLS12-381 curve.
//
// # Groups
//
// The BLS signature scheme can be instantiated with keys in one of the
// two groups: G1 or G2, which correspond to the input domain of a pairing
// function e(G1,G2) -> Gt.
// Thus, choosing keys in G1 implies that signature values are internally
// represented in G2; or viceversa. Use the types KeyG1SigG2 or KeyG2SigG1
// to express this preference.
//
// # Serialization
//
// The serialization of elements in G1 and G2 follows the recommendation
// given in [2], in order to be compatible with other implementations of
// BLS12-381 curve.
//
// # References
//
// [1] https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-bls-signature-05
//
// [2] https://github.com/zkcrypto/bls12_381/blob/0.7.0/src/notes/serialization.rs
package bls

import (
	"crypto"
	"errors"

	GG "github.com/cloudflare/circl/ecc/bls12381"
)

var (
	ErrInvalid    = errors.New("bls: invalid BLS instance")
	ErrInvalidKey = errors.New("bls: invalid key")
	ErrInvalidSig = errors.New("bls: invalid signature")
	ErrKeyGen     = errors.New("bls: too many unsuccessful key generation tries")
	ErrShortIKM   = errors.New("bls: IKM material shorter than 32 bytes")
	ErrAggregate  = errors.New("bls: error while aggregating signatures")
)

const (
	dstG1 = "BLS_SIG_BLS12381G1_XMD:SHA-256_SSWU_RO_NUL_"
	dstG2 = "BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_"
)

type Signature = []byte

type (
	// G1 group used for keys defined in pairing group G1.
	G1 struct{ g GG.G1 }
	// G2 group used for keys defined in pairing group G2.
	G2 struct{ g GG.G2 }
	// KeyG1SigG2 sets the keys to G1 and signatures to G2.
	KeyG1SigG2 = G1
	// KeyG2SigG1 sets the keys to G2 and signatures to G1.
	KeyG2SigG1 = G2
)

func (f *G1) setBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

func (f *G2) setBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

func (f *G1) hash(msg []byte) { _ = "STUB: not implemented"; return }
func (f *G2) hash(msg []byte) { _ = "STUB: not implemented"; return }

// KeyGroup determines the group used for keys, while the other
// group is used for signatures.
type KeyGroup interface{ G1 | G2 }

type PrivateKey[K KeyGroup] struct {
	key GG.Scalar
	pub *PublicKey[K]
}

type PublicKey[K KeyGroup] struct{ key K }

func (k *PrivateKey[K]) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *

	// PublicKey computes the corresponding public key. The key is cached
	// for further invocations to this function.
	new(crypto.PublicKey)
}

func (k *PrivateKey[K]) PublicKey() *PublicKey[K] { _ = "STUB: not implemented"; return nil }

func (k *PrivateKey[K]) Equal(x crypto.PrivateKey) bool { _ = "STUB: not implemented"; return false }

// Validate explicitly determines if a private key is valid.
func (k *PrivateKey[K]) Validate() bool { _ = "STUB: not implemented"; return false }

// MarshalBinary returns a slice with the representation of
// the underlying PrivateKey scalar (in big-endian order).
func (k *PrivateKey[K]) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *PrivateKey[K]) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// Validate explicitly determines if a public key is valid.
func (k *PublicKey[K]) Validate() bool { _ = "STUB: not implemented"; return false }

func (k *PublicKey[K]) Equal(x crypto.PublicKey) bool { _ = "STUB: not implemented"; return false }

// MarshalBinary returns a slice with the compressed
// representation of the underlying element in G1 or G2.
func (k *PublicKey[K]) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *PublicKey[K]) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// KeyGen derives a private key for the specified group (G1 or G2).
// The length of ikm material should be at least 32 bytes length.
// The salt value should be either empty or a uniformly random
// bytes whose length equals the output length of SHA-256.
func KeyGen[K KeyGroup](ikm, salt, keyInfo []byte) (*PrivateKey[K], error) {
	_ = "STUB: not implemented"
	// Implements recommended method at:
	// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-bls-signature-05#name-keygen
	return nil, nil
}

// Sign computes a signature of a message using a key (defined in
// G1 or G1).
func Sign[K KeyGroup](k *PrivateKey[K], msg []byte) Signature {
	_ = "STUB: not implemented"
	return *new(Signature)
}

// Verify returns true if the signature of a message is valid for the
// corresponding public key.
func Verify[K KeyGroup](pub *PublicKey[K], msg []byte, sig Signature) bool {
	_ = "STUB: not implemented"
	return false
}

// Aggregate produces a unified signature given a list of signatures.
// To specify the group of keys pass either G1{} or G2{} as the first
// parameter.
func Aggregate[K KeyGroup](k K, sigs []Signature) (Signature, error) {
	_ = "STUB: not implemented"
	return *new(Signature), nil
}

// VerifyAggregate returns true if the aggregated signature is valid for
// the list of messages and public keys provided. The slices must have
// equal size and have at least one element.
func VerifyAggregate[K KeyGroup](pubs []*PublicKey[K], msgs [][]byte, aggSig Signature) bool {
	_ = "STUB: not implemented"
	return false
}

// 1. If any two input messages are equal, return INVALID.

// 2. CoreAggregateVerify algorithm checks an aggregated signature over
// several (PK, message) pairs.
