//go:generate go run src.go -out ../amd64.s -stubs ../stubs_amd64.go -pkg common

// AVX2 optimized version of polynomial operations.  See the comments on the
// generic implementation for the details of the maths involved.
package main

import (
	. "github.com/mmcloughlin/avo/build"   // nolint:golint,stylecheck
	. "github.com/mmcloughlin/avo/operand" // nolint:golint,stylecheck
	. "github.com/mmcloughlin/avo/reg"     // nolint:golint,stylecheck
)

// XXX align Poly on 16 bytes such that we can use aligned moves
// XXX ensure Zetas and ZetasAVX2 are 16 byte aligned

// Barrett reduces the int16x16 where q must contain {q, q, …};
// num must contain {20159, 20159, …} the numerator in the approximation
// 20159/2²⁶ of 1/q and t is a temporary register that will be clobbered.
func barrettReduceX16(x, q, num, t Op) {
	_ = "STUB: not implemented"
	// Recall that the Barrett reduction of x is given by
	//
	//	x - int16((int32(x)*20159)>>26)*q
	return
}

// t := (int32(x) * 20159) >> 16
// t = int16(t)>>10 so that t = (int32(x)*20159) >> 26
// t *= q
// x -= t

func broadcastImm16(c int16, out Op) { _ = "STUB: not implemented"; return }

func addAVX2() { _ = "STUB: not implemented"; return }

func subAVX2() { _ = "STUB: not implemented"; return }

// For each lane in a that has a 1 on ith bit of its index, swap it with
// the corresponding lane in b where this 1 has been replaced by a 0.
//
// For instance, if i=2, then this will swap
//
//	a[0b0100] ↔ b[0b0000]    a[0b0101] ↔ b[0b0001]
//	a[0b0110] ↔ b[0b0010]    a[0b0111] ↔ b[0b0011]
//	a[0b1100] ↔ b[0b1000]    a[0b1101] ↔ b[0b1001]
//	a[0b1110] ↔ b[0b1010]    a[0b1111] ↔ b[0b1011]
//
// and keep all other lanes in their place.  If we index the lanes of a and
// b consecutively (i.e. 0wxyz is wxyz of a and 1wxyz = wxyz of b), then
// this corresponds to mapping lane vwxyz to xwvyz -- that is: flipping the
// fourth bit with the ith bit (where we start from zeroth.)  Hence the name.
//
// Why these permutations?  There are two reasons: these are reasonable
// easy to implement and they pull sequential butterflies in the NTT apart.
// Recall, namely, that on the fifth layer of the NTT we're computing
// butterflies between indices
//
//	abcd0fgh  abcd1fgh
//
// Applying bitflip with i=3 beforehand, the butterflies become
//
//	abc0dfgh  abc1dfgh
//
// which allows for 16 consecutive butterflies, which is convenient for AVX2.
// What we'll actually end up doing is a bit different: we'll apply both
// an i=3 and i=2 bitflip before then to also interleave the ζs correctly
// for the fourth layer.
//
// See the diagram linked to in the documentation of nttAVX2().
func bitflip(i int, a, b, t Op) { _ = "STUB: not implemented"; return }

func invNttAVX2() {
	_ = "STUB: not implemented"
	// This AVX2-optimized inverse NTT is close, but more disimilar from
	// the generic inverse NTT, then the AVX2-optimized forward NTT is
	// from the generic.
	//
	//  1. Just like in the AVX2-optimized forward NTT, we shuffle the
	//     coefficients around to ensure we can do consecutive butterflies and
	//  2. we use the same preshuffled and duplicated ZetasAVX2 table.
	//  3. Barrett reductions are computed at different moments as it's very
	//     efficient to do 16 at a time.
	return
}

// The butterflies and swaps are in the exact reverse order as those
// of the AVX2-optimized forward NTT.  See the comments on nttAVX2()
// and bitflip() for documentation on the shufflings.

// A diagram of the order of the butterflies and swaps can be found here:
//
//  https://github.com/cloudflare/circl/wiki/images/kyber-invntt-avx2.svg
//
// The vertical lines with circles on the end represent butterflies.
// The number in those butterflies refers to the index into the Zetas
// array of which ζ is used.  (Note that this array is different from
// the ZetasAVX2 array, which contains the elements of Zetas many times
// over in a way that is efficient for our implementation.)
//
// The green squares represent Barrett reductions.  The green numbers after
// the butterflies show the multiple of q that bounds the coefficient in
// absolute value.  (Recall that the lower coefficient is always bounded
// by one when computing the inverse butterflies in the obvious way.)
// The vertical lines with crosses on them represent a swap.

