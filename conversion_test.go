package byteutils

import "testing"

// TestToUint16 tests that Bytes get converted to 16-bit unsigned integer.
func TestToUint16(t *testing.T) {
	var b Convert = Bytes{0x0F, 0xA0}

	if actual := b.ToUint16(BigEndian); actual != 0xA00F {
		t.Errorf(`result = %04X, want A00F`, actual)
	}
	if actual := b.ToUint16(LittleEndian); actual != 0x0FA0 {
		t.Errorf(`result = %04X, want 0FA0`, actual)
	}
}

// TestToUint32 tests that Bytes get converted to 16-bit unsigned integer.
func TestToUint32(t *testing.T) {
	var b Convert = Bytes{0x0F, 0x00, 0x00, 0xA0}

	if actual := b.ToUint32(BigEndian); actual != 0xA0_00_00_0F {
		t.Errorf(`result = %08X, want A000000F`, actual)
	}
	if actual := b.ToUint32(LittleEndian); actual != 0x0F_00_00_A0 {
		t.Errorf(`result = %08X, want 0F0000A0`, actual)
	}
}

// TestIteration tests that little endian and big endian collections of bytes
// are correctly iterated over.
func TestIteration(t *testing.T) {
	b := Bytes{0x00, 0x08, 0x10}
	bigEndianOrder := []Byte{0x00, 0x08, 0x10}
	littleEndianOrder := []Byte{0x10, 0x08, 0x00}

	LittleEndian.IterateSmallestToLargest(
		b Byte, 
		func(b Byte, index int) {
			if want := littleEndianOrder[index]; b != want {
				t.Errorf(`little endian b = %02X, want %02X`, b, want)
			}
		},
	)
	BigEndian.IterateSmallestToLargest(
		b Byte, 
		func(b Byte, index int) {
			if want := bigEndianOrder[index]; b != want {
				t.Errorf(`big endian b = %02X, want %02X`, b, want)
			}
		},
	)
}
