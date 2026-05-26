package ed25519

import fp "github.com/cloudflare/circl/math/fp25519"

type (
	pointR1 struct{ x, y, z, ta, tb fp.Elt }
	pointR2 struct {
		pointR3
		z2 fp.Elt
	}
)
type pointR3 struct{ addYX, subYX, dt2 fp.Elt }

func (P *pointR1) neg() { _ = "STUB: not implemented"; return }

func (P *pointR1) SetIdentity() { _ = "STUB: not implemented"; return }

func (P *pointR1) toAffine() { _ = "STUB: not implemented"; return }

func (P *pointR1) ToBytes(k []byte) error { _ = "STUB: not implemented"; return nil }

func (P *pointR1) FromBytes(k []byte) bool { _ = "STUB: not implemented"; return false }

// u = y^2
// v = dy^2
// u = y^2-1
// v = dy^2+1
// x = sqrt(u/v)

// x = x mod p

// double calculates 2P for curves with A=-1.
func (P *pointR1) double() { _ = "STUB: not implemented"; return }

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

func (P *pointR1) mixAdd(Q *pointR3) { _ = "STUB: not implemented"; return }

// D = 2*z1

func (P *pointR1) add(Q *pointR2) { _ = "STUB: not implemented"; return }

// D = 2*z1*z2

// coreAddition calculates P=P+Q for curves with A=-1.
func (P *pointR1) coreAddition(Q *pointR3) { _ = "STUB: not implemented"; return }

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

func (P *pointR1) oddMultiples(T []pointR2) { _ = "STUB: not implemented"; return }

func (P *pointR1) isEqual(Q *pointR1) bool { _ = "STUB: not implemented"; return false }

func (P *pointR3) neg() { _ = "STUB: not implemented"; return }

func (P *pointR2) fromR1(Q *pointR1) { _ = "STUB: not implemented"; return }

func (P *pointR3) cneg(b int) { _ = "STUB: not implemented"; return }

func (P *pointR3) cmov(Q *pointR3, b int) { _ = "STUB: not implemented"; return }
