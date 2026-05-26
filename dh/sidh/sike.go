package sidh

import (
	"io"

	"github.com/cloudflare/circl/dh/sidh/internal/common"
	"github.com/cloudflare/circl/internal/sha3"
)

// SIKE KEM interface.
//
// Deprecated: not cryptographically secure.
type KEM struct {
	allocated   bool
	rng         io.Reader
	msg         []byte
	secretBytes []byte
	params      *common.SidhParams
	shake       sha3.State
}

// NewSike434 instantiates SIKE/p434 KEM.
//
// Deprecated: not cryptographically secure.
func NewSike434(rng io.Reader) *KEM { _ = "STUB: not implemented"; return nil }

// NewSike503 instantiates SIKE/p503 KEM.
//
// Deprecated: not cryptographically secure.
func NewSike503(rng io.Reader) *KEM { _ = "STUB: not implemented"; return nil }

// NewSike751 instantiates SIKE/p751 KEM.
//
// Deprecated: not cryptographically secure.
func NewSike751(rng io.Reader) *KEM { _ = "STUB: not implemented"; return nil }

// Allocate allocates KEM object for multiple SIKE operations. The rng
// must be cryptographically secure PRNG.
func (c *KEM) Allocate(id uint8, rng io.Reader) { _ = "STUB: not implemented"; return }

// Encapsulate receives the public key and generates SIKE ciphertext and shared secret.
// The generated ciphertext is used for authentication.
// Error is returned in case PRNG fails. Function panics in case wrongly formatted
// input was provided.
func (c *KEM) Encapsulate(ciphertext, secret []byte, pub *PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate ephemeral value

// Ensure bitlength is not bigger then to 2^e2-1

// K = H(msg||(c0||c1))

// Decapsulate given the keypair and ciphertext as inputs, Decapsulate outputs a shared
// secret if plaintext verifies correctly, otherwise function outputs random value.
// Decapsulation may panic in case input is wrongly formatted, in particular, size of
// the 'ciphertext' must be exactly equal to c.CiphertextSize().
func (c *KEM) Decapsulate(secret []byte, prv *PrivateKey, pub *PublicKey, ciphertext []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// r' = G(m'||pub)

// Ensure bitlength is not bigger than 2^e2-1

// S is chosen at random when generating a key and unknown to other party. It is
// important that S is unpredictable to the other party.  Without this check, would
// be possible to recover a secret, by providing series of invalid ciphertexts.
//
// See more details in "On the security of supersingular isogeny cryptosystems"
// (S. Galbraith, et al., 2016, ePrint #859).

// Resets internal state of KEM. Function should be used
// after Allocate and between subsequent calls to Encapsulate
// and/or Decapsulate.
func (c *KEM) Reset() { _ = "STUB: not implemented"; return }

// Returns size of resulting ciphertext.
func (c *KEM) CiphertextSize() int { _ = "STUB: not implemented"; return 0 }

// Returns size of resulting shared secret.
func (c *KEM) SharedSecretSize() int { _ = "STUB: not implemented"; return 0 }

// PublicKeySize returns size of the public key in bytes.
func (c *KEM) PublicKeySize() int { _ = "STUB: not implemented"; return 0 }

// Size returns size of the private key in bytes.
func (c *KEM) PrivateKeySize() int { _ = "STUB: not implemented"; return 0 }

func (c *KEM) generateCiphertext(ctext []byte, skA *PrivateKey, pkA, pkB *PublicKey, ptext []byte) {
	_ = "STUB: not implemented"
	return
}

// encrypt uses SIKE public key to encrypt plaintext. Requires cryptographically secure
// PRNG. Returns ciphertext in case encryption succeeds. Returns error in case PRNG fails
// or wrongly formated input was provided.
func (c *KEM) encrypt(ctext []byte, rng io.Reader, pub *PublicKey, ptext []byte) error {
	_ = "STUB: not implemented"
	return nil

	// c1 must be security level + 64 bits (see [SIKE] 1.4 and 4.3.3)
}

// decrypt uses SIKE private key to decrypt ciphertext. Returns plaintext in case
// decryption succeeds or error in case unexpected input was provided.
// Constant time.
func (c *KEM) decrypt(n []byte, prv *PrivateKey, ctext []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ctext is a concatenation of (ciphertext = pubkey_A || c1)
// it must be security level + 64 bits (see [SIKE] 1.4 and 4.3.3)
// Lengths has been already checked by Decapsulate()
