package byteutils

// Bytes can be converted to and from other built-in types.
type Bytes []Byte

// Convert is the type to manage converting to other types.
type Convert interface {
	ToUint16(Endian) uint16
	ToUint32(Endian) uint32
}

// BytesFromUint16 creates bytes from a uint16.
func BytesFromUint16(n uint16, e Endian) Bytes {
	bytes := make(Bytes, 2)
	e.IterateUint16(n, func(b Byte, index int) {
		bytes[len(bytes)-index-1] = b
	})
	return bytes
}

// BytesFromUint32 creates bytes from a uint32.
func BytesFromUint32(n uint32, e Endian) Bytes {
	bytes := make(Bytes, 4)
	e.IterateUint32(n, func(b Byte, index int) {
		bytes[len(bytes)-index-1] = b
	})
	return bytes
}

// ToUint16 converts Bytes to uint16
func (b Bytes) ToUint16(e Endian) uint16 {
	smallest, largest := e.byteRange(2)
	return uint16(b[smallest]) | (uint16(b[largest]) << 8)
}

// ToUint32 converts Bytes to uint16
func (b Bytes) ToUint32(e Endian) uint32 {
	var result uint32

	e.IterateSmallestToLargest(b, func(v Byte, enum int) {
		result |= uint32(v) << (enum * 8)
	})

	return result
}
