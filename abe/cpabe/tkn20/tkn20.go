//go:generate go run gen_testdata.go

// Package tkn20 implements a ciphertext-policy ABE by Tomida, Kawahara, Nishimaki.
//
// This is an implementation of an IND-CCA2 secure variant of the Ciphertext-Policy
// Attribute Based Encryption (CP-ABE) scheme by
// J. Tomida, Y. Kawahara, and R. Nishimaki. Fast, compact, and expressive
// attribute-based encryption. In A. Kiayias, M. Kohlweiss, P. Wallden, and
// V. Zikas, editors, PKC, volume 12110 of Lecture Notes in Computer Science,
// pages 3–33. Springer, 2020. https://eprint.iacr.org/2019/966
//
// # Update v1.3.8
//
// As of v1.3.8, ciphertext format changed to use wider prefixes.
// Ciphertexts in the previous format are still decryptable.
// The following functions are backwards-compatible:
//   - [AttributeKey.Decrypt]
//   - [Attributes.CouldDecrypt]
//   - [Policy.ExtractFromCiphertext]
package tkn20

import (
	"io"

	"github.com/cloudflare/circl/abe/cpabe/tkn20/internal/tkn"
)

type PublicKey struct {
	pp tkn.PublicParams
}

func (p *PublicKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *PublicKey) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (p *PublicKey) Equal(p2 *PublicKey) bool { _ = "STUB: not implemented"; return false }

func (p *PublicKey) Encrypt(rand io.Reader, policy Policy, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SystemSecretKey struct {
	sp tkn.SecretParams
}

func (msk *SystemSecretKey) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (msk *SystemSecretKey) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (msk *SystemSecretKey) Equal(msk2 *SystemSecretKey) bool {
	_ = "STUB: not implemented"
	return false
}

func (msk *SystemSecretKey) KeyGen(rand io.Reader, attrs Attributes) (AttributeKey, error) {
	_ = "STUB: not implemented"
	return *new(AttributeKey), nil
}

type AttributeKey struct {
	ak tkn.AttributesKey
}

func (s *AttributeKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *AttributeKey) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s *AttributeKey) Equal(s2 *AttributeKey) bool { _ = "STUB: not implemented"; return false }

func (s *AttributeKey) Decrypt(ct []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Policy struct {
	policy tkn.Policy
}

func (p *Policy) FromString(str string) error { _ = "STUB: not implemented"; return nil }

func (p *Policy) String() string { _ = "STUB: not implemented"; return "" }

func (p *Policy) ExtractFromCiphertext(ct []byte) error { _ = "STUB: not implemented"; return nil }

func (p *Policy) ExtractAttributeValuePairs() map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *Policy) Equal(p2 *Policy) bool { _ = "STUB: not implemented"; return false }

func (p *Policy) Satisfaction(a Attributes) bool { _ = "STUB: not implemented"; return false }

type Attributes struct {
	attrs tkn.Attributes
}

func (a *Attributes) Equal(a2 *Attributes) bool { _ = "STUB: not implemented"; return false }

func (a *Attributes) CouldDecrypt(ciphertext []byte) bool { _ = "STUB: not implemented"; return false }

func (a *Attributes) FromMap(in map[string]string) { _ = "STUB: not implemented"; return }

func Setup(rand io.Reader) (PublicKey, SystemSecretKey, error) {
	_ = "STUB: not implemented"
	return *new(PublicKey), *new(SystemSecretKey), nil
}