// Compute 4x16 Gentleman--Sande butterflies (a, b) ↦ (a + b, ζ(a - b)).
//
// There is a catch: the first two and the last two sets of butterflies
// have to use the same sets of zetas, as we don't have enough registers
// to keep everything around.  t1 up to t4 are temporary registers that
// will be clobbered.

// In the generic implementation, a single butterfly is computed as
// follows (unfolding the definition of montReduce and recalling
// zeta stores -ζ.)
//
//  t := b - a
//  a += b
//  m := int16(zeta * t * 62209)
//  b = int16(uint32(zeta * int32(t) - m * int32(Q)) >> 16)
//
// As ζt ≡ mq (mod 2¹⁶), see comments on montReduce(), we can
// also compute b as
//
//  b = (uint32(zeta * int32(t)) >> 16) - (uint32(m * int32(Q)) >> 16)
//
// m (x16) can be computed using a single VPMULLW with zeta * 62209
// as the second operand stored in a table.  The two multiplications
// and bitshifts for b can be performed using two VPMULHWs (again
// for 16 at a time.)

// t = b - a

// We don't use t4 yet, so that zeta12l may be used as t4.

// a += b

// m = int16(zeta * t * 62209)

// At this point zeta12l (which might equal t4) is free.

// uint32(zeta*int32(t)) >> 16

// uint32(m*int32(Q)) >> 16

// Compute b

// Registers and constants

// Layers 1 - 6

// Layer 1 (inverse of 7)

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 2 (inverse of 6)

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 3 (inverse of 5)

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 4 (inverse of 4)

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 5 (inverse of 3)

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 6 (inverse of 2)

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layers 7 (inverse of 1)

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Finally, we set x = montgomeryReduce(x * 1441). Just like in the
// the butterflies we compute the Montgomery reduction using
// VPMULHWs and VPMULLWs by observing:
//
//  m := int16(1441 * x * 62209) = int16(-10079 * x)
//  x' = int16(uint32(1441 * int32(x) - m * int32(Q)) >> 16)
//     = (uint32(1441 * int32(x)) >> 16) - (uint32(m * int32(Q)) >> 16)

// m = int16(-10079 * x)

// uint32(1441*int32(x)) >> 16

// uint32(m*int32(Q)) >> 16

// computes t

func nttAVX2() {
	_ = "STUB: not implemented"
	// We perform almost the same operations as the generic implementation of NTT,
	// but use AVX2 to perform 64 butterflies at the same time.  We can keep
	// 128 coefficients in registers at the same time.  We do the first level
	// separately writing back to memory.  Then we do levels 2 through 7 for
	// 128 coefficients in registers all at the same time.
	return
}

// As we can only perform butterflies of 8 consecutive coefficients,
// we need to shuffle coefficients around.  Similarly parallel
// multiplication in the NTT-domain (as implemented by MulHat) requires
// sequential coefficients to be pulled apart.  Thus we will use a
// different order of coefficients than the reference implementation.

// A diagram of the order of butterflies and swaps can be found here:
//
//  https://github.com/cloudflare/circl/wiki/images/kyber-ntt-avx2.svg
//
// The vertical lines with circles on the end represent butterflies.
// The number in those butterflies refers to the index into the Zetas
// array of which ζ is used.  (Note that this array is different from
// the ZetasAVX2 array, which contains the elements of Zetas many times
// over in a way that is efficient for our implementation.)
//
// The vertical lines with crosses on them represent a swap.

// Related reading: https://eprint.iacr.org/2018/039.pdf

// Compute 4x16 Cooley--Tukey butterflies (a, b) ↦ (a + ζb, a - ζb).
//
// There is a catch: the first two and the last two sets of butterflies
// have to use the same sets of zetas, as we don't have enough registers
// to keep everything around.  t1 up to t4 are temporary registers that
// will be clobbered.

// In the generic implementation, a single butterfly is computed as
// follows (unfolding the definition of montReduce):
//
//  m := int16(zeta * b * 62209)
//  t := int16(uint32(zeta * int32(b) - m * int32(Q)) >> 16)
//  b = a - t
//  a += t
//
// As ζb ≡ mq (mod 2¹⁶), see comments on montReduce(), we can
// also compute t as
//
//  t := (uint32(zeta * int32(b)) >> 16) - (uint32(m * int32(Q)) >> 16)
//
// m (x16) can be computed using a single VPMULLW with zeta * 62209
// as the second operand stored in a table.  The two multiplications
// and bitshifts for t can be performed using two VPMULHWs (again
// for 16 at a time.)

// m = int16(zeta * b * 62209)

// uint32(zeta*int32(b)) >> 16

// uint32(m*int32(Q)) >> 16

// computes t

// b = a - t

// a = a + t

// First level:

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layers 2 - 7
//
// Layers 2 and 3 are straight forward.  From layers 4 onwards, the
// shuffling begins.

