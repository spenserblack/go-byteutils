package byteutils

import "testing"

// TestAsBool checks that a bit is false when it is 0, true otherwise.
func TestAsBool(t *testing.T) {
	if Zero.AsBool() {
		t.Error("want Zero to be false")
	}
	if !One.AsBool() {
		t.Error("want One to be true")
	}
}

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
