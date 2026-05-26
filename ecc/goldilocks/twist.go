package goldilocks

// twistCurve is -x^2+y^2=1-39082x^2y^2 and is 4-isogenous to Goldilocks.
type twistCurve struct{}

// Identity returns the identity point.
func (twistCurve) Identity() *twistPoint { _ = "STUB: not implemented"; return nil }

// subYDiv16 update x = (x - y) / 16.
func subYDiv16(x *scalar64, y int64) { _ = "STUB: not implemented"; return }

func recodeScalar(d *[113]int8, k *Scalar) { _ = "STUB: not implemented"; return }

// ScalarMult returns kP.
func (e twistCurve) ScalarMult(k *Scalar, P *twistPoint) *twistPoint {
	_ = "STUB: not implemented"
	return nil
}

const (
	omegaFix = 7
	omegaVar = 5
)

// CombinedMult returns mG+nP.
func (e twistCurve) CombinedMult(m, n *Scalar, P *twistPoint) *twistPoint {
	_ = "STUB: not implemented"
	return nil
}

// Generator point

// Variable input point

// absolute returns always a positive value.
func absolute(x int32) int32 { _ = "STUB: not implemented"; return 0 }
