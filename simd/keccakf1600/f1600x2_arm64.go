//go:build arm64 && go1.16 && !purego
// +build arm64,go1.16,!purego

package keccakf1600

func permuteSIMDx2(state []uint64, turbo bool) { _ = "STUB: not implemented"; return }

func permuteSIMDx4(state []uint64, turbo bool) { _ = "STUB: not implemented"; return }

//go:noescape
func f1600x2ARM(state *uint64, rc *[24]uint64, turbo bool)
