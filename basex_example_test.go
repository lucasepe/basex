package basex_test

import (
	"fmt"

	"github.com/lucasepe/basex"
)

const (
	Base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Example_encode() {
	msg := []byte("hello, world")
	enc, err := basex.NewEncoding(Base62)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", enc.Encode(msg))
	// Output:
	// "g1CDtPywO1kG1foi"
}

func Example_decode() {
	enc, err := basex.NewEncoding(Base62)
	if err != nil {
		panic(err)
	}
	decoded, err := enc.Decode("g1CDtPywO1kG1foi")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", decoded)
	// Output:
	// "hello, world"
}
