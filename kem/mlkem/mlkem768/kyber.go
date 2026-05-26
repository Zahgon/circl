// Code generated from pkg.templ.go. DO NOT EDIT.

// Package mlkem768 implements the IND-CCA2 secure key encapsulation mechanism
// ML-KEM-768 as defined in FIPS203.
package mlkem768

import (
	"io"

	"github.com/cloudflare/circl/kem"
	cpapke "github.com/cloudflare/circl/pke/kyber/kyber768"
)

const (
	// Size of seed for NewKeyFromSeed
	KeySeedSize = cpapke.KeySeedSize + 32

	// Size of seed for EncapsulateTo.
	EncapsulationSeedSize = 32

	// Size of the established shared key.
	SharedKeySize = 32

	// Size of the encapsulated shared key.
	CiphertextSize = cpapke.CiphertextSize

	// Size of a packed public key.
	PublicKeySize = cpapke.PublicKeySize

	// Size of a packed private key.
	PrivateKeySize = cpapke.PrivateKeySize + cpapke.PublicKeySize + 64
)

// Type of a ML-KEM-768 public key
type PublicKey struct {
	pk *cpapke.PublicKey

	hpk [32]byte // H(pk)
}

// Type of a ML-KEM-768 private key
type PrivateKey struct {
	sk  *cpapke.PrivateKey
	pk  *cpapke.PublicKey
	hpk [32]byte // H(pk)
	z   [32]byte
}

// NewKeyFromSeed derives a public/private keypair deterministically
// from the given seed.
//
// Panics if seed is not of length KeySeedSize.
func NewKeyFromSeed(seed []byte) (*PublicKey, *PrivateKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Compute H(pk)

// GenerateKeyPair generates public and private keys using entropy from rand.
// If rand is nil, crypto/rand.Reader will be used.
func GenerateKeyPair(rand io.Reader) (*PublicKey, *PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EncapsulateTo generates a shared key and ciphertext that contains it
// for the public key using randomness from seed and writes the shared key
// to ss and ciphertext to ct.
//
// Panics if ss, ct or seed are not of length SharedKeySize, CiphertextSize
// and EncapsulationSeedSize respectively.
//
// seed may be nil, in which case crypto/rand.Reader is used to generate one.
func (pk *PublicKey) EncapsulateTo(ct, ss []byte, seed []byte) { _ = "STUB: not implemented"; return }

// (K', r) = G(m ‖ H(pk))

// c = Kyber.CPAPKE.Enc(pk, m, r)

// DecapsulateTo computes the shared key which is encapsulated in ct
// for the private key.
//
// Panics if ct or ss are not of length CiphertextSize and SharedKeySize
// respectively.
func (sk *PrivateKey) DecapsulateTo(ss, ct []byte) { _ = "STUB: not implemented"; return }

// m' = Kyber.CPAPKE.Dec(sk, ct)

// (K'', r') = G(m' ‖ H(pk))

// c' = Kyber.CPAPKE.Enc(pk, m', r')

// Compute shared secret in case of rejection: ss₂ = PRF(z ‖ c)

// Set ss2 to the real shared secret if c = c'.

// Packs sk to buf.
//
// Panics if buf is not of size PrivateKeySize.
func (sk *PrivateKey) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// Unpacks sk from buf.
//
// Panics if buf is not of size PrivateKeySize.
//
// Returns an error if buf is not of size PrivateKeySize, or private key
// doesn't pass the ML-KEM decapsulation key check.
func (sk *PrivateKey) Unpack(buf []byte) error { _ = "STUB: not implemented"; return nil }

// Packs pk to buf.
//
// Panics if buf is not of size PublicKeySize.
func (pk *PublicKey) Pack(buf []byte) { _ = "STUB: not implemented"; return }

// Unpacks pk from buf.
//
// Returns an error if buf is not of size PublicKeySize, or the public key
// is not normalized.
func (pk *PublicKey) Unpack(buf []byte) error { _ = "STUB: not implemented"; return nil }

// Compute cached H(pk)

// Boilerplate down below for the KEM scheme API.

type scheme struct{}

var sch kem.Scheme = &scheme{}

// Scheme returns a KEM interface.
func Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }

func (*scheme) Name() string               { _ = "STUB: not implemented"; return "" }
func (*scheme) PublicKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (*scheme) PrivateKeySize() int        { _ = "STUB: not implemented"; return 0 }
func (*scheme) SeedSize() int              { _ = "STUB: not implemented"; return 0 }
func (*scheme) SharedKeySize() int         { _ = "STUB: not implemented"; return 0 }
func (*scheme) CiphertextSize() int        { _ = "STUB: not implemented"; return 0 }
func (*scheme) EncapsulationSeedSize() int { _ = "STUB: not implemented"; return 0 }

func (sk *PrivateKey) Scheme() kem.Scheme { _ = "STUB: not implemented"; return *new(kem.Scheme) }
func (pk *PublicKey) Scheme() kem.Scheme  { _ = "STUB: not implemented"; return *new(kem.Scheme) }

func (sk *PrivateKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (sk *PrivateKey) Equal(other kem.PrivateKey) bool { _ = "STUB: not implemented"; return false }

func (pk *PublicKey) Equal(other kem.PublicKey) bool { _ = "STUB: not implemented"; return false }

func (sk *PrivateKey) Public() kem.PublicKey { _ = "STUB: not implemented"; return *new(kem.PublicKey) }

func (pk *PublicKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (*scheme) GenerateKeyPair() (kem.PublicKey, kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey), nil
}

func (*scheme) DeriveKeyPair(seed []byte) (kem.PublicKey, kem.PrivateKey) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), *new(kem.PrivateKey)
}

func (*scheme) Encapsulate(pk kem.PublicKey) (ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (*scheme) EncapsulateDeterministically(pk kem.PublicKey, seed []byte) (
	ct, ss []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (*scheme) Decapsulate(sk kem.PrivateKey, ct []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*scheme) UnmarshalBinaryPublicKey(buf []byte) (kem.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PublicKey), nil
}

func (*scheme) UnmarshalBinaryPrivateKey(buf []byte) (kem.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(kem.PrivateKey), nil
}
