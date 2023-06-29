# basex
Fast base encoding / decoding of any given alphabet using bitcoin style leading
zero compression.

Credits to [https://github.com/eknkc/basex](https://github.com/eknkc/basex) (this repo is a modified fork).

**WARNING:** This module is **NOT RFC3548** compliant,  it cannot be used for base16 (hex), base32, or base64 encoding in a standards compliant manner. 

## Example

Base62

```go

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
```

### Alphabets

See below for a list of commonly recognized alphabets, and their respective base.

Base | Alphabet
------------- | -------------
2 | `01`
8 | `01234567`
11 | `0123456789a`
16 | `0123456789abcdef`
32 | `0123456789ABCDEFGHJKMNPQRSTVWXYZ`
32 | `ybndrfg8ejkmcpqxot1uwisza345h769` (z-base-32)
36 | `0123456789abcdefghijklmnopqrstuvwxyz`
58 | `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`
62 | `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
64 | `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`
67 | `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.!~`


## How it works

It encodes octet arrays by doing long divisions on all significant digits in the
array, creating a representation of that number in the new base. Then for every
leading zero in the input (not significant as a number) it will encode as a
single leader character. This is the first in the alphabet and will decode as 8
bits. The other characters depend upon the base. For example, a base58 alphabet
packs roughly 5.858 bits per character.

This means the encoded string 000f (using a base16, 0-f alphabet) will actually decode
to 4 bytes unlike a canonical hex encoding which uniformly packs 4 bits into each
character.

While unusual, this does mean that no padding is required and it works for bases
like 43. 


## LICENSE [MIT](LICENSE)
A direct derivation of the base58 implementation from [`bitcoin/bitcoin`](https://github.com/bitcoin/bitcoin/blob/f1e2f2a85962c1664e4e55471061af0eaa798d40/src/base58.cpp),  generalized for variable length alphabets.