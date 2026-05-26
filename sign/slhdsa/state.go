package slhdsa

import (
	"hash"
	"io"

	"github.com/cloudflare/circl/internal/sha3"
)

// statePriv encapsulates common data for performing a private operation.
type statePriv struct {
	state
	PRF statePRF
}

func (s *statePriv) Size(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

func (p *params) NewStatePriv(skSeed, pkSeed []byte) (s statePriv) {
	_ = "STUB: not implemented"
	return *new(statePriv)
}

func (s *statePriv) Clear() { _ = "STUB: not implemented"; return }

// state encapsulates common data for performing a public operation.
type state struct {
	*params

	F stateF
	H stateH
	T stateT
}

func (s *state) Size(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

func (p *params) NewStatePub(pkSeed []byte) (s state) {
	_ = "STUB: not implemented"
	return *new(state)
}

func (s *state) init(p *params, c *cursor, pkSeed []byte) {
	s.params = p
	s.F.Init(p, c, pkSeed)
	s.H.Init(p, c, pkSeed)
	s.T.Init(p, c, pkSeed)
}

func (s *state) Clear() { _ = "STUB: not implemented"; return }

func sha256sum(out, in []byte) { _ = "STUB: not implemented"; return }
func sha512sum(out, in []byte) { _ = "STUB: not implemented"; return }

type baseHasher struct {
	hash          func(out, in []byte)
	input, output []byte
	address
}

func (b *baseHasher) Size(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

func (b *baseHasher) Clear() { _ = "STUB: not implemented"; return }

func (b *baseHasher) Final() []byte { _ = "STUB: not implemented"; return nil }

type statePRF struct{ baseHasher }

func (s *statePRF) Init(p *params, cur *cursor, skSeed, pkSeed []byte) {
	_ = "STUB: not implemented"
	return
}

func (s *statePRF) Size(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

func (s *statePRF) padSize(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

type stateF struct {
	msg []byte
	baseHasher
}

func (s *stateF) Init(p *params, cur *cursor, pkSeed []byte) { _ = "STUB: not implemented"; return }

func (s *stateF) SetMessage(msg []byte) { _ = "STUB: not implemented"; return }

func (s *stateF) Clear() { _ = "STUB: not implemented"; return }

func (s *stateF) Size(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

func (s *stateF) padSize(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

type stateH struct {
	msg0, msg1 []byte
	baseHasher
}

func (s *stateH) Init(p *params, cur *cursor, pkSeed []byte) { _ = "STUB: not implemented"; return }

func (s *stateH) SetMsgs(m0, m1 []byte) { _ = "STUB: not implemented"; return }

func (s *stateH) Clear() { _ = "STUB: not implemented"; return }

func (s *stateH) Size(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

func (s *stateH) padSize(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

type stateT struct {
	hash interface {
		io.Writer
		Reset()
		Final([]byte)
	}
	input, output []byte
	address
}

func (s *stateT) Init(p *params, cur *cursor, pkSeed []byte) { _ = "STUB: not implemented"; return }

func (s *stateT) Clear() { _ = "STUB: not implemented"; return }

func (s *stateT) Reset() { _ = "STUB: not implemented"; return }

func (s *stateT) WriteMessage(msg []byte) { _ = "STUB: not implemented"; return }

func (s *stateT) Final() []byte { _ = "STUB: not implemented"; return nil }

func (s *stateT) Size(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

func (s *stateT) outputSize(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

func (s *stateT) padSize(p *params) uint32 { _ = "STUB: not implemented"; return 0 }

type sha2rw struct{ hash.Hash }

func (s *sha2rw) Final(out []byte)         { _ = "STUB: not implemented"; return }
func (s *sha2rw) SumIdempotent(out []byte) { _ = "STUB: not implemented"; return }

type sha3rw struct{ sha3.State }

func (s *sha3rw) Final(out []byte)         { _ = "STUB: not implemented"; return }
func (s *sha3rw) SumIdempotent(out []byte) { _ = "STUB: not implemented"; return }

type (
	item struct {
		node []byte
		z    uint32
	}
	stackNode []item
)

func (p *params) NewStack(z uint32) stackNode { _ = "STUB: not implemented"; return *new(stackNode) }

func (s stackNode) isEmpty() bool { _ = "STUB: not implemented"; return false }
func (s stackNode) top() item     { _ = "STUB: not implemented"; return *new(item) }
func (s *stackNode) push(v item)  { _ = "STUB: not implemented"; return }

func (s *stackNode) pop() (v item) { _ = "STUB: not implemented"; return *new(item) }

func (s *stackNode) Clear() { _ = "STUB: not implemented"; return }

type cursor []byte

func (c *cursor) Rest() []byte               { _ = "STUB: not implemented"; return nil }
func (c *cursor) Next(n uint32) (out []byte) { _ = "STUB: not implemented"; return nil }

func clearSlice(s *[]byte) { _ = "STUB: not implemented"; return }
