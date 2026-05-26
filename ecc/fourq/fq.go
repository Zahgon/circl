package fourq

import (
	"math/big"
)

// Fq implements operations of a field of size q=p^2 as a quadratic
// extension of the base field where i^2=-1.
// An element in Fq is represented as f[0]+f[1]*i, where f[0],f[1] are in Fp.
type Fq [2]Fp

func (e *Fq) String() string              { _ = "STUB: not implemented"; return "" }
func (e *Fq) toBigInt() (f0, f1 *big.Int) { _ = "STUB: not implemented"; return nil, nil }
func (e *Fq) setBigInt(f0, f1 *big.Int)   { _ = "STUB: not implemented"; return }
func (e *Fq) setZero()                    { _ = "STUB: not implemented"; return }
func (e *Fq) setOne()                     { _ = "STUB: not implemented"; return }
func (e *Fq) isZero() bool                { _ = "STUB: not implemented"; return false }

func (e *Fq) toBytes(buf []byte) { _ = "STUB: not implemented"; return }

func (e *Fq) fromBytes(buf []byte) bool { _ = "STUB: not implemented"; return false }

func fqSgn(c *Fq) int { _ = "STUB: not implemented"; return 0 }

func fqCopy(c, a *Fq) { _ = "STUB: not implemented"; return }
func fqNeg(c, a *Fq)  { _ = "STUB: not implemented"; return }

// fqSqrt calculates c = sqrt(u/v) such that sgn(c)=s.
func fqSqrt(c, u, v *Fq, s int) { _ = "STUB: not implemented"; return }

// a = u0*v0 + u1*v1

// b = v0^2 + v1^2

// g = u1*v0 - u0*v1

// t = 2(a + sqrt(a^2+g^2)) = 2*(a + (a^2+g^2)^(2^125))
// if t=0; then t = 2*(a - (a^2+g^2)^(2^125))

// r = (t*b^3)^(2^125-1)

// x0 = (r*b*t)/2
// x1 = (r*b*g)

// if b*(2*x0)^2 == t then (x0,x1) <- (x1,x0)

func fqInv(c, a *Fq) { _ = "STUB: not implemented"; return }
