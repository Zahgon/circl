package ed25519

var order = [paramB]byte{
	0xed, 0xd3, 0xf5, 0x5c, 0x1a, 0x63, 0x12, 0x58,
	0xd6, 0x9c, 0xf7, 0xa2, 0xde, 0xf9, 0xde, 0x14,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10,
}

// isLessThan returns true if 0 <= x < y, and assumes that slices have the same length.
func isLessThan(x, y []byte) bool { _ = "STUB: not implemented"; return false }

// reduceModOrder calculates k = k mod order of the curve.
func reduceModOrder(k []byte, is512Bit bool) { _ = "STUB: not implemented"; return }

// red512 calculates x = x mod Order of the curve.
func red512(x *[8]uint64, full bool) {
	_ = "STUB: not implemented"
	// Implementation of Algs.(14.47)+(14.52) of Handbook of Applied
	// Cryptography, by A. Menezes, P. van Oorschot, and S. Vanstone.
	return
}

// if q=0 then m=0...0 else m=1..1

// calculateS performs s = r+k*a mod Order of the curve.
func calculateS(s, r, k, a []byte) { _ = "STUB: not implemented"; return }
