//go:build !amd64 || purego
// +build !amd64 purego

package fourq

func (P *pointR1) double()           { _ = "STUB: not implemented"; return }
func (P *pointR1) add(Q *pointR2)    { _ = "STUB: not implemented"; return }
func (P *pointR1) mixAdd(Q *pointR3) { _ = "STUB: not implemented"; return }
