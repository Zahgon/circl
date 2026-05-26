package bls12381

import "github.com/cloudflare/circl/ecc/bls12381/ff"

// Pair calculates the ate-pairing of P and Q.
func Pair(P *G1, Q *G2) *Gt { _ = "STUB: not implemented"; return nil }

func miller(f *ff.Fp12, P *G1, Q *G2) { _ = "STUB: not implemented"; return }

// paramX is -2 ^ 63 - 2 ^ 62 - 2 ^ 60 - 2 ^ 57 - 2 ^ 48 - 2 ^ 16

// inverts f as paramX is negative.

// line contains the coefficients of a sparse element of Fp12.
// Evaluating the line on P' = (xP',yP') results in
//
//	f = evalLine(P') = l[0]*xP' + l[1]*yP' + l[2] \in Fp12.
type line [3]ff.Fp2

// evalLine updates f = f * line(P'), where f lives in Fp12 = Fp6[w]/(w^2-v)
// and P' is the image of P on the twist curve.
func evalLine(f *ff.LineValue, l *line, P *G1) {
	_ = "STUB: not implemented"
	// Send P \in E to the twist
	//
	//	   E    -->        E'
	//	(xP,yP) |-> (xP*w^2,yP*w^3) = (xP',yP')
	//
	// f = line(P') = l[0]*xP' + l[1]*yP' + l[2] \in Fp12.
	//
	//	= l[0]*xP*w^2 + l[1]*yP*w^3 + l[2] \in Fp12.
	return
}

// First perform the products: l[0]*xP and l[1]*yP \in Fp2.

func finalExp(g *Gt, f *ff.Fp12) { _ = "STUB: not implemented"; return }

// ProdPair calculates the product of pairings:
//
//	e = \Prod_i pair(Pi, Qi)^ni
//	  = \Prod_i pair(ni*Pi, Qi)
//	  = \Prod_i pair(Pi, ni*Qi)
//
// For efficiency, it performs operations in G1.
func ProdPair(P []*G1, Q []*G2, n []*Scalar) *Gt { _ = "STUB: not implemented"; return nil }

// ProdPairFrac computes the product e(P, Q)^sign where sign is 1 or -1
func ProdPairFrac(P []*G1, Q []*G2, signs []int) *Gt { _ = "STUB: not implemented"; return nil }
