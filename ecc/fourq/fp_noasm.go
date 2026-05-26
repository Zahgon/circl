//go:build !amd64 || purego
// +build !amd64 purego

package fourq

func fpMod(c *Fp)       { _ = "STUB: not implemented"; return }
func fpAdd(c, a, b *Fp) { _ = "STUB: not implemented"; return }
func fpSub(c, a, b *Fp) { _ = "STUB: not implemented"; return }
func fpMul(c, a, b *Fp) { _ = "STUB: not implemented"; return }
func fpSqr(c, a *Fp)    { _ = "STUB: not implemented"; return }
func fpHlf(c, a *Fp)    { _ = "STUB: not implemented"; return }
