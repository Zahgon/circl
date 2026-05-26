//go:build !amd64 || purego
// +build !amd64 purego

package x448

import fp "github.com/cloudflare/circl/math/fp448"

func double(x, z *fp.Elt)             { _ = "STUB: not implemented"; return }
func diffAdd(w *[5]fp.Elt, b uint)    { _ = "STUB: not implemented"; return }
func ladderStep(w *[5]fp.Elt, b uint) { _ = "STUB: not implemented"; return }
func mulA24(z, x *fp.Elt)             { _ = "STUB: not implemented"; return }
