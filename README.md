# byteutils

[![codecov](https://codecov.io/gh/spenserblack/go-byteutils/branch/master/graph/badge.svg?token=DfSRqth9QW)](https://codecov.io/gh/spenserblack/go-byteutils)

Utilities for managing bytes, such as indexing into a byte to get an individual
bit and converting types like `uint` to and from `[]byte`.

## Example

```go
fmt.Println(byteutils.ToUint16([]byte{0x01, 0x00})) // 256

b := Byte(0b0001_0000)
fmt.Println(b.Indexl(3)) // 1 (index from left)
b.Setl(4, byteutils.One) // set index 4 from right
fmt.Println(b.Indexr(3)) // 1 (index from right)
```
