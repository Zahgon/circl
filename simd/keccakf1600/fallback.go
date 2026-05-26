//go:build (!amd64 && !arm64) || (arm64 && !go1.16) || purego
// +build !amd64,!arm64 arm64,!go1.16 purego

package keccakf1600

func permuteSIMDx2(state []uint64, turbo bool) { _ = "STUB: not implemented"; return }

func permuteSIMDx4(state []uint64, turbo bool) { _ = "STUB: not implemented"; return }
