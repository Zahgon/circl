//go:build amd64 && !purego
// +build amd64,!purego

package fp448

import (
	"golang.org/x/sys/cpu"
)

var hasBmi2Adx = cpu.X86.HasBMI2 && cpu.X86.HasADX

var _ = hasBmi2Adx

func cmov(x, y *Elt, n uint)  { _ = "STUB: not implemented"; return }
func cswap(x, y *Elt, n uint) { _ = "STUB: not implemented"; return }
func add(z, x, y *Elt)        { _ = "STUB: not implemented"; return }
func sub(z, x, y *Elt)        { _ = "STUB: not implemented"; return }
func addsub(x, y *Elt)        { _ = "STUB: not implemented"; return }
func mul(z, x, y *Elt)        { _ = "STUB: not implemented"; return }
func sqr(z, x *Elt) {
	_ = "STUB: not implemented"
	return

	/* Functions defined in fp_amd64.s */
}

//go:noescape
func cmovAmd64(x, y *Elt, n uint)

//go:noescape
func cswapAmd64(x, y *Elt, n uint)

//go:noescape
func addAmd64(z, x, y *Elt)

//go:noescape
func subAmd64(z, x, y *Elt)

//go:noescape
func addsubAmd64(x, y *Elt)

//go:noescape
func mulAmd64(z, x, y *Elt)

//go:noescape
func sqrAmd64(z, x *Elt)
