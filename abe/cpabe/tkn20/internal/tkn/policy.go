package tkn

import (
	pairing "github.com/cloudflare/circl/ecc/bls12381"
)

const (
	bkAttribute   = "internal-boneh-katz-transform-attribute"
	attributeSize = pairing.ScalarSize + 1
)

type Wire struct {
	Label    string
	RawValue string
	Value    *pairing.Scalar
	Positive bool
}

func (w *Wire) String() string { _ = "STUB: not implemented"; return "" }

type Policy struct {
	Inputs []Wire
	F      Formula // monotonic boolean formula
}

type Attribute struct {
	wild  bool // false if tame
	Value *pairing.Scalar
}

func (a *Attribute) marshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *Attribute) unmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (a *Attribute) Equal(b *Attribute) bool { _ = "STUB: not implemented"; return false }

type Attributes map[string]Attribute

func (a *Attributes) marshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *Attributes) unmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (a *Attributes) Equal(b *Attributes) bool { _ = "STUB: not implemented"; return false }

func (w *Wire) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Wire) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (w *Wire) Equal(w2 *Wire) bool { _ = "STUB: not implemented"; return false }

func (p *Policy) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Policy) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	// Extract formula
	return nil
}

// Extract wires

func (p *Policy) Equal(p2 *Policy) bool { _ = "STUB: not implemented"; return false }

func (p *Policy) String() string {
	_ = "STUB: not implemented"
	// gateAssign takes n wires (intermediates and outputs) and maps to the gate
	// that set them. For details, refer to [Formula].
	return ""
}

func (p *Policy) printWire(gateAssign []int, wire int) string { _ = "STUB: not implemented"; return "" }

type match struct {
	wire  int
	label string
}

type Satisfaction struct {
	matches []match
}

func (p *Policy) pi() []int { _ = "STUB: not implemented"; return nil }

// Paper would have us put a +1 here
// we change the indexing instead

func (p *Policy) Satisfaction(attr *Attributes) (*Satisfaction, error) {
	_ = "STUB: not implemented"
	// For now its all of the wires, so we don't need to look at the formula.
	return nil, nil
}

// missing Attribute might not be needed

// Carry Out the augmentation under the BK transform
func (p *Policy) transformBK(val *pairing.Scalar) *Policy { _ = "STUB: not implemented"; return nil }

func transformAttrsBK(attr *Attributes) *Attributes { _ = "STUB: not implemented"; return nil }
