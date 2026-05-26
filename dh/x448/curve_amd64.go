//go:build amd64 && !purego
// +build amd64,!purego

package x448

import (
	fp "github.com/cloudflare/circl/math/fp448"
	"golang.org/x/sys/cpu"
)

var hasBmi2Adx = cpu.X86.HasBMI2 && cpu.X86.HasADX

var _ = hasBmi2Adx

func double(x, z *fp.Elt)             { _ = "STUB: not implemented"; return }
func diffAdd(w *[5]fp.Elt, b uint)    { _ = "STUB: not implemented"; return }
func ladderStep(w *[5]fp.Elt, b uint) { _ = "STUB: not implemented"; return }
func mulA24(z, x *fp.Elt)             { _ = "STUB: not implemented"; return }

//go:noescape
func doubleAmd64(x, z *fp.Elt)

//go:noescape
func diffAddAmd64(w *[5]fp.Elt, b uint)

//go:noescape
func ladderStepAmd64(w *[5]fp.Elt, b uint)

//go:noescape
func mulA24Amd64(z, x *fp.Elt)
