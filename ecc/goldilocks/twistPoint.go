package goldilocks

import (
	fp "github.com/cloudflare/circl/math/fp448"
)

type twistPoint struct{ x, y, z, ta, tb fp.Elt }

type preTwistPointAffine struct{ addYX, subYX, dt2 fp.Elt }

type preTwistPointProy struct {
	preTwistPointAffine
	z2 fp.Elt
}

func (P *twistPoint) String() string { _ = "STUB: not implemented"; return "" }

// cneg conditionally negates the point if b=1.
func (P *twistPoint) cneg(b uint) { _ = "STUB: not implemented"; return }

// Double updates P with 2P.
func (P *twistPoint) Double() {
	_ = "STUB: not implemented"
	// This is formula (7) from "Twisted Edwards Curves Revisited" by
	// Hisil H., Wong K.KH., Carter G., Dawson E. (2008)
	// https://doi.org/10.1007/978-3-540-89255-7_20
	return
}

// x+y
// A = x^2
// B = y^2
// z^2
// C = 2*z^2
// H = A+B
// (x+y)^2
// E = (x+y)^2-A-B
// G = B-A
// F = C-G
// Z = F * G
// X = E * F
// Y = G * H, T = E * H

// mixAdd calculates P= P+Q, where Q is a precomputed point with Z_Q = 1.
func (P *twistPoint) mixAddZ1(Q *preTwistPointAffine) { _ = "STUB: not implemented"; return }

// D = 2*z1 (z2=1)

// coreAddition calculates P=P+Q for curves with A=-1.
func (P *twistPoint) coreAddition(Q *preTwistPointAffine) {
	_ = "STUB: not implemented"
	// This is the formula following (5) from "Twisted Edwards Curves Revisited" by
	// Hisil H., Wong K.KH., Carter G., Dawson E. (2008)
	// https://doi.org/10.1007/978-3-540-89255-7_20
	return
}

// t1 = ta*tb
// y1-x1
// y1+x1
// A = (y1-x1)*(y2-x2)
// B = (y1+x1)*(y2+x2)
// C = 2*D*t1*t2
// E = B-A
// H = B+A
// F = D-C
// G = D+C
// Z = F * G
// X = E * F
// Y = G * H, T = E * H

func (P *preTwistPointAffine) neg() { _ = "STUB: not implemented"; return }

func (P *preTwistPointAffine) cneg(b int) { _ = "STUB: not implemented"; return }

func (P *preTwistPointAffine) cmov(Q *preTwistPointAffine, b uint) {
	_ = "STUB: not implemented"
	return
}

// mixAdd calculates P= P+Q, where Q is a precomputed point with Z_Q != 1.
func (P *twistPoint) mixAdd(Q *preTwistPointProy) { _ = "STUB: not implemented"; return }

// D = 2*z1*z2

// oddMultiples calculates T[i] = (2*i-1)P for 0 < i < len(T).
func (P *twistPoint) oddMultiples(T []preTwistPointProy) { _ = "STUB: not implemented"; return }

// cmov conditionally moves Q into P if b=1.
func (P *preTwistPointProy) cmov(Q *preTwistPointProy, b uint) { _ = "STUB: not implemented"; return }

// FromTwistPoint precomputes some coordinates of Q for missed addition.
func (P *preTwistPointProy) FromTwistPoint(Q *twistPoint) { _ = "STUB: not implemented"; return }

// addYX = X + Y
// subYX = Y - X
// T = ta*tb
// D*T
// dt2 = 2*D*T
// z2 = 2*Z
