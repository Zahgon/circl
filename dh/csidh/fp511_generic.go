package csidh

// mul576 implements schoolbook multiplication of
// 64x512-bit integer. Returns result modulo 2^512.
// r = m1*m2.
func mul512Generic(r, m1 *fp, m2 uint64) { _ = "STUB: not implemented"; return }

// mul576 implements schoolbook multiplication of
// 64x512-bit integer. Returns 576-bit result of
// multiplication.
// r = m1*m2.
func mul576Generic(r *[9]uint64, m1 *fp, m2 uint64) { _ = "STUB: not implemented"; return }

// cswap512 implements constant time swap operation.
// If choice = 0, leave x,y unchanged. If choice = 1, set x,y = y,x.
// If choice is neither 0 nor 1 then behaviour is undefined.
func cswap512Generic(x, y *fp, choice uint8) { _ = "STUB: not implemented"; return }

// mulRdc performs montgomery multiplication r = x * y mod P.
// Returned result r is already reduced and in Montgomery domain.
func mulRdcGeneric(r, x, y *fp) { _ = "STUB: not implemented"; return }

// if p <= r < 2p then r = r-p

func mulGeneric(r, x, y *fp) {
	_ = "STUB: not implemented"
	// keeps intermediate results
	return
}

// x[i]*y + q_i*p

// s = (s + x[i]*y + q_i * p) / R

// last iteration stores result in r
