package common

// Constant time select.
// if pick == 1 (out = in1)
// if pick == 0 (out = in2)
// else out is undefined.
func Cpick(pick int, out, in1, in2 []byte) { _ = "STUB: not implemented"; return }

// Read 2*bytelen(p) bytes into the given ExtensionFieldElement.
//
// It is an error to call this function if the input byte slice is less than 2*bytelen(p) bytes long.
func BytesToFp2(fp2 *Fp2, input []byte, bytelen int) { _ = "STUB: not implemented"; return }

// Convert the input to wire format.
//
// The output byte slice must be at least 2*bytelen(p) bytes long.
func Fp2ToBytes(output []byte, fp2 *Fp2, bytelen int) { _ = "STUB: not implemented"; return }
