package csidh

// xAdd implements differential arithmetic in P^1 for Montgomery
// curves E(x): x^3 + A*x^2 + x by using x-coordinate only arithmetic.
//
//	x(PaQ) = x(P) + x(Q) by using x(P-Q)
//
// This algorithms is correctly defined only for cases when
// P!=inf, Q!=inf, P!=Q and P!=-Q.
func xAdd(PaQ, P, Q, PdQ *point) { _ = "STUB: not implemented"; return }

// sqr
// sqr

// xDbl implements point doubling on a Montgomery curve
// E(x): x^3 + A*x^2 + x by using x-coordinate only arithmetic.
//
//	x(Q) = [2]*x(P)
//
// It is correctly defined for all P != inf.
func xDbl(Q, P, A *point) { _ = "STUB: not implemented"; return }

// sqr

// sqr

// xDblAdd implements combined doubling of point P
// and addition of points P and Q on a Montgomery curve
// E(x): x^3 + A*x^2 + x by using x-coordinate only arithmetic.
//
//	x(PaP) = x(2*P)
//	x(PaQ) = x(P+Q)
func xDblAdd(PaP, PaQ, P, Q, PdQ *point, A24 *coeff) { _ = "STUB: not implemented"; return }

// cswappoint swaps P1 with P2 in constant time. The 'choice'
// parameter must have a value of either 1 (results
// in swap) or 0 (results in no-swap).
func cswappoint(P1, P2 *point, choice uint8) { _ = "STUB: not implemented"; return }

// xMul implements point multiplication with left-to-right Montgomery
// adder. co is A coefficient of x^3 + A*x^2 + x curve. k must be > 0
//
// Non-constant time!
func xMul(kP, P *point, co *coeff, k *fp) { _ = "STUB: not implemented"; return }

// Precompyte A24 = (A+2C:4C) => (A24.x = A.x+2A.z; A24.z = 4*A.z)

// Skip initial 0 bits.

// performance hit from making it constant-time is actually
// quite big, so... unsafe branch for now

// xIso computes the isogeny with kernel point kern of a given order
// kernOrder. Returns the new curve coefficient co and the image img.
//
// During computation function switches between Montgomery and twisted
// Edwards curves in order to compute image curve parameters faster.
// This technique is described by Meyer and Reith in ia.cr/2018/782.
//
// Non-constant time.
func xIso(img *point, co *coeff, kern *point, kernOrder uint64) { _ = "STUB: not implemented"; return }

// Compute twisted Edwards coefficients
// coEd.a = co.a + 2*co.c
// coEd.c = co.a - 2*co.c
// coEd.a*X^2 + Y^2 = 1 + coEd.c*X^2*Y^2

// Transfer point to twisted Edwards YZ-coordinates
// (X:Z)->(Y:Z) = (X-Z : X+Z)

// NOTE: Not constant time.

// coEd.a^kernOrder and coEd.c^kernOrder

// prod^8

// Compute image curve params

// Convert curve coefficients back to Montgomery

// montEval evaluates x^3 + Ax^2 + x.
func montEval(res, A, x *fp) { _ = "STUB: not implemented"; return }
