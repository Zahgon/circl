// Package dleq provides zero-knowledge proofs of Discrete-Logarithm Equivalence (DLEQ).
//
// This implementation is compatible with the one used for VOPRFs [1].
// It supports batching proofs to amortize the cost of the proof generation and
// verification.
//
// References:
//
//	[1] RFC-9497: https://www.rfc-editor.org/info/rfc9497
package dleq

import (
	"crypto"
	"io"

	"github.com/cloudflare/circl/group"
)

const (
	labelSeed         = "Seed-"
	labelChallenge    = "Challenge"
	labelComposite    = "Composite"
	labelHashToScalar = "HashToScalar-"
)

type Params struct {
	G   group.Group
	H   crypto.Hash
	DST []byte
}

type Proof struct {
	c, s group.Scalar
}

type Prover struct{ Params }

func (p Prover) Prove(k group.Scalar, a, ka, b, kb group.Element, rnd io.Reader) (*Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Prover) ProveWithRandomness(k group.Scalar, a, ka, b, kb group.Element, rnd group.Scalar) (*Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Prover) ProveBatch(k group.Scalar, a, ka group.Element, bi, kbi []group.Element, rnd io.Reader) (*Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Prover) ProveBatchWithRandomness(
	k group.Scalar,
	a, ka group.Element,
	bi, kbi []group.Element,
	rnd group.Scalar,
) (*Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Params) computeComposites(
	k group.Scalar,
	ka group.Element,
	bi []group.Element,
	kbi []group.Element,
) (m, z group.Element, err error) {
	_ = "STUB: not implemented"
	return *new(group.Element), *new(group.Element), nil
}

func (p Params) doChallenge(a [5][]byte) group.Scalar {
	_ = "STUB: not implemented"
	return *new(group.Scalar)
}

type Verifier struct{ Params }

func (v Verifier) Verify(a, ka, b, kb group.Element, p *Proof) bool {
	_ = "STUB: not implemented"
	return false
}

func (v Verifier) VerifyBatch(a, ka group.Element, bi, kbi []group.Element, p *Proof) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Proof) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Proof) UnmarshalBinary(g group.Group, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func mustWrite(h io.Writer, bytes []byte) { _ = "STUB: not implemented"; return }
