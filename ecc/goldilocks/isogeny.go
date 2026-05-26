package goldilocks

func (Curve) pull(P *twistPoint) *Point      { _ = "STUB: not implemented"; return nil }
func (twistCurve) pull(P *Point) *twistPoint { _ = "STUB: not implemented"; return nil }

// push sends a point on the Goldilocks curve to a point on the twist curve.
func (Curve) push(P *Point) *twistPoint { _ = "STUB: not implemented"; return nil }

// x+y
// A = x^2
// B = y^2
// z^2
// C = 2*z^2
// D = A
// (x+y)^2
// (x+y)^2-A
// E = (x+y)^2-A-B
// H = B+D
// G = B-D
// F = C-H
// Z = F * G
// X = E * F
// Y = G * H, // T = E * H

// push sends a point on the twist curve to a point on the Goldilocks curve.
func (twistCurve) push(P *twistPoint) *Point { _ = "STUB: not implemented"; return nil }

// x+y
// A = x^2
// B = y^2
// z^2
// C = 2*z^2
// D = -A
// (x+y)^2
// (x+y)^2-A
// E = (x+y)^2-A-B
// H = B+D
// G = B-D
// F = C-H
// Z = F * G
// X = E * F
// Y = G * H, // T = E * H
