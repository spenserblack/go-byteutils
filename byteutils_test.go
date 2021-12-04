package byteutils

import "testing"

// TestSetL checks that a bit from the left will be set.
func TestSetL(t *testing.T) {
	b := Byte(0)
	b.SetL(0)
	if b != 0b1000_0000 {
		t.Fatalf(`b = %08b, want 10000000`, b)
	}
}

// TestSetR checks that a bit from the left will be set.
func TestSetR(t *testing.T) {
	b := Byte(0)
	b.SetR(0)
	if b != 0b0000_0001 {
		t.Fatalf(`b = %08b, want 00000001`, b)
	}
}
