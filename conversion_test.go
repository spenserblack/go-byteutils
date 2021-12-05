package byteutils

import "testing"

// TestToUint16 tests that Bytes get converted to 16-bit unsigned integer.
func TestToUint16(t *testing.T) {
	var b Convert = Bytes{0x0F, 0xA0}

	if actual := b.ToUint16(BigEndian); actual != 0xA00F {
		t.Errorf(`result = %X, want A00F`, actual)
	}
	if actual := b.ToUint16(LittleEndian); actual != 0x0FA0 {
		t.Errorf(`result = %X, want 0FA0`, actual)
	}
}

// TestToUint32 tests that Bytes get converted to 16-bit unsigned integer.
func TestToUint32(t *testing.T) {
	var b Convert = Bytes{0x0F, 0x00, 0x00, 0xA0}

	if actual := b.ToUint32(BigEndian); actual != 0xA0_00_00_0F {
		t.Errorf(`result = %X, want A000000F`, actual)
	}
	if actual := b.ToUint32(LittleEndian); actual != 0x0F_00_00_A0 {
		t.Errorf(`result = %X, want 0F0000A0`, actual)
	}
}