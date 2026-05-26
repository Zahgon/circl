package fourq

type pointR1 struct {
	X, Y, Z, Ta, Tb Fq // (x,y,z,t=ta*tb)
}

type pointR3 struct {
	addYX Fq // y + x
	subYX Fq // y - x
	dt2   Fq // 2*d*t
}

type pointR2 struct {
	pointR3
	z2 Fq // 2 * z
}

// subYDiv16 update x = (x - y) / 16.
func subYDiv16(x *[5]uint64, y int64) { _ = "STUB: not implemented"; return }

// condAddOrderN updates x = x+order if x is even, otherwise x remains unchanged.
func condAddOrderN(x *[5]uint64) { _ = "STUB: not implemented"; return }

func recodeScalar(d *[65]int8, k *[32]byte) { _ = "STUB: not implemented"; return }

func (P *pointR1) oddMultiples(T *[8]pointR2) { _ = "STUB: not implemented"; return }

// scalarMult calculates P = k*Q.
func (P *pointR1) ScalarMult(k *[32]byte, Q *pointR1) { _ = "STUB: not implemented"; return }

// absolute returns always a positive value.
func absolute(x int32) int32 { _ = "STUB: not implemented"; return 0 }

// div2subY update x = (x/2) - y.
func div2subY(x *[5]uint64, y int64) { _ = "STUB: not implemented"; return }

// mLSBRecoding is the odd-only modified LSB-set.
//
// Reference:
//
//	"Efficient and secure algorithms for GLV-based scalar multiplication and
//	 their implementation on GLV–GLS curves" by (Faz-Hernandez et al.)
//	 http://doi.org/10.1007/s13389-014-0085-7.
func mLSBRecoding(L []int8, k []byte) { _ = "STUB: not implemented"; return }

// right-shift by d

func (P *pointR1) ScalarBaseMult(scalar *[Size]byte) { _ = "STUB: not implemented"; return }

func (P *pointR1) copy(Q *pointR1) { _ = "STUB: not implemented"; return }

func (P *pointR1) SetIdentity() { _ = "STUB: not implemented"; return }

func (P *pointR1) IsIdentity() bool { _ = "STUB: not implemented"; return false }

func (P *pointR1) ToAffine() { _ = "STUB: not implemented"; return }

// Marshal encodes a point P into out buffer.
func (P *Point) Marshal(out *[Size]byte) {
	_ = "STUB: not implemented"

	// b=0 if x is positive or zero
	// b=1 if x is negative
	return
}

// Unmarshal retrieves a point P from the input buffer. On success, returns true.
func (P *Point) Unmarshal(in *[Size]byte) bool { _ = "STUB: not implemented"; return false }

// t0 = y^2
// t1 = d*y^2
// t0 = y^2 - 1
// t1 = d*y^2 + 1
// x = sqrt(t0/t1)

func (P *pointR1) IsOnCurve() bool { _ = "STUB: not implemented"; return false }

// Check z != 0

// Check Eq 1: -X^2 + Y^2 == Z^2 + dT^2
// t0  = y + x
// lhs = y - x
// lhs = y^2 - x^2
// rhs = T = Ta * Tb
// rhs = T^2
// rhs = dT^2
// t0  = Z^2
// rhs = Z^2 + dT^2
// t0  = (-X^2 + Y^2) - (Z^2 + dT^2)

// Check Eq 2: (Ta*Tb)*Z == X*Y
// lhs = Ta*Tb = T
// lhs = T * Z
// rhs = X * Y
// t0  = Ta*Tb*Z - X*Y

func (P *pointR1) isEqual(Q *pointR1) bool { _ = "STUB: not implemented"; return false }

// l = X1*Z2
// r = X2*Z1
// l = l-r

// l = Y1*Z2
// r = Y2*Z1
// l = l-r

// l = T1 = Ta1*Tb1
// l = T1*Z2
// r = T2 = Ta2*Tb2
// r = T2*Z1
// l = l-r

func (P *pointR1) ClearCofactor() { _ = "STUB: not implemented"; return }

func (P *pointR2) FromR1(Q *pointR1) { _ = "STUB: not implemented"; return }

func (P *pointR2) cmov(Q *pointR2, b int) { _ = "STUB: not implemented"; return }

func (P *pointR3) cneg(b int) { _ = "STUB: not implemented"; return }

func (P *pointR3) cmov(Q *pointR3, b int) { _ = "STUB: not implemented"; return }
