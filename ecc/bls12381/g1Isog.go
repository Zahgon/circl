package bls12381

import (
	"github.com/cloudflare/circl/ecc/bls12381/ff"
)

type isogG1Point struct{ x, y, z ff.Fp }

func (p isogG1Point) String() string { _ = "STUB: not implemented"; return "" }

// IsOnCurve returns true if g is a valid point on the curve.
func (p *isogG1Point) IsOnCurve() bool { _ = "STUB: not implemented"; return false }

// y2 = y^2
// y2 = y^2*z
// z2 = z^2
// z3 = z^3
// z3 = B*z^3
// x2 = x^2
// x3 = A*z^2
// x3 = x^2 + A*z^2
// x3 = x^3 + A*x*z^2
// x3 = x^3 + A*x*z^2 + Bz^3

// sswu implements the Simplified Shallue-van de Woestijne-Ulas method for
// mapping a field element to a point on the isogenous curve.
func (p *isogG1Point) sswu(u *ff.Fp) {
	_ = "STUB: not implemented"
	// Method in Appendix-G.2.1 of
	// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-hash-to-curve-11
	return
}

// 1.  tv1 = u^2
// 2.  tv3 = Z * tv1
// 3.  tv2 = tv3^2
// 4.   xd = tv2 + tv3
// 5.  tv4 = 1
//     x1n = xd + tv4
// 6.  x1n = x1n * B
// 7.   xd = A * xd
//      xd = -xd
// 8.   e1 = xd == 0
// 9.  tv4 = Z * A
//      xd = CMOV(xd, tv4, e1)
// 10. tv2 = xd^2
// 11. gxd = tv2 * xd
// 12. tv2 = A * tv2
// 13. gx1 = x1n^2
// 14. gx1 = gx1 + tv2
// 15. gx1 = gx1 * x1n
// 16. tv2 = B * gxd
// 17. gx1 = gx1 + tv2
// 18. tv4 = gxd^2
// 19. tv2 = gx1 * gxd
// 20. tv4 = tv4 * tv2
// 21.  y1 = tv4^c1
// 22.  y1 = y1 * tv2
// 23. x2n = tv3 * x1n
// 24.  y2 = y1 * c2
// 25.  y2 = y2 * tv1
// 26.  y2 = y2 * u
// 27. tv2 = y1^2
// 28. tv2 = tv2 * gxd
// 29.  e2 = tv2 == gx1
// 30.  xn = CMOV(x2n, x1n, e2)
// 31.   y = CMOV(y2, y1, e2)
// 32.  e3 = sgn0(u) == sgn0(y)
// 33. tv1 = y
//     tv1 = -y
//       y = CMOV(tv1, y, e3)
// 34. return
//       (x,y) = (xn/xd, y/1)
//       (X,Y,Z) = (xn, y*xd, xd)

// evalIsogG1 calculates g = g1Isog11(p), where g1Isog11 is an isogeny of
// degree 11 to the curve used in G1.
func (g *G1) evalIsogG1(p *isogG1Point) { _ = "STUB: not implemented"; return }
