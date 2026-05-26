// Package qndleq provides zero-knowledge proofs of Discrete-Logarithm Equivalence (DLEQ) on Qn.
//
// This package implements proofs on the group Qn (the subgroup of squares in (Z/nZ)*).
//
// # Notation
//
//	Z/nZ is the ring of integers modulo N.
//	(Z/nZ)* is the multiplicative group of Z/nZ, a.k.a. the units of Z/nZ, the elements with inverse mod N.
//	Qn is the subgroup of squares in (Z/nZ)*.
//
// A number x belongs to Qn if
//
//	gcd(x, N) = 1, and
//	exists y such that x = y^2 mod N.
//
// # References
//
// [DLEQ Proof] "Wallet databases with observers" by Chaum-Pedersen.
// https://doi.org/10.1007/3-540-48071-4_7
//
// [Qn] "Practical Threshold Signatures" by Shoup.
// https://www.iacr.org/archive/eurocrypt2000/1807/18070209-new.pdf
package qndleq

import (
	"errors"
	"io"
	"math/big"
)

type Proof struct {
	z, c     *big.Int
	secParam uint
}

// SampleQn returns an element of Qn (the subgroup of squares in (Z/nZ)*).
// SampleQn will return error for any error returned by crypto/rand.Int.
func SampleQn(random io.Reader, N *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// x is a square by construction.

// now check whether h is coprime to N.

// Prove creates a DLEQ Proof that attests that the pairs (g,gx)
// and (h,hx) have the same discrete logarithm equal to x.
//
// Given g, h in Qn (the subgroup of squares in (Z/nZ)*), it holds
//
//	gx = g^x mod N
//	hx = h^x mod N
//	x  = Log_g(g^x) = Log_h(h^x)
//
// Note: this function does not run in constant time because it uses
// big.Int arithmetic.
func Prove(random io.Reader, x, g, gx, h, hx, N *big.Int, secParam uint) (*Proof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Challenge must not be congruent to zero.
//   c != 0 mod m, where m = (p-1)(q-1)/4, and N = p*q.
// Check this by doing an Exp because m is unknown.
//
// This is valid assuming N is the product of two safe prime numbers.
// In the verification equation, c multiplies the witness.
// When c is zero, it removes the witness allowing to trivially
// pass the verification check.

// Verify checks whether x = Log_g(g^x) = Log_h(h^x).
func (p Proof) Verify(g, gx, h, hx, N *big.Int) bool { _ = "STUB: not implemented"; return false }

// Check c != 0 (mod m), where m = (p-1)(q-1)/4,
// by doing an Exp as m is unknown.

func doChallenge(g, gx, h, hx, gP, hP, N *big.Int, secParam uint) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checkBounds returns nil if 0 < x[i] < N for all 0 <= i < len(x);
// otherwise, returns ErrBounds.
func checkBounds(N *big.Int, x ...*big.Int) error { _ = "STUB: not implemented"; return nil }

var (
	// ErrSecParam is returned when the security parameter is less than 128.
	ErrSecParam = errors.New("zk/qndleq: the security parameter must be greater than 128")
	// ErrBounds is returned when a value is not in the range 0 to N.
	ErrBounds = errors.New("zk/qndleq: input must be greater than 0 and less than N")
	// ErrProve is returned when Prove exhausted the number of proof tries.
	ErrProve = errors.New("zk/qndleq: exhausted the number of proof tries")
)
