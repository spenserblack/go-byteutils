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
	*b |= (1 << (7 - index))
}

// SetR sets the nth bit from the right.
func (b *Byte) SetR(index byte) {
	*b |= (1 << index)
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
