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

// TestBytesFromUint16 tests that a uint16 can be created from a little-endian
// and big-endian slice of bytes.
func TestBytesFromUint16(t *testing.T) {
	littleEndian := BytesFromUint16(0xF010, LittleEndian)
	bigEndian := BytesFromUint16(0xF010, BigEndian)
	littleByteOrder := Bytes{0xF0, 0x10}
	bigByteOrder := Bytes{0x10, 0xF0}

	for i, b := range littleEndian {
		want := littleByteOrder[i]
		t.Errorf(`byte %d = %02X, want %02X`, i, b, want)
	}
	for i, b := range bigEndian {
		want := bigByteOrder[i]
		t.Errorf(`byte %d = %02X, want %02X`, i, b, want)
	}
}
