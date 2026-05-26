// Package polynomial provides representations of polynomials over the scalars
// of a group.
package polynomial

import "github.com/cloudflare/circl/group"

// Polynomial stores a polynomial over the set of scalars of a group.
type Polynomial struct {
	// Internal representation is in polynomial basis:
	// Thus,
	//     p(x) = \sum_i^k c[i] x^i,
	// where k = len(c)-1 is the degree of the polynomial.
	c []group.Scalar
}

// New creates a new polynomial given its coefficients in ascending order.
// Thus,
//
//	p(x) = \sum_i^k c[i] x^i,
//
// where k = len(c)-1 is the degree of the polynomial.
//
// The zero polynomial has degree equal to -1 and can be instantiated passing
// nil to New.
func New(coeffs []group.Scalar) (p Polynomial) { _ = "STUB: not implemented"; return *new(Polynomial) }

// Degree returns the degree of the polynomial. The zero polynomial has degree
// equal to -1.
func (p Polynomial) Degree() int { _ = "STUB: not implemented"; return 0 }

// Evaluate returns the evaluation of p on x.
func (p Polynomial) Evaluate(x group.Scalar) group.Scalar {
	_ = "STUB: not implemented"
	return *new(group.Scalar)
}

// Coefficient returns a deep-copy of the n-th polynomial's coefficient.
// Note coefficients are sorted in ascending order with respect to the degree.
func (p Polynomial) Coefficient(n uint) group.Scalar {
	_ = "STUB: not implemented"
	return *new(group.Scalar)
}

// LagrangePolynomial stores a Lagrange polynomial over the set of scalars of a group.
type LagrangePolynomial struct {
	// Internal representation is in Lagrange basis:
	// Thus,
	//     p(x) = \sum_i^k y[i] L_j(x), where k is the degree of the polynomial,
	//     L_j(x) = \prod_i^k (x-x[i])/(x[j]-x[i]),
	//     y[i] = p(x[i]), and
	//     all x[i] are different.
	x, y []group.Scalar
}

// NewLagrangePolynomial creates a polynomial in Lagrange basis given a list
// of nodes (x) and values (y), such that:
//
//	p(x) = \sum_i^k y[i] L_j(x), where k is the degree of the polynomial,
//	L_j(x) = \prod_i^k (x-x[i])/(x[j]-x[i]),
//	y[i] = p(x[i]), and
//	all x[i] are different.
//
// It panics if one of these conditions does not hold.
//
// The zero polynomial has degree equal to -1 and can be instantiated passing
// (nil,nil) to NewLagrangePolynomial.
func NewLagrangePolynomial(x, y []group.Scalar) (l LagrangePolynomial) {
	_ = "STUB: not implemented"
	return *new(LagrangePolynomial)
}

func (l LagrangePolynomial) Degree() int { _ = "STUB: not implemented"; return 0 }

func (l LagrangePolynomial) Evaluate(x group.Scalar) group.Scalar {
	_ = "STUB: not implemented"
	return *new(group.Scalar)
}

// LagrangeBase returns the j-th Lagrange polynomial base evaluated at x.
// Thus, L_j(x) = \prod (x - x[i]) / (x[j] - x[i]) for 0 <= i < k, and i != j.
func LagrangeBase(jth uint, xi []group.Scalar, x group.Scalar) group.Scalar {
	_ = "STUB: not implemented"
	return *new(group.Scalar)
}

func baseRatio(jth uint, xi []group.Scalar, x group.Scalar) group.Scalar {
	_ = "STUB: not implemented"
	return *new(group.Scalar)
}

func areAllDifferent(x []group.Scalar) bool { _ = "STUB: not implemented"; return false }
