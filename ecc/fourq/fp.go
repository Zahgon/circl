package fourq

import (
	"math/big"
)

var modulusP = Fp{
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f,
}

// SizeFp is the length in bytes to represent an element in the base field.
const SizeFp = 16

// Fp is an element (in littleEndian order) of prime field GF(2^127-1).
type Fp [SizeFp]byte

func (f *Fp) String() string       { _ = "STUB: not implemented"; return "" }
func (f *Fp) isZero() bool         { _ = "STUB: not implemented"; return false }
func (f *Fp) toBigInt() *big.Int   { _ = "STUB: not implemented"; return nil }
func (f *Fp) setBigInt(b *big.Int) { _ = "STUB: not implemented"; return }
func (f *Fp) toBytes(buf []byte)   { _ = "STUB: not implemented"; return }

func (f *Fp) fromBytes(buf []byte) bool { _ = "STUB: not implemented"; return false }

func fpNeg(c, a *Fp) { _ = "STUB: not implemented"; return }

// fqSgn returns the sign of an element.
//
//	-1 if x >  (p+1)/2
//	 0 if x == 0
//	+1 if x >  (p+1)/2.
func fpSgn(c *Fp) int { _ = "STUB: not implemented"; return 0 }

// fpTwo1251 sets c = a^(2^125-1).
func fpTwo1251(c, a *Fp) { _ = "STUB: not implemented"; return }

// fpInv sets z to a^(-1) mod p.
func fpInv(z, a *Fp) { _ = "STUB: not implemented"; return }
