package bls12381

import (
	"errors"

	"github.com/cloudflare/circl/ecc/bls12381/ff"
)

// Scalar represents positive integers in the range 0 <= x < Order.
type Scalar = ff.Scalar

const ScalarSize = ff.ScalarSize

// Order returns the order of the pairing groups, returned as a big-endian slice.
//
//	Order = 0x73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001
func Order() []byte { _ = "STUB: not implemented"; return nil }

var (
	bls12381 struct { // Let z be the BLS12 parameter.
		minusZ    [8]byte  //      (-z), (integer big-endian).
		oneMinusZ [8]byte  //     (1-z), (integer big-endian).
		g1Check   [16]byte // (z^2-1)/3, (integer big-endian).
	}
	g1Params struct{ b, _3b, genX, genY ff.Fp }
	g2Params struct{ b, _3b, genX, genY ff.Fp2 }

	// g1Isog11 is an isogeny of degree 11 from g1Iso(a,b) to G1 and is given
	// by rational maps:
	//  g1Iso(a,b) --> G1
	//  (x,y,z)    |-> (x,y,1)
	//                 (xNum/xDen, y * yNum/yDen, 1)
	//                 (xNum*yDen, y * yNum*xDen, z*xDen*yDen)
	// such that
	//  xNum = \sum ai * x^i * z^(n-1-i), for 0 <= i < n, and n=12.
	//  xDen = \sum bi * x^i * z^(n-1-i), for 0 <= i < n, and n=11.
	//  yNum = \sum ci * x^i * z^(n-1-i), for 0 <= i < n, and n=16.
	//  yDen = \sum di * x^i * z^(n-1-i), for 0 <= i < n, and n=16.
	g1Isog11 struct {
		a, b ff.Fp
		xNum [12]ff.Fp
		xDen [11]ff.Fp
		yNum [16]ff.Fp
		yDen [16]ff.Fp
	}

	// g2Isog3 is an isogeny of degree 3 from g2Iso(a,b) to G2 and is given
	// by rational maps:
	//  g2Iso(a,b) --> G2
	//  (x,y,z)    |-> (x,y,1)
	//                 (xNum/xDen, y * yNum/yDen, 1)
	//                 (xNum*yDen, y * yNum*xDen, z*xDen*yDen)
	// such that
	//  xNum = \sum ai * x^i * z^(n-1-i), for 0 <= i < n, and n=4.
	//  xDen = \sum bi * x^i * z^(n-1-i), for 0 <= i < n, and n=3.
	//  yNum = \sum ci * x^i * z^(n-1-i), for 0 <= i < n, and n=4.
	//  yDen = \sum di * x^i * z^(n-1-i), for 0 <= i < n, and n=4.
	g2Isog3 struct {
		a, b ff.Fp2
		xNum [4]ff.Fp2
		xDen [3]ff.Fp2
		yNum [4]ff.Fp2
		yDen [4]ff.Fp2
	}
	g1sswu struct {
		Z  ff.Fp    // Z = 11.
		c1 [48]byte // integer c1 = (p - 3) / 4 (big-endian)
		c2 ff.Fp
	}
	g2sswu struct {
		Z  ff.Fp2   // -(2 + I)
		c1 [95]byte // integer c1 = (p^2 - 9) / 16 (big-endian)
		c2 ff.Fp2   // sqrt(-1)
		c3 ff.Fp2   // sqrt(c2)
		c4 ff.Fp2   // sqrt(Z^3 / c3)
		c5 ff.Fp2   // sqrt(Z^3 / (c2 * c3))
	}
	g1Sigma struct {
		beta0 ff.Fp // beta0 = F(2)^(2*(p-1)/3) where F = GF(p).
		beta1 ff.Fp // beta1 = F(2)^(1*(p-1)/3) where F = GF(p).
	}
	g2Psi struct {
		alpha ff.Fp2 // alpha = w^2/Frob(w^2)
		beta  ff.Fp2 // beta = w^3/Frob(w^3)
	}
)

var (
	errInputLength = errors.New("incorrect input length")
	errEncoding    = errors.New("incorrect encoding")
)

func headerEncoding(isCompressed, isInfinity, isBigYCoord byte) byte {
	_ = "STUB: not implemented"
	return 0
}

func err(e error) { _ = "STUB: not implemented"; return }

func init() {
	bls12381.oneMinusZ = [8]byte{ // (big-endian)
		0xd2, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01,
	}
	bls12381.minusZ = [8]byte{ // (big-endian)
		0xd2, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	}
	bls12381.g1Check = [16]byte{ // (big-endian)
		0x39, 0x6c, 0x8c, 0x00, 0x55, 0x55, 0xe1, 0x56,
		0x00, 0x00, 0x00, 0x00, 0x55, 0x55, 0x55, 0x55,
	}
	initG1Params()
	initG2Params()
	initG1Isog11()
	initG2Isog3()
	initG1sswu()
	initG2sswu()
	initSigma()
	initPsi()
}

func initG1Params() { _ = "STUB: not implemented"; return }

func initG2Params() { _ = "STUB: not implemented"; return }

func initG1Isog11() { _ = "STUB: not implemented"; return }

func initG2Isog3() { _ = "STUB: not implemented"; return }

func initG1sswu() { _ = "STUB: not implemented"; return }

// (big-endian)

func initG2sswu() { _ = "STUB: not implemented"; return }

// (big-endian)

func initSigma() { _ = "STUB: not implemented"; return }

func initPsi() {
	_ = "STUB: not implemented"
	// ratioKummer sets z = t/Frob(t) if it falls in Fp2, panics otherwise.
	return
}
