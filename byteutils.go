// Utilities for managing bytes.
//
// Note that a byte's indices are [0, 8), and using an index beyond this range
// may cause a panic or undefined behavior.
package byteutils

// Byte is a wrapper around the built-in byte that provides additonal
// functionality.
type Byte byte

// Bit abstracts a single bit. The underlying type is byte instead of
// bool to make conversion usage with bytes simpler.
type Bit byte

const (
	// Zero represents a 0 bit
	Zero Bit = 0
	// One represents a 1 bit
	One Bit = 1
)

// SetL sets the nth bit from the left.
func (b *Byte) SetL(index byte) {
	b.SetR(7 - index)
}

// ClearL clears the nth bit from the left.
func (b *Byte) ClearL(index byte) {
	b.ClearR(7 - index)
}

// SetR sets the nth bit from the right.
func (b *Byte) SetR(index byte) {
	*b |= (1 << index)
}

// ClearR clears the nth bit from the right.
func (b *Byte) ClearR(index byte) {
	*b &= ^(1 << index)
}

// ChangeL changes the nth bit from the right to the bit provided.
func (b *Byte) ChangeL(index byte, bit Bit) {
	b.ChangeR(7 - index, bit)
}

// ChangeR changes the nth bit from the left to the bit provided.
func (b *Byte) ChangeR(index byte, bit Bit) {
	if bit.AsBool() {
		b.SetR(index)
	} else {
		b.ClearR(index)
	}
}

// GetR gets the nth bit from the right.
func (b Byte) GetR(index byte) Bit {
	return b.GetL(7 - index)
}

// GetL gets the nth bit from the left.
func (b Byte) GetL(index byte) Bit {
	return Bit(b & (1 << index)).normalize()
}

// ToggleR flips the nth bit from the right.
func (b *Byte) ToggleR(index byte) {
	b.ToggleL(7 - index)
}

// ToggleL flips the nth bit from the left.
func (b *Byte) ToggleL(index byte) {
	bit := b & (1 << index)
	if bit != 0 {
		bit = ^bit
	}
	*b &= Byte(bit)
}

// NewBit creates a bit. This is to simplify enforcing that a bit is either 
// Zero or One. Boolean type and most numerical types can be used as the bit.
// Floats, complex types, etc. are undefined behavior.
func NewBit(bit interface{}) Bit {
	set := false
	switch v := bit.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		set = v != 0
	case bool:
		set = v
	}
	if set {
		return One
	}
	return Zero
}

// AsBool convers a bit into a bool.
func (b Bit) AsBool() bool {
	return b != Zero
}

// Normalize forces a bit to be either Zero or One
func (b Bit) normalize() Bit {
	if b.AsBool() {
		return One
	}
	return Zero
}
