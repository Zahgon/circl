package slhdsa

import (
	"crypto"
	"io"

	"github.com/cloudflare/circl/xof"
	_ "golang.org/x/crypto/sha3"
)

// [PreHash] is a helper for hashing a message before signing.
// It implements the [io.Writer] interface, so the message can be provided
// in chunks before calling the [SignDeterministic], [SignRandomized], or
// [Verify] functions.
// Pre-hash must not be used for generating pure signatures.
type PreHash struct {
	writer interface {
		io.Writer
		Reset()
	}
	size int
	oid  byte
}

// [NewPreHashWithHash] is used to prehash messages using either the SHA2 or
// SHA3 hash functions.
// Returns [ErrPreHash] if the function is not supported.
func NewPreHashWithHash(h crypto.Hash) (*PreHash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// [NewPreHashWithXof] is used to prehash messages using either SHAKE-128
// or SHAKE-256.
// Returns [ErrPreHash] if the function is not supported.
func NewPreHashWithXof(x xof.ID) (*PreHash, error) { _ = "STUB: not implemented"; return nil, nil }

func (ph *PreHash) Reset() { _ = "STUB: not implemented"; return }
func (ph *PreHash) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// BuildMessage returns a [Message] for signing, and resets the writer.
		nil
}

func (ph *PreHash) BuildMessage() (*Message, error) {
	_ = "STUB: not implemented"
	// Source https://csrc.nist.gov/Projects/computer-security-objects-register/algorithm-registration
	return nil, nil
}

// [Message] wraps a message for signing.
type Message struct {
	msg       []byte
	isPreHash byte
}

// For pure signatures, use [NewMessage] to pass the message to be signed.
// For pre-hashed signatures, use [PreHash] to hash the message first, and
// then use [PreHash.BuildMessage] to get a [Message] to be signed.
func NewMessage(msg []byte) *Message { _ = "STUB: not implemented"; return nil }

func (m *Message) getMsgPrime(context []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// See FIPS 205 -- Section 10.2 -- Algorithm 23 and Algorithm 25.
	return nil, nil
}
