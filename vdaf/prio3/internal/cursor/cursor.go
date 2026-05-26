// Package cursor aids with iteration of slices.
package cursor

type Cursor[V ~[]E, E any] []E

func New[V ~[]E, E any](x V) Cursor[V, E] { _ = "STUB: not implemented"; return nil }

// Next return an slice of size n and advances the pointer.
func (s *Cursor[V, E]) Next(n uint) (out V) { _ = "STUB: not implemented"; return *new(V) }
