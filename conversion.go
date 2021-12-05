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

	e.IterateSmallestToLargest(b, func(v Byte, enum int) {
		result |= uint32(v) << (enum * 8)
	})

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

// ByteIteratorFunc takes a byte and the enumeration (count of calls to
// function).
type ByteIteratorFunc = func(b Byte, enumeration int)

// IterateSmallestToLargest iterates from the smallest byte to the largest
// byte given the endianness. It will call the provided function on each byte.
func (e Endian) IterateSmallestToLargest(b Bytes, f ByteIteratorFunc) {
	smallest, largest := e.byteRange(len(b))

	if e == BigEndian {
		for i := smallest; i <= largest; i++ {
			f(b[i], i)
		}
	} else {
		for i := smallest; i >= largest; i-- {
			f(b[i], smallest - i)
		}
	}
}