// Layer 2

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 3

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 4
//
// On this layer, the butterflies are of length 16 and so would still
// fit.  However, the first set of butterflies uses Zetas[8], whereas
// the second set uses Zetas[9].  On the next layer the butterflies
// are each of length 8 and wouldn't fit.  We solve both issues now
// by swapping the second part of the first set with the first part
// of the second set, etc.  This is a bitflip() with i=4.

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 5
//
// On this layer, the butterflies are of length 8 and wouldn't fit
// directly.  However, because of the previous shuffling we have
// sets of 16 consecutive butterflies.  However, just like in the
// previous layer, the ζs required by the two sets are different:
// the first sets uses 16 & 17 whereas the second uses 18 - 19.
// We solve the issue by swapping two pairs of quarters.  At the same
// time this swapping also ensures that the next layer has consecutive
// butterflies.  This is bitflip() with i=3.  There is some freedom
// which pairs to flip.  We try to keep the permutations as local
// as possible: there is only mixing between xs[0], xs[1], xs[2]
// and xs[3].  As an added benefit this ensures that the final
// complete permutation is convenient for multiplication.

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 6
//
// We continue on with the same principle.

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

// Layer 7

// a1, b1
// a2, b2
// zs12l, zs12h
// a3, b3,
// a4, b4,
// zs34l, zs34h
// t1, t2, t3, t4
// q

func mulHatAVX2() { _ = "STUB: not implemented"; return }

// In case BP register is needed.

// = q⁻¹ (mod 2¹⁶)

// Recall that quite conveniently for this computation (when j=0),
//
//  a[0] contains a₀, a₄, ..., a₆₀,
//  a[1] contains a₁, a₅, ..., a₆₁,
//  a[2] contains a₂, a₆, ..., a₆₂ and
//  a[3] contains a₃, a₇, ..., a₆₃
//
// and a similar thing for b, p and j>0.

// We have to compute several products of the form t=montReduce(a*b).
// From the discussion in AVX2-optimized NTT, recall that we can
// compute this as follows.
//
//  m := int16(a * b * 62209) = int16(a * b * -3227)
//  t := int16(uint32(int32(a) * int32(b) - m * int32(Q)) >> 16)
//     = (uint32(int32(a) * int32(b))>>16) - (uint32(m * int32(Q))>>16)

// We start with the first four lines of
//
//  p0 := montReduce(int32(a[i+1]) * int32(b[i+1]))
//  p2 := montReduce(int32(a[i]) * int32(b[i]))
//  p1 := montReduce(int32(a[i]) * int32(b[i+1]))
//  p1 += montReduce(int32(a[i+1]) * int32(b[i]))
//  p0 = montReduce(int32(p0) * zeta) + p2

// zl and zh are used as temporary registers here
// will end up in b[0]
// will end up in b[1]

// a[i+1]*b[i+1]
// a[i]*b[i]
// a[i]*b[i+1]
// a[i+1]*b[i]

// Compute p0 = montReduce(int32(p0) * zeta) + p2

// p1

// Now the same but then for the next two

// zl and zh are used as temporary registers here
// will end up in b[2]
// will end up in b[3]

// a[i+1]*b[i+1]
// a[i]*b[i]
// a[i]*b[i+1]
// a[i+1]*b[i]

// Compute p0 = p2 - montReduce(int32(p0) * zeta)

// p1

func tangleAVX2() { _ = "STUB: not implemented"; return }

func detangleAVX2() { _ = "STUB: not implemented"; return }

func barrettReduceAVX2() { _ = "STUB: not implemented"; return }

// Recall that the Barrett reduction of x is given by
//
//  x - int16((int32(x)*20159)>>26)*q

// t := (int32(x) * 20159) >> 16

// t = int16(t)>>10 so that t = (int32(x)*20159) >> 26

// t *= q

// x -= t

func normalizeAVX2() { _ = "STUB: not implemented"; return }

// Just like the generic implementation, we do a Barrett reduction
// followed by a conditional subtraction.

// Recall that the Barrett reduction of x is given by
//
//  x - int16((int32(x)*20159)>>26)*q

// t := (int32(x) * 20159) >> 16

// t = int16(t)>>10 so that t = (int32(x)*20159) >> 26

// t *= q

// x -= t

// x is now Barrett reduced.  Next we conditionally subtract q to
// normalize it.
//
//  x -= Q
//  x += (x >> 15) & Q

// x -= q

// t := x >> 15

// t &= q

// x += t

func main() {
	ConstraintExpr("amd64,!purego")

	addAVX2()
	subAVX2()
	nttAVX2()
	invNttAVX2()
	mulHatAVX2()
	detangleAVX2()
	tangleAVX2()
	barrettReduceAVX2()
	normalizeAVX2()

	Generate()
}
