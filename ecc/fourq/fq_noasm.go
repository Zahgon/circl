//go:build !amd64 || purego
// +build !amd64 purego

package fourq

func fqCmov(c, a *Fq, b int) { _ = "STUB: not implemented"; return }

func fqAdd(c, a, b *Fq) { _ = "STUB: not implemented"; return }
func fqSub(c, a, b *Fq) { _ = "STUB: not implemented"; return }
func fqMul(c, a, b *Fq) { _ = "STUB: not implemented"; return }
func fqSqr(c, a *Fq)    { _ = "STUB: not implemented"; return }
