//go:build amd64 && !purego
// +build amd64,!purego

package csidh

import (
	"golang.org/x/sys/cpu"
)

var (
	// Signals support for BMI2 (MULX)
	hasBMI2 = cpu.X86.HasBMI2
	// Signals support for ADX and BMI2
	hasADXandBMI2 = cpu.X86.HasBMI2 && cpu.X86.HasADX
)

var _ = hasBMI2

func mul512(r, m1 *fp, m2 uint64)     { _ = "STUB: not implemented"; return }
func cswap512(x, y *fp, choice uint8) { _ = "STUB: not implemented"; return }
func mulRdc(r, x, y *fp)              { _ = "STUB: not implemented"; return }

//go:noescape
func mul512Amd64(a, b *fp, c uint64)

//go:noescape
func cswap512Amd64(x, y *fp, choice uint8)

//go:noescape
func mulBmiAsm(res, x, y *fp)

// mulRdc performs montgomery multiplication r = x * y mod P.
// Returned result r is already reduced and in Montgomery domain.
func mulRdcAmd64(r, x, y *fp) { _ = "STUB: not implemented"; return }

// if p <= r < 2p then r = r-p
