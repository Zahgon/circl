// Implements the scheme of https://eprint.iacr.org/2019/966

package tkn

import (
	"io"

	pairing "github.com/cloudflare/circl/ecc/bls12381"
)

type PublicParams struct {
	b2  *matrixG2
	wb1 *matrixG1
	btk *matrixGT
}

func (p *PublicParams) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *PublicParams) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (p *PublicParams) Equal(p2 *PublicParams) bool { _ = "STUB: not implemented"; return false }

type SecretParams struct {
	a       *matrixZp
	wtA     *matrixZp
	bstar   *matrixZp
	bstar12 *matrixZp
	k       *matrixZp // vectors are represented as 1 x n matrices
	prfKey  []byte
}

func (s *SecretParams) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SecretParams) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SecretParams) Equal(s2 *SecretParams) bool { _ = "STUB: not implemented"; return false }

type AttributesKey struct {
	a      *Attributes
	k1     *matrixG2
	k2     *matrixG1
	k3     map[string]*matrixG1
	k3wild map[string]*matrixG1 // only contains wildcards
}

func (a *AttributesKey) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *AttributesKey) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (a *AttributesKey) Equal(b *AttributesKey) bool { _ = "STUB: not implemented"; return false }

type ciphertextHeader struct {
	p     *Policy
	c1    *matrixG2
	c2    []*matrixG2
	c3    []*matrixG1
	c3neg []*matrixG1 // additional vector for negated attributes
}

func (hdr *ciphertextHeader) marshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Now we need to indicate how long c2, c3, c3neg are.
// Each array will be the same size (or nil), so with more work we can specialize
// but for now we will ignore that.

func (hdr *ciphertextHeader) unmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func GenerateParams(rand io.Reader) (*PublicParams, *SecretParams, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func max(in []int) int { _ = "STUB: not implemented"; return 0 }

// encapsulate creates a new ephemeral key and header that can be opened to it. This is
// the transformation of an Elgamal like scheme to a KEM.
func encapsulate(rand io.Reader, pp *PublicParams, policy *Policy) (*ciphertextHeader, *pairing.Gt, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func deriveAttributeKeys(rand io.Reader, sp *SecretParams, attrs *Attributes) (*AttributesKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For wild k3 is y term, k3wild is constant term

// Decapsulate decapsulates
func decapsulate(header *ciphertextHeader, key *AttributesKey) (*pairing.Gt, error) {
	_ = "STUB: not implemented"
	// First we need to determine the satisfying assignment: which attributes in attr
	// are needed.
	return nil, nil
}

// We use pi to determine which D to sum into

// p1, p2 are the left halves of the pairings.
