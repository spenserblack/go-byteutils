package byteutils

// Byte is a wrapper around the built-in byte that provides additonal
// functionality.
type Byte byte

// Bit abstracts a single bit. The underlying type is byte instead of
// bool to make conversion to primitive types easier.
type Bit byte

const (
	// Zero represents a 0 bit
	Zero Bit = 0
	// One represents a 1 bit
	One = 1
)

// SetL sets the nth bit from the left.
func (b *Byte) SetL(index byte) {
	*b |= (1 << (7 - index))
}
