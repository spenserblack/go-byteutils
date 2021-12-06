package byteutils

import "testing"

// TestToUint16 tests that Bytes get converted to 16-bit unsigned integer.
func TestToUint16(t *testing.T) {
	b := Bytes{0x0F, 0xA0}

	if actual := ToUint16(b, BigEndian); actual != 0xA00F {
		t.Errorf(`result = %04X, want A00F`, actual)
	}
	if actual := ToUint16(b, LittleEndian); actual != 0x0FA0 {
		t.Errorf(`result = %04X, want 0FA0`, actual)
	}
}

// TestToUint32 tests that Bytes get converted to 16-bit unsigned integer.
func TestToUint32(t *testing.T) {
	b := Bytes{0x0F, 0x00, 0x00, 0xA0}

	if actual := ToUint32(b, BigEndian); actual != 0xA0_00_00_0F {
		t.Errorf(`result = %08X, want A000000F`, actual)
	}
	if actual := ToUint32(b, LittleEndian); actual != 0x0F_00_00_A0 {
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
		if b != want {
			t.Errorf(`byte %d = %02X, want %02X`, i, b, want)
		}
	}
	for i, b := range bigEndian {
		want := bigByteOrder[i]
		if b != want {
			t.Errorf(`byte %d = %02X, want %02X`, i, b, want)
		}
	}
}

// TestBytesFromUint32 tests that a uint32 can be created from a little-endian
// and big-endian slice of bytes.
func TestBytesFromUint32(t *testing.T) {
	var v uint32 = 0x3F_7F_BF_FF
	littleEndian := BytesFromUint32(v, LittleEndian)
	bigEndian := BytesFromUint32(v, BigEndian)
	littleByteOrder := Bytes{0x3F, 0x7F, 0xBF, 0xFF}
	bigByteOrder := Bytes{0xFF, 0xBF, 0x7F, 0x3F}

	for i, b := range littleEndian {
		want := littleByteOrder[i]
		if b != want {
			t.Errorf(`byte %d = %02X, want %02X`, i, b, want)
		}
	}
	for i, b := range bigEndian {
		want := bigByteOrder[i]
		if b != want {
			t.Errorf(`byte %d = %02X, want %02X`, i, b, want)
		}
	}
}
