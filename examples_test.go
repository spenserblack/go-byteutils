package byteutils_test

import (
	"fmt"

	"github.com/spenserblack/go-byteutils"
)

// Bitwise methods to manage bytes and their bits are abstracted with this
// library.
//
// Functions are provided to "index" into a byte from the left and from the
// right.
func Example() {
	var b byte = 0b0000_0000
	byteutils.SetL(&b, 1)
	fmt.Printf("%08b\n", b)
	byteutils.ChangeR(&b, 0, byteutils.One)
	fmt.Printf("%08b\n", b)
	byteutils.ToggleL(&b, 7)
	fmt.Printf("%08b\n", b)
	fmt.Println(byteutils.GetR(b, 6))
	// Output:
	// 01000000
	// 01000001
	// 01000000
	// 1
}

// Endianness determines the least and most significant bytes.
//
// Little endian means that the last byte is the smallest.
// Big endian means that the last byte is the largest.
func Example_endianness() {
	bytes := []byte{0x00, 0x01}

	littleEndianConversion := byteutils.ToUint16(bytes, byteutils.LittleEndian)
	bigEndianConversion := byteutils.ToUint16(bytes, byteutils.BigEndian)

	fmt.Printf("%d (0x%04X)\n", littleEndianConversion, littleEndianConversion)
	fmt.Printf("%d (0x%04X)\n", bigEndianConversion, bigEndianConversion)
	// Output:
	// 1 (0x0001)
	// 256 (0x0100)
}
