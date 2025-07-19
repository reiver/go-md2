package md2_test

import (
	"fmt"

	"github.com/reiver/go-md2"
)

func ExampleSum() {

	var data []byte = []byte("Hello world!")

	digest := md2.Sum(data)

	fmt.Printf("MD2-DIGEST: 0x%032X\n", digest)

	// Output:
	// MD2-DIGEST: 0x63503D3117AD33F941D20F57144ECE64
}
