package md2

// Sum returns the MD2 digest (checksum) of `data` calculated using the MD2 hash-function.
func Sum(data []byte) [Size]byte {
	var h hasher
        h.Reset()
	h.Write(data)

	return h.sum()
}
