//go:generate go run src.go -out ../amd64.s -stubs ../stubs_amd64.go -pkg dilithium

// AVX2 optimized version of Poly.[Inv]NTT().  See the comments on the generic
// implementation for details on the maths involved.
package main

import (
	. "github.com/mmcloughlin/avo/build"   // nolint:golint,stylecheck
	. "github.com/mmcloughlin/avo/operand" // nolint:golint,stylecheck
	. "github.com/mmcloughlin/avo/reg"     // nolint:golint,stylecheck
)

// XXX align Poly on 16 bytes such that we can use aligned moves
// XXX ensure Zetas and InvZetas are 16 byte aligned.

func broadcastImm32(c uint32, out Op) { _ = "STUB: not implemented"; return }

// Performs AND with an 64b immediate.
func andImm64(c uint64, inout Op) { _ = "STUB: not implemented"; return }

// Executes the permutation (a[2] b[0]) (a[3] b[1]) when considering only the
// even positions of a and b seen as [8]uint32.
func swapInner(a, b Op) { _ = "STUB: not implemented"; return }

// 0 + 2*16
// 1 + 3*16

// Executes the permutation (a[1] b[0]) (a[3] b[2]) when considering only the
// even positions of a and b seen as [8]uint32.
func oddCrossing(a, b Op) { _ = "STUB: not implemented"; return }

// nolint:funlen
func nttAVX2() {
	_ = "STUB: not implemented"
	// We perform the same operations as the generic implementation of NTT,
	// but use AVX2 to perform 16 butterflies at the same time.  For the
	// first few levels this is straight forward.  For the final levels we
	// need to move some coefficients around to be able to use the AVX2
	// instructions.
	return
}

// We allocate a [256]uint64 on the stack aligned to 32 bytes to hold
// "buf" which contains intermediate coefficients like "p" in the generic
// algorithm, but then in uint64s instead of uint32s.

// +32 to be able to align

// 4236238847 = -(q^-1) mod 2³²

// Computes 4x4 Cooley--Tukey butterflies (a,b) ↦ (a + ζb, a - ζb).

// Set b = bζ.

// Now we reduce b below 2Q with the method of reduceLe2Q():
//
//      t := ((b * 4236238847) & 0xffffffff) * uint64(Q)
//      return uint32((b + t) >> 32)

// t = b * 4236238847.

// t = (t & 0xffffffff) * Q.  The and is implicit as VPMULUDQ
// is a parallel 32b x 32b -> 64b multiplication.

// t = b + t

// t = t >> 32

// b = a + 2Q

// a += t

// b = b - t

// With AVX2 we can compute 4*4 Cooley--Tukey butterflies at the same time.
// As loading and storing from memory is expensive, we try to compute
// as much at the same time.

// First, second and third level.
// The first butterfly at the third level is (0, 32).  To compute it, we
// need to compute some butterflies on the second level and in turn
// the butterflies (0, 128), (32, 160), (64, 192) and (96, 224) on the
// first level.  As we need to compute them anyway, we compute the
// butterflies (0, 32), (64, 96), (128, 160) and (192, 224) on the
// third level at the same time.  Using the uint64x4 AVX2 registers,
// we compute (0, 32), (1, 33), ..., (4, 36), (64, 96), (64, 97), ...
// in one go.  This is one eighth of the third level.  We repeat another
// seven times with a shifted offset to compute the third level.

// XXX should we really unroll this loop?

// First level.
// Load the coefficients.  First uint32s of xs[0], xs[1], ...
// contains p[0], p[32], p[64], ..., p[224].

// Loads 4 32b coefficients at the same time; zeropads them to 64b
// and puts them in xs[i].

// XXX At the moment we've completely unrolled, so we could, if we want,
//     hardcode the Zetas here instead of looking them up from memory.
//     Is that worth it?

// Zetas[1]

// Second level
// Zetas[2]
// Zetas[3]

// Third level
// Zetas[4]
// Zetas[5]
// Zetas[6]
// Zetas[7]

