# byteutils

[![CI](https://github.com/spenserblack/go-byteutils/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/go-byteutils/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/go-byteutils/branch/master/graph/badge.svg?token=DfSRqth9QW)](https://codecov.io/gh/spenserblack/go-byteutils)

Utilities for managing bytes, such as indexing into a byte to get an individual
bit and converting types like `uint` to and from `[]byte`.

## Example

```go
fmt.Println(byteutils.ToUint16([]byte{0x01, 0x00})) // 256

b := Byte(0b0001_0000)
fmt.Println(b.GetL(3)) // 1 (index from left)
b.SetL(4) // set index 4 from left to 1
fmt.Println(b.GetR(3)) // 1 (index from right)
```
