// Package hpke implements the Hybrid Public Key Encryption (HPKE) standard
// specified by draft-irtf-cfrg-hpke-07.
//
// HPKE works for any combination of a public-key encapsulation mechanism
// (KEM), a key derivation function (KDF), and an authenticated encryption
// scheme with additional data (AEAD).
//
// Specification in
// https://datatracker.ietf.org/doc/draft-irtf-cfrg-hpke
//
// BUG(cjpatton): This package does not implement the "Export-Only" mode of the
// HPKE context. In particular, it does not recognize the AEAD codepoint
// reserved for this purpose (0xFFFF).
package hpke

import (
	"encoding"
	"errors"
	"io"

	"github.com/cloudflare/circl/kem"
)

const versionLabel = "HPKE-v1"

// Context defines the capabilities of an HPKE context.
type Context interface {
	encoding.BinaryMarshaler
	// Export takes a context string exporterContext and a desired length (in
	// bytes), and produces a secret derived from the internal exporter secret
	// using the corresponding KDF Expand function. It panics if length is
	// greater than 255*N bytes, where N is the size (in bytes) of the KDF's
	// output.
	Export(exporterContext []byte, length uint) []byte
	// Suite returns the cipher suite corresponding to this context.
	Suite() Suite
}

// Sealer encrypts a plaintext using an AEAD encryption.
type Sealer interface {
	Context
	// Seal takes a plaintext and associated data to produce a ciphertext.
	// The nonce is handled by the Sealer and incremented after each call.
	Seal(pt, aad []byte) (ct []byte, err error)
}

// Opener decrypts a ciphertext using an AEAD encryption.
type Opener interface {
	Context
	// Open takes a ciphertext and associated data to recover, if successful,
	// the plaintext. The nonce is handled by the Opener and incremented after
	// each call.
	Open(ct, aad []byte) (pt []byte, err error)
}

// modeID represents an HPKE variant.
type modeID = uint8

const (
	// modeBase to enable encryption to the holder of a given KEM private key.
	modeBase modeID = 0x00
	// modePSK extends the base mode by allowing the Receiver to authenticate
	// that the sender possessed a given pre-shared key (PSK).
	modePSK modeID = 0x01
	// modeAuth extends the base mode by allowing the Receiver to authenticate
	// that the sender possessed a given KEM private key.
	modeAuth modeID = 0x02
	// modeAuthPSK provides a combination of the PSK and Auth modes.
	modeAuthPSK modeID = 0x03
)

// Suite is an HPKE cipher suite consisting of a KEM, KDF, and AEAD algorithm.
type Suite struct {
	kemID  KEM
	kdfID  KDF
	aeadID AEAD
}

// NewSuite builds a Suite from a specified set of algorithms. Panics
// if an algorithm identifier is not valid.
func NewSuite(kemID KEM, kdfID KDF, aeadID AEAD) Suite {
	_ = "STUB: not implemented"
	return *new(Suite)
}

type state struct {
	Suite
	modeID modeID
	skS    kem.PrivateKey
	pkS    kem.PublicKey
	psk    []byte
	pskID  []byte
	info   []byte
}

// Sender performs hybrid public-key encryption.
type Sender struct {
	state
	pkR kem.PublicKey
}

// NewSender creates a Sender with knowledge of the receiver's public-key.
func (suite Suite) NewSender(pkR kem.PublicKey, info []byte) (*Sender, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setup generates a new HPKE context used for Base Mode encryption.
// Returns the Sealer and corresponding encapsulated key.
func (s *Sender) Setup(rnd io.Reader) (enc []byte, seal Sealer, err error) {
	_ = "STUB: not implemented"
	return nil, *new(Sealer), nil
}

// SetupAuth generates a new HPKE context used for Auth Mode encryption.
// Returns the Sealer and corresponding encapsulated key.
func (s *Sender) SetupAuth(rnd io.Reader, skS kem.PrivateKey) (
	enc []byte, seal Sealer, err error,
) {
	_ = "STUB: not implemented"
	return nil, *new(Sealer), nil
}

// SetupPSK generates a new HPKE context used for PSK Mode encryption.
// Returns the Sealer and corresponding encapsulated key.
func (s *Sender) SetupPSK(rnd io.Reader, psk, pskID []byte) (
	enc []byte, seal Sealer, err error,
) {
	_ = "STUB: not implemented"
	return nil, *new(Sealer), nil
}

// SetupAuthPSK generates a new HPKE context used for Auth-PSK Mode encryption.
// Returns the Sealer and corresponding encapsulated key.
func (s *Sender) SetupAuthPSK(rnd io.Reader, skS kem.PrivateKey, psk, pskID []byte) (
	enc []byte, seal Sealer, err error,
) {
	_ = "STUB: not implemented"
	return nil, *new(Sealer), nil
}

// Receiver performs hybrid public-key decryption.
type Receiver struct {
	state
	skR kem.PrivateKey
	enc []byte
}

// NewReceiver creates a Receiver with knowledge of a private key.
func (suite Suite) NewReceiver(skR kem.PrivateKey, info []byte) (
	*Receiver, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setup generates a new HPKE context used for Base Mode encryption.
// Setup takes an encapsulated key and returns an Opener.
func (r *Receiver) Setup(enc []byte) (Opener, error) {
	_ = "STUB: not implemented"
	return *new(Opener), nil
}

// SetupAuth generates a new HPKE context used for Auth Mode encryption.
// SetupAuth takes an encapsulated key and a public key, and returns an Opener.
func (r *Receiver) SetupAuth(enc []byte, pkS kem.PublicKey) (Opener, error) {
	_ = "STUB: not implemented"
	return *new(Opener), nil
}

// SetupPSK generates a new HPKE context used for PSK Mode encryption.
// SetupPSK takes an encapsulated key, and a pre-shared key; and returns an
// Opener.
func (r *Receiver) SetupPSK(enc, psk, pskID []byte) (Opener, error) {
	_ = "STUB: not implemented"
	return *new(Opener), nil
}

// SetupAuthPSK generates a new HPKE context used for Auth-PSK Mode encryption.
// SetupAuthPSK takes an encapsulated key, a public key, and a pre-shared key;
// and returns an Opener.
func (r *Receiver) SetupAuthPSK(
	enc, psk, pskID []byte, pkS kem.PublicKey,
) (Opener, error) {
	_ = "STUB: not implemented"
	return *new(Opener), nil
}

func (s *Sender) allSetup(rnd io.Reader) ([]byte, Sealer, error) {
	_ = "STUB: not implemented"
	return nil, *new(Sealer), nil
}

func (r *Receiver) allSetup() (Opener, error) { _ = "STUB: not implemented"; return *new(Opener), nil }

var (
	ErrInvalidHPKESuite       = errors.New("hpke: invalid HPKE suite")
	ErrInvalidKDF             = errors.New("hpke: invalid KDF identifier")
	ErrInvalidKEM             = errors.New("hpke: invalid KEM identifier")
	ErrInvalidAuthKEM         = errors.New("hpke: KEM does not support Auth mode")
	ErrInvalidAEAD            = errors.New("hpke: invalid AEAD identifier")
	ErrInvalidKEMPublicKey    = errors.New("hpke: invalid KEM public key")
	ErrInvalidKEMPrivateKey   = errors.New("hpke: invalid KEM private key")
	ErrInvalidKEMSharedSecret = errors.New("hpke: invalid KEM shared secret")
	ErrInvalidKEMDeriveKey    = errors.New("hpke: too many tries to derive KEM key")
	ErrAEADSeqOverflows       = errors.New("hpke: AEAD sequence number overflows")
)
