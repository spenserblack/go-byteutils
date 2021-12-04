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

// AsBool convers a bit into a bool.
func (b Bit) AsBool() bool {
	return b != Zero
}
