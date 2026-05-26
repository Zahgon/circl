// Code generated from ./templates/poly.go.tmpl. DO NOT EDIT.

package fp64

type Poly []Fp

func (p Poly) AddAssign(x Poly) { _ = "STUB: not implemented"; return }
func (p Poly) SubAssign(x Poly) { _ = "STUB: not implemented"; return }
func (p Poly) Mul(x, y Poly)    { _ = "STUB: not implemented"; return }

func (p Poly) MulNSquare(x, y Poly) { _ = "STUB: not implemented"; return }

func (p Poly) MulNlogN(x, y Poly) { _ = "STUB: not implemented"; return }

func (p Poly) Sqr(x Poly) { _ = "STUB: not implemented"; return }

func (p Poly) Evaluate(x *Fp) (px Fp) { _ = "STUB: not implemented"; return *new(Fp) }

func (p Poly) Strip() Poly { _ = "STUB: not implemented"; return *new(Poly) }
