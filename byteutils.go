// Package byteutils contains utilities for managing bytes.
//
// Note that a byte's indices are [0, 8), and using an index beyond this range
// may cause a panic or undefined behavior.
package byteutils

// Bit abstracts a single bit. The underlying type is byte instead of
// bool to make conversion usage with bytes simpler.
type Bit = byte

const (
	// Zero represents a 0 bit
	Zero Bit = 0
	// One represents a 1 bit
	One Bit = 1
)

// SetL sets the nth bit from the left.
func SetL(b *byte, index byte) {
	SetR(b, 7-index)
}

// ClearL clears the nth bit from the left.
func ClearL(b *byte, index byte) {
	ClearR(b, 7-index)
}

// SetR sets the nth bit from the right.
func SetR(b *byte, index byte) {
	*b |= (1 << index)
}

// ClearR clears the nth bit from the right.
func ClearR(b *byte, index byte) {
	*b &= ^(1 << index)
}

// ChangeL changes the nth bit from the right to the bit provided.
func ChangeL(b *byte, index byte, bit Bit) {
	ChangeR(b, 7-index, bit)
}

// ChangeR changes the nth bit from the left to the bit provided.
func ChangeR(b *byte, index byte, bit Bit) {
	if BitAsBool(bit) {
		SetR(b, index)
	} else {
		ClearR(b, index)
	}
}

// GetL gets the nth bit from the right.
func GetL(b byte, index byte) Bit {
	return GetR(b, 7-index)
}

// GetR gets the nth bit from the left.
func GetR(b byte, index byte) Bit {
	return normalize(b & (1 << index))
}

// SliceL gets the [i, end) bits from the left.
func SliceL(b byte, i, end byte) byte {
	return (b >> (8 - end)) & (byte(0xFF) >> i)
}

// SliceR gets the [i, end) bits from the right.
//
// i will be the "index" of the bit farthest to the right, and end is the index
// of the bit further left from the right.
func SliceR(b byte, i, end byte) byte {
	return (b >> i) & (byte(0xFF) >> (8 - end))
}

// ToggleL flips the nth bit from the left.
func ToggleL(b *byte, index byte) {
	ToggleR(b, 7-index)
}

// ToggleR flips the nth bit from the right.
func ToggleR(b *byte, index byte) {
	*b ^= 1 << index

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

// BitAsBool converts a bit into a bool.
func BitAsBool(b Bit) bool {
	return b != Zero
}

// Normalize forces a bit to be either Zero or One
func normalize(b Bit) Bit {
	if BitAsBool(b) {
		return One
	}
	return Zero
}
