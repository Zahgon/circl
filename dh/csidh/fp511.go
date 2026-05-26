package csidh

// Constant time select.
// if pick == 0xFF..FF (out = in1)
// if pick == 0 (out = in2)
// else out is undefined.
func ctPick64(which uint64, in1, in2 uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// ctIsNonZero64 returns 0 in case i == 0, otherwise it returns 1.
// Constant-time.
func ctIsNonZero64(i uint64) int {
	_ = "STUB: not implemented"
	// In case i==0 then i-1 will set MSB. Only in such case (i OR ~(i-1))
	// will result in MSB being not set (logical implication: (i-1)=>i is
	// false iff (i-1)==0 and i==non-zero). In every other case MSB is
	// set and hence function returns 1.
	return 0
}

// Returns result of x<y operation.
func isLess(x, y *fp) bool { _ = "STUB: not implemented"; return false }

// x == y

// r = x + y mod p.
func addRdc(r, x, y *fp) { _ = "STUB: not implemented"; return }

// r = x - y.
func sub512(r, x, y *fp) uint64 { _ = "STUB: not implemented"; return 0 }

// r = x - y mod p.
func subRdc(r, x, y *fp) {
	_ = "STUB: not implemented"

	// Same as sub512(r,x,y). Unfortunately
	// compiler is not able to inline it.
	return
}

// if x<y => r=x-y+p

// Fixed-window mod exp for fpBitLen bit value with 4 bit window. Returned
// result is a number in montgomery domain.
// r = b ^ e (mod p).
// Constant time.
func modExpRdcCommon(r, b, e *fp, fpBitLen int) { _ = "STUB: not implemented"; return }

// Precompute step, computes an array of small powers of 'b'. As this
// algorithm implements 4-bit window, we need 2^4=16 of such values.
// b^0 = 1, which is equal to R from REDC.
// b ^ 0
// b ^ 1

// OPTIMIZE: implement fast squaring. Then interleaving fast squaring
// with multiplication should improve performance.
// sqr

// note: non resistant to cache SCA

// if p <= r < 2p then r = r-p

// modExpRdc does modular exponentiation of 512-bit number.
// Constant-time.
func modExpRdc512(r, b, e *fp) { _ = "STUB: not implemented"; return }

// modExpRdc does modular exponentiation of 64-bit number.
// Constant-time.
func modExpRdc64(r, b *fp, e uint64) { _ = "STUB: not implemented"; return }

// isNonQuadRes checks whether value v is quadratic residue.
// Implementation uses Fermat's little theorem (or
// Euler's criterion)
//
//	a^(p-1) == 1, hence
//	(a^2) ((p-1)/2) == 1
//
// Which means v is a quadratic residue iff v^((p-1)/2) == 1.
// Caller provided v must be in montgomery domain.
// Returns 0 in case v is quadratic residue or 1 in case
// v is quadratic non-residue.
func (v *fp) isNonQuadRes() int { _ = "STUB: not implemented"; return 0 }

// isZero returns false in case v is equal to 0, otherwise
// true. Constant time.
func (v *fp) isZero() bool { _ = "STUB: not implemented"; return false }

// equal checks if v is equal to in. Constant time.
func (v *fp) equal(in *fp) bool { _ = "STUB: not implemented"; return false }