// Fourth, fifth, sixth, seventh and eighth level.
// If we want to compute the butterfly (0, 1) in the eighth level, we need
// to compute the first 2 butterflies in the seventh level; the first 4
// of the sixth, ... and the first 16 in the fourth level which needs the
// first 32 coefficients already computed in the third level.
// Going forward again, we see that we can use these to compute the first
// 32 coefficients.  As each level requires 16 butterflies, we can
// conveniently perform these all in our YMM registers.
// After that we repeat the same method for the next 32 coefficients and
// continue for a total of eight times to finish the computation of
// the NTT.

// XXX should we really unroll this loop?

// Load the first 32 coefficients from level 3.  Recall that bufPtr
// has 64 bits of space for each coefficient.

// Fourth level

// Fifth level

// Sixth level

// Seventh level
// Now things get a bit trickier.  We have to compute the butterflies
// (0, 2), (1, 3), (4, 6), (5, 7), etc which don't fit our ctButterfly()
// routine, which likes to have four consecutive butterflies.
// To work around this, we swap 2 with 4 and 3 with 5, etc., which
// allows us to use our old routine.

// XXX optimize?  We might want to add a small extra table for just
//     these zetas so that we don't have to blend them.

// Eighth level
// Finally, we have to perform the butterflies (0, 1), (2, 3), etc.
// Swapping 1 with 4 and 3 with 6 (etc.) will ensure that a
// straight-forward call to our ctButterfly() routine will do the right
// thing.

// Packing.
// Due to swapInner() and oddCrossing() our coefficients are laid out
// as 0, 2, 4, 6, 1, 3, 5, 7, 8, 10, ... in xs[0], xs[1], ...
// with junk 32b in between.  By shifting the odd xss 32b to the
// left and merging them with the even xss, we get the desired
// order 0, 1, 2, 3, ... without any padding, which can then be
// moved out into memory.

// nolint:funlen
func invNttAVX2() {
	_ = "STUB: not implemented"
	// Just like with the generic implementation, we do the operations of
	// NTT in reverse, except for two things: we hoist out all divisions by
	// two from the Gentleman-Sande butterflies and accumulate them to one
	// big division by 2⁸ at the end.
	return
}

// We allocate a [256]uint64 on the stack aligned to 32 bytes to hold
// "buf" which contains intermediate coefficients like "p" in the generic
// algorithm, but then in uint64s instead of uint32s.

// +32 to be able to align

// Computes 4x4 doubled Gentleman--Sande butterflies (a,b) ↦ (a+b, ζ(a-b)).

// XXX be more parallel when we have more registers available, when
//     we don't use the full four registers for zetas.

// Set t = 256Q + a in preparation of subtracting b

// Set t = t - b

// Set a = a + b

// Set b = tζ

// Now we reduce b below 2Q with the method of reduceLe2Q():
//
//      t := ((b * 4236238847) & 0xffffffff) * uint64(Q)
//      return uint32((b + t) >> 32)

// t = b * 4236238847.

// t = (t & 0xffffffff) * Q.  The and is implicit as VPMULUDQ
// is a parallel 32b x 32b -> 64b multiplication.

// t = b + t

// b = t >> 32

// XXX should we really unroll this loop?

// Load coeffs 0 1 2 3 4 5 6 7 into xs[0], 8 ... 16 into xs[1], etc.

// Move odd coeffs of xs[2*i] into xs[2*i+1] and shift down.  Ignoring
// the odd coefficients, we have 0 2 4 6 in xs[0] and 1 3 4 5 in xs[1].

// Eighth level

// See comments in nttAVX2() above about oddCrossing() and swapInner().

// Seventh level

// XXX optimize?  We might want to add a small extra table for just
//     these zetas so that we don't have to blend them.

// See comments in nttAVX2() above about oddCrossing() and swapInner()

// Sixth level

// Fifth level

// Fourth level

// XXX should we really unroll this loop?

// Third level

// Second level

// First level

