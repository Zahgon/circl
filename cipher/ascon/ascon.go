// Package ascon provides ASCON family of light-weight AEAD ciphers.
//
// This package implements Ascon128 and Ascon128a two AEAD ciphers as specified
// in ASCON v1.2 by C. Dobraunig, M. Eichlseder, F. Mendel, M. Schläffer.
// https://ascon.iaik.tugraz.at/index.html
//
// It also implements Ascon-80pq, which has an increased key-size to provide
// more resistance against a quantum adversary using Grover’s algorithm for
// key search. Since Ascon-128 and Ascon-80pq share the same building blocks
// and same parameters except the size of the key, it is claimed the same
// security for Ascon-80pq against classical attacks as for Ascon-128.
package ascon

import (
	"errors"
)

const (
	KeySize     = 16 // For Ascon128 and Ascon128a.
	KeySize80pq = 20 // Only for Ascon80pq.
	NonceSize   = 16
	TagSize     = 16
)

type Mode int

// KeySize is 16 for Ascon128 and Ascon128a, or 20 for Ascon80pq.
func (m Mode) KeySize() int { _ = "STUB: not implemented"; return 0 }

func (m Mode) String() string { _ = "STUB: not implemented"; return "" }

const (
	Ascon128  Mode = 1
	Ascon128a Mode = 2
	Ascon80pq Mode = -1
)

const permA = 12

type Cipher struct {
	key  [3]uint64
	mode Mode
}

// New returns a Cipher struct implementing the crypto/cipher.AEAD interface.
// The key must be Mode.KeySize() bytes long, and the mode is one of Ascon128,
// Ascon128a or Ascon80pq.
func New(key []byte, m Mode) (*Cipher, error) { _ = "STUB: not implemented"; return nil, nil }

// NonceSize returns the size of the nonce that must be passed to Seal
// and Open.
func (a *Cipher) NonceSize() int {
	_ = "STUB: not implemented"

	// Overhead returns the maximum difference between the lengths of a
	// plaintext and its ciphertext.
	return 0
}

func (a *Cipher) Overhead() int {
	_ = "STUB: not implemented"

	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	//
	// To reuse plaintext's storage for the encrypted output, use plaintext[:0]
	// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
	return 0
}

func (a *Cipher) Seal(dst, nonce, plaintext, additionalData []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Open decrypts and authenticates ciphertext, authenticates the
// additional data and, if successful, appends the resulting plaintext
// to dst, returning the updated slice. The nonce must be NonceSize()
// bytes long and both it and the additional data must match the
// value passed to Seal.
//
// To reuse ciphertext's storage for the decrypted output, use ciphertext[:0]
// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
//
// Even if the function fails, the contents of dst, up to its capacity,
// may be overwritten.
func (a *Cipher) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func abs(x int) int { _ = "STUB: not implemented"; return 0 }

// blockSize = 8 for Ascon128 and Ascon80pq, or 16 for Ascon128a.
func (a *Cipher) blockSize() int { _ = "STUB: not implemented"; return 0 }

// permB = 6 for Ascon128 and Ascon80pq, or 8 for Ascon128a.
func (a *Cipher) permB() int { _ = "STUB: not implemented"; return 0 }

func (a *Cipher) initialize(nonce []byte, s *[5]uint64) { _ = "STUB: not implemented"; return }

func (a *Cipher) assocData(add []byte, s *[5]uint64) { _ = "STUB: not implemented"; return }

func (a *Cipher) procText(in, out []byte, enc bool, s *[5]uint64) {
	_ = "STUB: not implemented"
	return
}

func (a *Cipher) finalize(tag []byte, s *[5]uint64) { _ = "STUB: not implemented"; return }

func perm(n int, s *[5]uint64) { _ = "STUB: not implemented"; return }

// pC -- addition of constants

// pS -- substitution layer
// Figure 6 from Spec [DHVV18,Dae18]
// https://ascon.iaik.tugraz.at/files/asconv12-nist.pdf

// pL -- linear diffusion layer

// sliceForAppend takes a slice and a requested number of bytes. It returns a
// slice with the contents of the given slice followed by that many bytes and a
// second slice that aliases into it and contains only the extra bytes. If the
// original slice has sufficient capacity then no allocation is performed.
func sliceForAppend(in []byte, n int) (head, tail []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	ErrKeySize    = errors.New("ascon: bad key size")
	ErrNonceSize  = errors.New("ascon: bad nonce size")
	ErrDecryption = errors.New("ascon: invalid ciphertext")
	ErrMode       = errors.New("ascon: invalid cipher mode")
)
