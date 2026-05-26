package bls12381

import (
	"github.com/cloudflare/circl/ecc/bls12381/ff"
)

type isogG2Point struct{ x, y, z ff.Fp2 }

func (p isogG2Point) String() string { _ = "STUB: not implemented"; return "" }

// IsOnCurve returns true if g is a valid point on the curve.
func (p *isogG2Point) IsOnCurve() bool { _ = "STUB: not implemented"; return false }

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
func (p *isogG2Point) sswu(u *ff.Fp2) {
	_ = "STUB: not implemented"
	// Method in Appendix-G.2.3 of
	// https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-hash-to-curve-11
	return
}

// 1.  tv1 = u^2
// 2.  tv3 = Z * tv1
// 3.  tv5 = tv3^2
// 4.   xd = tv5 + tv3
// 5.  tv2 = 1
//     x1n = xd + tv2
// 6.  x1n = x1n * B
// 7.   xd = A * xd
//      xd = -xd
// 8.   e1 = xd == 0
// 9.  tv2 = Z * A
//      xd = CMOV(xd, tv2, e1)
// 10. tv2 = xd^2
// 11. gxd = tv2 * xd
// 12. tv2 = A * tv2
// 13. gx1 = x1n^2
// 14. gx1 = gx1 + tv2
// 15. gx1 = gx1 * x1n
// 16. tv2 = B * gxd
// 17. gx1 = gx1 + tv2
// 18. tv4 = gxd^2
// 19. tv2 = tv4 * gxd
// 20. tv4 = tv4^2
// 21. tv2 = tv2 * tv4
// 22. tv2 = tv2 * gx1
// 23. tv4 = tv4^2
// 24. tv4 = tv2 * tv4
// 25.   y = tv4^c1
// 26.   y = y * tv2
// 27. tv4 = y * c2
// 28. tv2 = tv4^2
// 29. tv2 = tv2 * gxd
// 30.  e2 = tv2 == gx1
// 31.   y = CMOV(y, tv4, e2)
// 32. tv4 = y * c3
// 33. tv2 = tv4^2
// 34. tv2 = tv2 * gxd
// 35.  e3 = tv2 == gx1
// 36.   y = CMOV(y, tv4, e3)
// 37. tv4 = tv4 * c2
// 38. tv2 = tv4^2
// 39. tv2 = tv2 * gxd
// 40.  e4 = tv2 == gx1
// 41.   y = CMOV(y, tv4, e4)
// 42. gx2 = gx1 * tv5
// 43. gx2 = gx2 * tv3
// 44. tv5 = y * tv1
// 45. tv5 = tv5 * u
// 46. tv1 = tv5 * c4
// 47. tv4 = tv1 * c2
// 48. tv2 = tv4^2
// 49. tv2 = tv2 * gxd
// 50.  e5 = tv2 == gx2
// 51. tv1 = CMOV(tv1, tv4, e5)
// 52. tv4 = tv5 * c5
// 53. tv2 = tv4^2
// 54. tv2 = tv2 * gxd
// 55.  e6 = tv2 == gx2
// 56. tv1 = CMOV(tv1, tv4, e6)
// 57. tv4 = tv4 * c2
// 58. tv2 = tv4^2
// 59. tv2 = tv2 * gxd
// 60.  e7 = tv2 == gx2
// 61. tv1 = CMOV(tv1, tv4, e7)
// 62. tv2 = y^2
// 63. tv2 = tv2 * gxd
// 64.  e8 = tv2 == gx1
// 65.   y = CMOV(tv1, y, e8)
// 66. tv2 = tv3 * x1n
// 67.  xn = CMOV(tv2, x1n, e8)
// 68.  e9 = sgn0(u) == sgn0(y)
// 69. tv1 = y
//     tv1 = -y
//       y = CMOV(tv1, y, e9)
// 70. return
//       (x,y) = (xn/xd, y/1)
//       (X,Y,Z) = (xn, y*xd, xd)

// evalIsogG2 calculates g = g2Isog3(p), where g2Isog3 is an isogeny of
// degree 3 to the curve used in G2.
func (g *G2) evalIsogG2(p *isogG2Point) { _ = "STUB: not implemented"; return }
