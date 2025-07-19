package md2

import (
	"hash"
)

// hasher is used to calculate the MD2 hash-function.
type hasher struct {
	state    [48]byte
	checksum [Size]byte
	buffer   [BlockSize]byte // input buffer
	index    int             // this is the index into buffer
}

var _ hash.Hash = &hasher{}

// BlockSize returns the size of the MD2 block measured in byte.
//
// Size will always return 16.
//
// So, the size of the MD2 block is 16 bytes (i.e., 128 bits).
func (hasher) BlockSize() int {
	return BlockSize
}

// processBlock processes a single 16-byte (i.e., 128-bit) block.
func (receiver *hasher) processBlock(block []byte) {
	if nil == receiver {
		return
	}

	block = block[:BlockSize]

	copy(receiver.state[16:32], block)

	for i := 0; i < 16; i++ {
		receiver.state[32+i] = receiver.state[16+i] ^ receiver.state[i]
	}

	// Do 18 rounds.
	{
		var t byte = 0

		for j := 0; j < 18; j++ {
			for k := 0; k < 48; k++ {
				t = receiver.state[k] ^ sbox[t]
				receiver.state[k] = t
			}

			// We don't have to (and can't) do a MOD 256 here.
			// I.e.,:
			//
			//	t = (t + byte(j)) % 256
			//
			// Addition using a byte is already MOD 256.
			//
			// IETF RFC-1319 is show a MOD 256 done here.
			t = (t + byte(j))
		}
	}

	// update the checksum
	{
		l := receiver.checksum[15]
		for i := 0; i < len(receiver.checksum); i++ {
			c := block[i]
			receiver.checksum[i] ^= sbox[c^l]
			l = receiver.checksum[i]
		}
	}
}

// Reset resets hasher to the initial state,
// as per how it is defined in IETF RFC-1319.
//
// https://datatracker.ietf.org/doc/html/rfc1319
func (receiver *hasher) Reset() {
	if nil == receiver {
		return
	}

	// zero the whole struct
	*receiver = hasher{}
}

// Size returns the size of the MD2 digest measured in byte.
//
// Size will always return 16.
//
// So, the size of the MD2 digest is 16 bytes (i.e., 128 bits).
func (hasher) Size() int {
	return Size
}

// Sum appends the MD2 digest of the data that had been written to the hasher so far.
//
// Sum appends the the digest to `p` (rather than just returning the digest).
//
// To have Sum return the digest, pass nil to Sum.
// For example:
//
//	digest := hasher.Sum(nil)
func (receiver *hasher) Sum(p []byte) []byte {

	// Make a copy of the hasher, so that when sum() writes to the receiver
	// as part of calculating the (current) digest, it doesn't affect the
	// original receiver. And, the original receiver can have more data written to it.
	clone := *receiver

	digest := clone.sum()
	return append(p, digest[:]...)
}

// sum returns the MD2 digest (checksum) of the data already written to the hasher.
//
// Note that sum mutates the receiver.
func (receiver *hasher) sum() [Size]byte {
	if nil == receiver {
		return nothing
	}

	{
		// There might be extra space in this array
		// that we won't actually use.
		var padding [BlockSize]byte

		padLength := byte(BlockSize - receiver.index)

		var limit uint = uint(padLength)
		for i:=uint(0); i<limit; i++ {
			padding[i] = padLength
		}

		receiver.Write(padding[:padLength])
	}

	receiver.Write(receiver.checksum[:])

	{
		var digest [Size]byte
		copy(digest[:], receiver.state[:Size])
		return digest
	}
}

// Write puts in more data into the ongoing MD2 hasher.
//
// Write makes hasher match the io.Writer interface.
func (receiver *hasher) Write(p []byte) (n int, err error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	n = len(p)

	if receiver.index > 0 {
		numCopied := copy(receiver.buffer[receiver.index:], p)
		receiver.index += numCopied
		if receiver.index == BlockSize {
			receiver.processBlock(receiver.buffer[:])
			receiver.index = 0
		}
		p = p[numCopied:]
	}

	for len(p) >= BlockSize {

		// .processBlock() will only process the first 16 bytes of the slise,
		// so do NOT need to make the slice passed to .processBlock() be
		// exactly 64 bytes.
		receiver.processBlock(p)
		p = p[BlockSize:]
	}

	if len(p) > 0 {
		receiver.index = copy(receiver.buffer[:], p)
	}

	return
}
