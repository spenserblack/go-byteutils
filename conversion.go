package byteutils

// Endian represents the endianness for conversion.
type Endian bool

const (
	LittleEndian Endian = false
	BigEndian    Endian = true
)

// Bytes can be converted to and from other built-in types.
type Bytes []Byte

// Convert is the type to manage converting to other types.
type Convert interface{
	ToUint16(Endian) uint16
	ToUint32(Endian) uint32
}

// ToUint16 converts Bytes to uint16
func (b Bytes) ToUint16(e Endian) uint16 {
	smallest, largest := e.byteRange(2)
	return uint16(b[smallest]) | (uint16(b[largest]) << 8)
}

// ToUint32 converts Bytes to uint16
func (b Bytes) ToUint32(e Endian) uint32 {
	var result uint32
	smallest, largest := e.byteRange(4)

	enumeration := 0
	if e == BigEndian {
		for i := smallest; i <= largest; i++ {
			result |= uint32(b[i]) << (enumeration * 8)
			enumeration++
		}
	} else {
		for i := largest; i <= smallest; i-- {
			result |= uint32(b[i]) << ((enumeration) * 8)
			enumeration++
		}
	}
	return result
}

// ByteRange returns the index of the smallest and the largest bytes.
func (e Endian) byteRange(byteCount int) (smallest, largest int) {
	if e == LittleEndian {
		smallest = byteCount - 1
	} else {
		largest = byteCount - 1
	}
	return
}
