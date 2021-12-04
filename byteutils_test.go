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
