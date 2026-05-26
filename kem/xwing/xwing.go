// Package xwing implements the X-Wing PQ/T hybrid KEM
//
//	https://datatracker.ietf.org/doc/draft-connolly-cfrg-xwing-kem
//
// Implements the final version (-05).
package xwing

import (
	"errors"
	"io"

	"github.com/cloudflare/circl/dh/x25519"
	"github.com/cloudflare/circl/kem/mlkem/mlkem768"
)

// An X-Wing private key.
type PrivateKey struct {
	seed [32]byte
	m    mlkem768.PrivateKey
	x    x25519.Key
	xpk  x25519.Key
}

// An X-Wing public key.
type PublicKey struct {
	m mlkem768.PublicKey
	x x25519.Key
}

const (
	// Size of a seed of a keypair
	SeedSize = 32

	// Size of an X-Wing public key
	PublicKeySize = 1216

	// Size of an X-Wing private key
	PrivateKeySize = 32

	// Size of the seed passed to EncapsulateTo
	EncapsulationSeedSize = 64

	// Size of the established shared key
	SharedKeySize = 32

	// Size of an X-Wing ciphertext.
	CiphertextSize = 1120
)

func combiner(
	out []byte,
	ssm *[mlkem768.SharedKeySize]byte,
	ssx *x25519.Key,
	ctx *x25519.Key,
	pkx *x25519.Key,
) {
	_ = "STUB: not implemented"
	return
}

//   \./
//   /^\

// Packs sk to buf.
//
// Panics if buf is not of size PrivateKeySize
func (sk *PrivateKey) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// Packs pk to buf.
//
// Panics if buf is not of size PublicKeySize.
func (pk *PublicKey) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// DeriveKeyPair derives a public/private keypair deterministically
// from the given seed.
//
// Panics if seed is not of length SeedSize.
func DeriveKeyPair(seed []byte) (*PrivateKey, *PublicKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deriveKeyPair(seed []byte, sk *PrivateKey, pk *PublicKey) { _ = "STUB: not implemented"; return }

// DeriveKeyPairPacked derives a keypair like DeriveKeyPair, and
// returns them packed.
func DeriveKeyPairPacked(seed []byte) ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// GenerateKeyPair generates public and private keys using entropy from rand.
// If rand is nil, crypto/rand.Reader will be used.
func GenerateKeyPair(rand io.Reader) (*PrivateKey, *PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GenerateKeyPairPacked generates a keypair like GenerateKeyPair, and
// returns them packed.
func GenerateKeyPairPacked(rand io.Reader) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Encapsulate generates a shared key and ciphertext that contains it
// for the public key pk using randomness from seed.
//
// seed may be nil, in which case crypto/rand.Reader is used.
//
// Warning: note that the order of the returned ss and ct matches the
// X-Wing standard, which is the reverse of the Circl KEM API.
//
// Returns ErrPubKey if ML-KEM encapsulation key check fails.
//
// Panics if pk is not of size PublicKeySize, or randomness could not
// be read from crypto/rand.Reader.
func Encapsulate(pk, seed []byte) (ss, ct []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Decapsulate computes the shared key which is encapsulated in ct
// for the private key sk.
//
// Panics if sk or ct are not of length PrivateKeySize and CiphertextSize
// respectively.
func Decapsulate(ct, sk []byte) (ss []byte) { _ = "STUB: not implemented"; return nil }

// Raised when passing a byte slice of the wrong size for the shared
// secret to the EncapsulateTo or DecapsulateTo functions.
var ErrSharedKeySize = errors.New("wrong size for shared key")

// EncapsulateTo generates a shared key and ciphertext that contains it
// for the public key using randomness from seed and writes the shared key
// to ss and ciphertext to ct.
//
// Panics if ss, ct or seed are not of length SharedKeySize, CiphertextSize
// and EncapsulationSeedSize respectively.
//
// seed may be nil, in which case crypto/rand.Reader is used to generate one.
func (pk *PublicKey) EncapsulateTo(ct, ss, seed []byte) { _ = "STUB: not implemented"; return }

// A peer public key with low order points results in an all-zeroes
// shared secret. Ignored for now pending clarification in the spec,
// https://github.com/dconnolly/draft-connolly-cfrg-xwing-kem/issues/28

// DecapsulateTo computes the shared key which is encapsulated in ct
// for the private key.
//
// Panics if ct or ss are not of length CiphertextSize and SharedKeySize
// respectively.
func (sk *PrivateKey) DecapsulateTo(ss, ct []byte) { _ = "STUB: not implemented"; return }

// A peer public key with low order points results in an all-zeroes
// shared secret. Ignored for now pending clarification in the spec,
// https://github.com/dconnolly/draft-connolly-cfrg-xwing-kem/issues/28

// Unpacks pk from buf.
//
// Panics if buf is not of size PublicKeySize.
//
// Returns ErrPubKey if pk fails the ML-KEM encapsulation key check.
func (pk *PublicKey) Unpack(buf []byte) error { _ = "STUB: not implemented"; return nil }

// Unpacks sk from buf.
//
// Panics if buf is not of size PrivateKeySize.
func (sk *PrivateKey) Unpack(buf []byte) { _ = "STUB: not implemented"; return }
