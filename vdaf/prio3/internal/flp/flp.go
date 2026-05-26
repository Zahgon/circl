// Package flp provides a Fully-Linear Proof (FLP) system.
package flp

import (
	"errors"

	"github.com/cloudflare/circl/vdaf/prio3/arith"
)

// FLP is an instance of a FLP by Boneh et al. Crypto, 2019 paper
// "Zero-Knowledge Proofs on Secret-Shared Data via Fully Linear PCPs",
// https://ia.cr/2019/188
// plus some changes described in VDAF specification.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.1
type FLP[
	G Gadget[P, V, E, F],
	P arith.Poly[P, E], V arith.Vec[V, E], E arith.Elt, F arith.Fp[E],
] struct {
	// Eval evaluates the arithmetic circuit on a measurement and joint randomness.
	Eval func(out V, g Gadget[P, V, E, F], numCalls uint, meas, jointRand V, numShares uint8)
	Valid[G, P, V, E, F]
}

// Prove returns a proof attesting validity to the given measurement.
// Prove randomness must be provided.
// Some statements may require joint randomness too.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.3.3
func (f *FLP[G, P, V, E, F]) Prove(meas, proveRand, jointRand V) V {
	_ = "STUB: not implemented"
	return *new(V)
}

// invN is the inverse of N = g.p.
// Also, since g.p is always a power of two, we call a faster inversion
// method that receives the log2 of g.p.

// Query is the linear Query algorithm run by each verifier on a share of the
// measurement and proof.
// Query randomness must be provided.
// Some statements may require joint randomness too.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.3.4
func (f *FLP[G, P, V, E, F]) Query(
	measShare, proofShare, queryRand, jointRand V, numShares uint8,
) (verifierMsg V, err error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

// Check that t^p != 1. Since p=2^log2p, this requires log2P squares.

// invN is the inverse of N = g.p.
// Also, since g.p is always a power of two, we call a faster inversion
// method that receives the log2 of g.p.

// Extracts the constant factor (1/N) to be multiplied after
// polynomial interpolation and evaluation.

// Decide returns true if the measurement from which it was generated is valid.
//
// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-vdaf-13#section-7.3.5
func (f *FLP[G, P, V, E, F]) Decide(verifierMsg V) bool { _ = "STUB: not implemented"; return false }

var (
	ErrOutputLen        = errors.New("wrong output length")
	ErrMeasurementLen   = errors.New("invalid measurement length")
	ErrMeasurementValue = errors.New("invalid measurement value")
	ErrInvalidEval      = errors.New("invalid evaluation point")
)
