package frodo640shake

func add(out *nbarByNbarU16, lhs *nbarByNbarU16, rhs *nbarByNbarU16) {
	_ = "STUB: not implemented"
	return
}

func sub(out *nbarByNbarU16, lhs *nbarByNbarU16, rhs *nbarByNbarU16) {
	_ = "STUB: not implemented"
	return
}

func pack(out []byte, in []uint16) { _ = "STUB: not implemented"; return }

func unpack(out []uint16, in []byte) { _ = "STUB: not implemented"; return }

func encodeMessage(out *nbarByNbarU16, msg *[messageSize]byte) { _ = "STUB: not implemented"; return }

// 16 = bit size of out[i]

func decodeMessage(out *[messageSize]byte, msg *nbarByNbarU16) { _ = "STUB: not implemented"; return }

func mulAddSBPlusE(out *nbarByNbarU16, s []uint16, b *nByNbarU16, e []uint16) {
	_ = "STUB: not implemented"
	// Multiply by s on the left
	// Inputs: b (N x N_BAR), s (N_BAR x N), e (N_BAR x N_BAR)
	// Output: out = s*b + e (N_BAR x N_BAR)
	return
}

func mulBS(out *nbarByNbarU16, b *nbarByNU16, s *nByNbarU16) { _ = "STUB: not implemented"; return }

func ctCompareU16(lhs []uint16, rhs []uint16) int {
	_ = "STUB: not implemented"
	// Compare lhs and rhs in constant time.
	// Returns 0 if they are equal, 1 otherwise.
	return 0
}
