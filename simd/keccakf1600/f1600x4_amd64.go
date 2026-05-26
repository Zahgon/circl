//go:build amd64 && !purego
// +build amd64,!purego

package keccakf1600

func permuteSIMDx4(state []uint64, turbo bool) { _ = "STUB: not implemented"; return }

func permuteSIMDx2(state []uint64, turbo bool) { _ = "STUB: not implemented"; return }
