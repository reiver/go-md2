package md2

// Size is the size of an MD2 digest (checksum) measured in bytes.
//
// The size of an MD2 digest 16 bytes (i.e., 128 bits).
// For example:
//
//	//          1    2    3    4    5    6    7    8    9   10   11   12   13   14   15   16
//	[16]byte{0x83,0x50,0xe5,0xa3,0xe2,0x4c,0x15,0x3d,0xf2,0x27,0x5c,0x9f,0x80,0x69,0x27,0x73}
//
// Note that if you are representing an MD2 digest in hexadecimal, then it will double the number of characters to 32 bytes.
// For example:
//
//	"8350e5a3e24c153df2275c9f80692773"
//
// Or, if you include the "0x" prefix in front of the hexadecimal representation, then it will be 32+2=34 bytes long:
//
//	0x8350e5a3e24c153df2275c9f80692773
//
// Size represents the length of the MD2 digest in a binary representation.
const Size = 16
