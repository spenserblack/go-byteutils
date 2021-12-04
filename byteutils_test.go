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

// TestNewBit checks that NewBit returns One for non-zero values with defined
// behavior, Zero otherwise.
func TestNewBit(t *testing.T) {
	if NewBit(true) != One {
		t.Error("want true bit to be set")
	}
	if NewBit(0) != Zero {
		t.Error("want zero bit to be clear")
	}
	if NewBit("undefined") != Zero {
		t.Error("expected undefined behavior to leave bit clear")
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

// TestClear checks that a bits are cleared.
func TestClear(t *testing.T) {
	b := Byte(0xFF)
	b.ClearL(1)
	b.ClearR(1)
	want := Byte(0b1011_1101)
	if b != want {
		t.Fatalf(`b = %08b, want %08b`, b, want)
	}
}

// TestChange checks that bits are changed.
func TestChange(t *testing.T) {
	b := Byte(0x0F)

	b.ChangeL(0, One)
	b.ChangeR(0, Zero)
	want := Byte(0b1000_1110)
	if b != want {
		t.Fatalf(`b = %08b, want %08b`, b, want)
	}
}

// TestGet checks that the correct value is returned when "indexing" into
// a byte.
func TestGet(t *testing.T) {
	b := Byte(0x0F)

	if actual := b.GetL(4); actual != Zero {
		t.Errorf(`4 from left of %08b = %v, want %v`, b, actual, Zero)
	}
	if actual := b.GetR(4); actual != One {
		t.Errorf(`4 from right of %08b = %v, want %v`, b, actual, One)
	}
}

// TestToggle checks that a bit is flipped and remaining bits remain the same.
func TestToggle(t *testing.T) {
	b := Byte(0x0F)
	b.ToggleR(3)
	if b != 0x07 {
		t.Errorf(`b = %08b, want %08b`, b, 0x07)
	}
	b.ToggleL(3)
	if b != 0x17 {
		t.Errorf(`b = %08b, want %08b`, b, 0x17)
	}
}
