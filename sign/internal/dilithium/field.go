package dilithium

// Returns a y with y < 2q and y = x mod q.
// Note that in general *not*: ReduceLe2Q(ReduceLe2Q(x)) == x.
func ReduceLe2Q(x uint32) uint32 {
	_ = "STUB: not implemented"
	// Note 2²³ = 2¹³ - 1 mod q. So, writing  x = x₁ 2²³ + x₂ with x₂ < 2²³
	// and x₁ < 2⁹, we have x = y (mod q) where
	// y = x₂ + x₁ 2¹³ - x₁ ≤ 2²³ + 2¹³ < 2q.
	return 0
}

// 2²³-1

// Returns x mod q.
func modQ(x uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// For x R ≤ q 2³², find y ≤ 2q with y = x mod q.
func montReduceLe2Q(x uint64) uint32 {
	_ = "STUB: not implemented"
	// Qinv = 4236238847 = -(q⁻¹) mod 2³²
	return 0
}

// Returns x mod q for 0 ≤ x < 2q.
func le2qModQ(x uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// mask is 2³²-1 if x was neg.; 0 otherwise

// Splits 0 ≤ a < Q into a0 and a1 with a = a1*2ᴰ + a0
// and -2ᴰ⁻¹ < a0 < 2ᴰ⁻¹.  Returns a0 + Q and a1.
func power2round(a uint32) (a0plusQ, a1 uint32) {
	_ = "STUB: not implemented"
	// We effectively compute a0 = a mod± 2ᵈ
	//
	//	and a1 = (a - a0) / 2ᵈ.
	return 0, 0
}

// a mod 2ᵈ

// a0 is one of  0, 1, ..., 2ᵈ⁻¹-1, 2ᵈ⁻¹, 2ᵈ⁻¹+1, ..., 2ᵈ-1

// now a0 is     -2ᵈ⁻¹-1, -2ᵈ⁻¹, ..., -2, -1, 0, ..., 2ᵈ⁻¹-2
// Next, we add 2ᴰ to those a0 that are negative (seen as int32).

// now a0 is     2ᵈ⁻¹-1, 2ᵈ⁻¹, ..., 2ᵈ-2, 2ᵈ-1, 0, ..., 2ᵈ⁻¹-2

// now a0 id     0, 1, 2, ..., 2ᵈ⁻¹-1, 2ᵈ⁻¹-1, -2ᵈ⁻¹-1, ...
// which is what we want.
