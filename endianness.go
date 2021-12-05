package byteutils

// Endian represents the endianness for conversion.
type Endian bool

const (
	LittleEndian Endian = false
	BigEndian    Endian = true
)

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

// IterateUint16 iterates over a uint16 as bytes, from smallest to largest.
// Endianness determines if iteration goes from the left-most byte to the
// right-most (big endian), or the right-most byte to the left-most (little
// endian).
func (e Endian) IterateUint16(n uint16, f ByteIteratorFunc) {
	smallest, largest := e.byteRange(2)

	if e == BigEndian {
		for i := smallest; i <= largest; i++ {
			var intersect uint16 = 0xFF << i
			b := Byte((n & intersect) >> i)
			f(b, i)
		}
	} else {
		for i := smallest; i >= largest; i-- {
			enumeration := smallest - i
			var intersect uint16 = 0xFF << i
			b := Byte((n & intersect) >> i)
			f(b, enumeration)
		}
	}
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
