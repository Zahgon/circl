package fp25519

func cmovGeneric(x, y *Elt, n uint) { _ = "STUB: not implemented"; return }

func cswapGeneric(x, y *Elt, n uint) { _ = "STUB: not implemented"; return }

func addGeneric(z, x, y *Elt) { _ = "STUB: not implemented"; return }

func subGeneric(z, x, y *Elt) { _ = "STUB: not implemented"; return }

func addsubGeneric(x, y *Elt) { _ = "STUB: not implemented"; return }

func mulGeneric(z, x, y *Elt) { _ = "STUB: not implemented"; return }

func sqrGeneric(z, x *Elt) { _ = "STUB: not implemented"; return }

func modpGeneric(x *Elt) { _ = "STUB: not implemented"; return }

// CX = C[255] ? 38 : 19

// PUT BIT 255 IN CARRY FLAG AND CLEAR

// TEST FOR BIT 255 AGAIN; ONLY TRIGGERED ON OVERFLOW MODULO 2^255-19
// cx = C[255] ? 0 : 19

// CLEAR BIT 255

func red64(z *Elt, x0, x1, x2, x3, x4, x5, x6, x7 uint64) { _ = "STUB: not implemented"; return }
