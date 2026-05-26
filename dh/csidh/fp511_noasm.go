//go:build !amd64 || purego
// +build !amd64 purego

package csidh

func mul512(r, m1 *fp, m2 uint64)            { _ = "STUB: not implemented"; return }
func mul576(r *[9]uint64, m1 *fp, m2 uint64) { _ = "STUB: not implemented"; return }
func cswap512(x, y *fp, choice uint8)        { _ = "STUB: not implemented"; return }
func mulRdc(r, x, y *fp)                     { _ = "STUB: not implemented"; return }