// Finally, we multiply by 41978 = (256)^-1 R² ...

// (we need this loop, otherwise we run out of YMM registers.)

// ... and reduce below 2Q with the method of reduceLe2Q():
//
//      t := ((x * 4236238847) & 0xffffffff) * uint64(Q)
//      return uint32((x + t) >> 32)

// t = x * 4236238847.

// t = (t & 0xffffffff) * Q.  The and is implicit as VPMULUDQ
// is a parallel 32b x 32b -> 64b multiplication.

// t = x + t

// x = t >> 32

// Finally, we copy the 32b results from the [256]uint64 buf to
// the [256]uint32 p.
// XXX is this the most efficient way?

// Recall that oddCrossing after swapInner will permute the
// even coefficients from 0 1 2 3 4 5 6 7 to 0 2 4 6 1 3 5 7 and so
// then we can simply shift and blend the last four into the first four
// as we did at the end of nttAVX2().

// XXX Split out into separate file.  To do this we need to figure out how
//
//	to share code properly between avo modules.
func mulHatAVX2() { _ = "STUB: not implemented"; return }

// XXX Is this loop unrolling worthwhile?

// XXX We could use 6 registers each (instead of 4).  Does that make
//     it faster?

// Now we reduce b below 2Q with the method of reduceLe2Q():
//
//      a := ((b * 4236238847) & 0xffffffff) * uint64(Q)
//      return uint32((b + a) >> 32)

// a = b * 4236238847.

// t = (t & 0xffffffff) * Q.  The and is implicit as VPMULUDQ
// is a parallel 32b x 32b -> 64b multiplication.

// t = b + a

// b = a >> 32

// Pack into p.  See end of invNttAvx2() for a description of the method.
// XXX is there a better way to do this that avoids the PERM
//     in oddCrossing?

func addAVX2() { _ = "STUB: not implemented"; return }

// XXX is unrolling worth it?

func subAVX2() { _ = "STUB: not implemented"; return }

// XXX is unrolling worth it?

func packLe16AVX2() { _ = "STUB: not implemented"; return }

// We load p[0], ..., p[7] into a[0], p[8], ..., p[15] into a[1], etc.,
// so we may consider a as a matrix.  We transpose a in the usual way.

// a has been transposed, so a[0] contains p[0], p[8], ... and
// a[1] contains p[1], p[9], ..., etc.  We shift a[i] by 4*i to the left
// and or them together.

func reduceLe2QAVX2() { _ = "STUB: not implemented"; return }

// We use the same computation as used in reduceLe2Q() for the separate
// coefficients.

// b = a >> 23

// a = a & 2²³-1

// c = (b << 13) - b

// a = a + c

// Write back

func le2qModQAVX2() { _ = "STUB: not implemented"; return }

// We use the same method as le2qModQ().

// a -= Q

// m = uint32(int32(a) >> 31)

// m &= q

// a += m

func exceedsAVX2() { _ = "STUB: not implemented"; return }

// We use the same method as Poly.exceedsGeneric().

// a = (Q-1)/2 - a

// b = a >> 31

// a = a ^ b

// a = (Q-1)/2 - a

// Here exceedsGeneric() checks if a ⩾ bound.  We'll be more clever.
// a ⩾ bound iff a - bound ⩾ 0, so set a = a - bound first.

// a &= 0x80000000.  Leaves the sign.  Should be zero.

// Move the high bits, which are all zero except possibly for
// the sign bits, into tmp.

// If one of the sign bits is zero, then one of the as is
// positive hence the bound is exceeded.
// 0b10001000100010001000100010001000

func mulBy2toDAVX2() { _ = "STUB: not implemented"; return }

func main() {
	ConstraintExpr("amd64,!purego")

	nttAVX2()
	invNttAVX2()
	mulHatAVX2()
	addAVX2()
	subAVX2()
	packLe16AVX2()
	reduceLe2QAVX2()
	le2qModQAVX2()
	exceedsAVX2()
	mulBy2toDAVX2()

	Generate()
}
