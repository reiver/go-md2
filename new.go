package md2

import (
	"hash"
)

// New returns a hash.Hash that calculates an MD2 digest (checksum) using the MD2 cryptographic hash-function.
func New() hash.Hash {
	var h hash.Hash = newHasher()
	return h
}

func newHasher() *hasher {
	var h hasher
	h.Reset()
	return &h
}
