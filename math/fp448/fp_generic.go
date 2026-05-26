package fp448

func cmovGeneric(x, y *Elt, n uint) { _ = "STUB: not implemented"; return }

func cswapGeneric(x, y *Elt, n uint) { _ = "STUB: not implemented"; return }

func addGeneric(z, x, y *Elt) { _ = "STUB: not implemented"; return }

func subGeneric(z, x, y *Elt) { _ = "STUB: not implemented"; return }

func addsubGeneric(x, y *Elt) { _ = "STUB: not implemented"; return }

func mulGeneric(z, x, y *Elt) { _ = "STUB: not implemented"; return }

func sqrGeneric(z, x *Elt) { _ = "STUB: not implemented"; return }

func red64(z *Elt, l, h *[7]uint64) {
	_ = "STUB: not implemented"
	/* (2C13, 2C12, 2C11, 2C10|C10, C9, C8, C7) + (C6,...,C0) */ return
}

/* (C10C9, C9C8,C8C7,C7C13,C13C12,C12C11,C11C10) + (C6,...,C0) */

/* (C7) + (C6,...,C0) */

/* (C7) + (C6,...,C0) */
