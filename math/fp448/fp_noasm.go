//go:build !amd64 || purego
// +build !amd64 purego

package fp448

func cmov(x, y *Elt, n uint)  { _ = "STUB: not implemented"; return }
func cswap(x, y *Elt, n uint) { _ = "STUB: not implemented"; return }
func add(z, x, y *Elt)        { _ = "STUB: not implemented"; return }
func sub(z, x, y *Elt)        { _ = "STUB: not implemented"; return }
func addsub(x, y *Elt)        { _ = "STUB: not implemented"; return }
func mul(z, x, y *Elt)        { _ = "STUB: not implemented"; return }
func sqr(z, x *Elt)           { _ = "STUB: not implemented"; return }
