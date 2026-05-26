package hpke

import (
	"crypto/cipher"
)

type encdecContext struct {
	// Serialized parameters
	suite              Suite
	sharedSecret       []byte
	secret             []byte
	keyScheduleContext []byte
	exporterSecret     []byte
	key                []byte
	baseNonce          []byte
	sequenceNumber     []byte

	// Operational parameters
	cipher.AEAD
	nonce []byte
}

type (
	sealContext struct{ *encdecContext }
	openContext struct{ *encdecContext }
)

// Export takes a context string exporterContext and a desired length (in
// bytes), and produces a secret derived from the internal exporter secret
// using the corresponding KDF Expand function. It panics if length is
// greater than 255*N bytes, where N is the size (in bytes) of the KDF's
// output.
func (c *encdecContext) Export(exporterContext []byte, length uint) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (c *encdecContext) Suite() Suite { _ = "STUB: not implemented"; return *new(Suite) }

func (c *encdecContext) calcNonce() []byte { _ = "STUB: not implemented"; return nil }

func (c *encdecContext) increment() error {
	_ = "STUB: not implemented"
	// tests whether the sequence number is all-ones, which prevents an
	// overflow after the increment.
	return nil
}

// performs an increment by 1 and verifies whether the sequence overflows.

func (c *sealContext) Seal(pt, aad []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *openContext) Open(ct, aad []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
