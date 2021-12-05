package byteutils

import "testing"


// TestIteration tests that little endian and big endian collections of bytes
// are correctly iterated over.
func TestIteration(t *testing.T) {
	b := Bytes{0x00, 0x08, 0x10}
	bigEndianOrder := []Byte{0x00, 0x08, 0x10}
	littleEndianOrder := []Byte{0x10, 0x08, 0x00}

	LittleEndian.IterateSmallestToLargest(b, func(b Byte, index int) {
		if want := littleEndianOrder[index]; b != want {
			t.Errorf(`little endian b = %02X, want %02X`, b, want)
		}
	})
	BigEndian.IterateSmallestToLargest(b, func(b Byte, index int) {
		if want := bigEndianOrder[index]; b != want {
			t.Errorf(`big endian b = %02X, want %02X`, b, want)
		}
	})
}

// TestIterateUint16 tests that little endian and big endian collections of bytes
// are correctly iterated over.
func TestIterateUint16(t *testing.T) {
	var v uint16 = 0xF010
	littleEndian := []Byte{0x10, 0xF0}
	bigEndian := []Byte{0xF0, 0x10}

	LittleEndian.IterateUint16(b, func(b Byte, index int) {
		if want := littleEndian[index]; b != want {
			t.Errorf(`little endian b = %02X, want %02X`, b, want)
		}
	})
	BigEndian.IterateUint16(b, func(b Byte, index int) {
		if want := bigEndian[index]; b != want {
			t.Errorf(`big endian b = %02X, want %02X`, b, want)
		}
	})
}
