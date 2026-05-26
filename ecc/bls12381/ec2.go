package bls12381

func doubleAndLine(P *G2, l *line) {
	_ = "STUB: not implemented"
	// Reference:
	//
	//	"Faster Pairing Computations on Curves with High-Degree Twists" by
	//	Costello-Lange-Naehrig. [Sec. 5] (eprint.iacr.org/2009/615).
	//	"Complete addition formulas for prime order elliptic curves" by
	//	Costello-Renes-Batina. [Alg.9] (eprint.iacr.org/2015/1060).
	return
}

// 1. B = Y1^2
// 2. C = Z1^2
// 3. D = 3b*C
// 4. F = (Y1+Z1)
//    F = (Y1+Z1)^2
//    F = (Y1+Z1)^2-B
//    F = (Y1+Z1)^2-B-C

// 5.  A  = X1^2
//     E  = (X1+Y1)
//     E  = (X1+Y1)^2
//     E  = (X1+Y1)^2-A
//     E  = (X1+Y1)^2-A-B = 2X*Y
// 5a. l0 = 2A
//     l0 = 3A = 3X1^2
// 5b. l1 = F
//     l1 = -F = -2Y1Z1
// 5c. l2 = D-B = 3b*Z1^2-Y1^2

// 5. E = X*Y
//    E = 2X*Y

// 6.  T  = 2D
// 7.  G  = 3D
// 8.  X3 = (B-G)
//     X3 = E*(B-G)
// 9 .  T = 4D^2
// 10. Y3 = (B+G)
//     Y3 = (B+G)^2
//     Y3 = (B+G)^2-4D^2
//     Y3 = (B+G)^2-8D^2
//     Y3 = (B+G)^2-12D^2
// 11. Z3 = B*F
//     Z3 = 2B*F
//     Z3 = 4B*F

func addAndLine(PQ, P, Q *G2, l *line) {
	_ = "STUB: not implemented"
	// Reference:
	//
	//	"Faster Pairing Computations on Curves with High-Degree Twists" by
	//	Costello-Lange-Naehrig. [Sec. 5] (eprint.iacr.org/2009/615).
	//	"Complete addition formulas for prime order elliptic curves" by
	//	Costello-Renes-Batina. [Alg.7] (eprint.iacr.org/2015/1060).
	return
}

// A = 2X1X2
//   = 3X1X2
// B = Y1Y2+3bZ1Z2
// C = Y1Y2-3bZ1Z2

// t0 = (X1 + Y1)
// D  = (X2 + Y2)
//    = X1X2 + X1Y2 + X2Y1 + Y1Y2
//    = X1Y2 + X2Y1 + Y1Y2
//    = X1Y2 + X2Y1

// t0 = Y1Z2
// t1 = Y2Z1
// E  = Y1Z2 + Y2Z1
// EE = Y1Z2 - Y2Z1

// t0 = X1Z2
// t1 = X2Z1
// F  = X1Z2 + X2Z1
// FF = X1Z2 - X2Z1

// l0 = (Y1Z2 - Y2Z1)*Z2
//    = -(Y1Z2 - Y2Z1)*Z2
// l1 = (X1Z2 - X2Z1)*Z2
// t0 = (X1Z2 - X2Z1)*Y2
// l2 = (Y1Z2 - Y2Z1)*X2
//    = (Y1Z2 - Y2Z1)*X2 - (X1Z2 - X2Z1)*Y2

// t0 = (Y1 + Z1)
// t1 = (Y2 + Z2)
// E  = Y1Y2 + Y1Z2 + Y2Z1 + Z1Z2
//    = Y1Z2 + Y2Z1 + Z1Z2
//    = Y1Z2 + Y2Z1

// t0 = (X1 + Z1)
// t1 = (X2 + Z2)
// F  = X1X2 + X1Z2 + X2Z1 + Z1Z2
//    = X1Z2 + X2Z1 + Z1Z2
//    = X1Z2 + X2Z1

// G = 3b*F

// t0 = E*G
// X3 = D*C
//    = D*C - E*G

// t0 = A*G
// Y3 = B*C
//    = B*C + A*G

// t0 = A*D
// Z3 = E*B
//    = E*B + A*D
