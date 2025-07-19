package md2_test

import (
	"fmt"
	"io"

	"github.com/reiver/go-md2"
)

func ExampleNew() {

	hasher := md2.New()

	// “Yesterday I was clever, so I wanted to change the world. Today I am wise, so I am changing myself.”
	// ⸺ Rumi
	io.WriteString(hasher, "“")
	io.WriteString(hasher, "Yesterday I was clever, so I wanted to change the world.")
	io.WriteString(hasher, " ")
	io.WriteString(hasher, "Today I am wise, so I am changing myself.")
	io.WriteString(hasher, "”")
	io.WriteString(hasher, "\n")
	io.WriteString(hasher, "”")
	io.WriteString(hasher, "⸺")
	io.WriteString(hasher, " ")
	io.WriteString(hasher, "Rumi")

	digest := hasher.Sum(nil)

	fmt.Printf("MD2-DIGEST: 0x%032X\n", digest)

	// Output:
	// MD2-DIGEST: 0xDA17B23E385A5B2F3756D762E1CFDA43
}
