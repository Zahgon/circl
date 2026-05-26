package slhdsa

// See FIPS 205 -- Section 4.2
// Functions and Addressing

type addrType = uint32

const (
	addressWotsHash = addrType(iota)
	addressWotsPk
	addressTree
	addressForsTree
	addressForsRoots
	addressWotsPrf
	addressForsPrf
)

const (
	addressSizeCompressed    = 22
	addressSizeNonCompressed = 32
)

type address struct {
	b []byte
	o int
}

func (p *params) addressSize() uint32 { _ = "STUB: not implemented"; return 0 }

func (p *params) addressOffset() int { _ = "STUB: not implemented"; return 0 }

func (p *params) NewAddress() (a address) { _ = "STUB: not implemented"; return *new(address) }

func (a *address) fromBytes(p *params, c *cursor) { _ = "STUB: not implemented"; return }

func (a *address) Set(x address)              { _ = "STUB: not implemented"; return }
func (a *address) Clear()                     { _ = "STUB: not implemented"; return }
func (a *address) SetKeyPairAddress(i uint32) { _ = "STUB: not implemented"; return }
func (a *address) SetChainAddress(i uint32)   { _ = "STUB: not implemented"; return }
func (a *address) SetTreeHeight(i uint32)     { _ = "STUB: not implemented"; return }
func (a *address) SetHashAddress(i uint32)    { _ = "STUB: not implemented"; return }
func (a *address) SetTreeIndex(i uint32)      { _ = "STUB: not implemented"; return }
func (a *address) GetKeyPairAddress() uint32  { _ = "STUB: not implemented"; return 0 }
func (a *address) SetLayerAddress(l addrType) { _ = "STUB: not implemented"; return }

func (a *address) SetTreeAddress(t [3]uint32) { _ = "STUB: not implemented"; return }

func (a *address) SetTypeAndClear(t uint32) { _ = "STUB: not implemented"; return }
