package byteutils

// Bytes can be converted to and from other built-in types.
type Bytes = []byte

// BytesFromUint16 creates bytes from a uint16.
func BytesFromUint16(n uint16, e Endian) Bytes {
	bytes := make(Bytes, 2)
	e.IterateUint16(n, func(b *byte, index int) {
		bytes[len(bytes)-index-1] = *b
	})
	return bytes
}

// BytesFromUint32 creates bytes from a uint32.
func BytesFromUint32(n uint32, e Endian) Bytes {
	bytes := make(Bytes, 4)
	e.IterateUint32(n, func(b *byte, index int) {
		bytes[len(bytes)-index-1] = *b
	})
	return bytes
}

// ToUint16 converts Bytes to uint16
func ToUint16(bytes Bytes, e Endian) uint16 {
	smallest, largest := e.byteRange(2)
	return uint16(bytes[smallest]) | (uint16(bytes[largest]) << 8)
}

// ToUint32 converts Bytes to uint16
func ToUint32(bytes Bytes, e Endian) uint32 {
	var result uint32

	e.IterateSmallestToLargest(bytes, func(v *byte, enum int) {
		result |= uint32(*v) << (enum * 8)
	})

	return result
}
