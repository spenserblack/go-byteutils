package byteutils

// Endian represents the endianness for conversion.
type Endian bool

const (
	LittleEndian Endian = false
	BigEndian    Endian = true
)

// ByteIteratorFunc takes a byte and the enumeration (count of calls to
// function). A pointer is passed so that the byte can be potentially
// modified.
type ByteIteratorFunc = func(b *byte, enumeration int)

// IterateSmallestToLargest iterates from the smallest byte to the largest
// byte given the endianness. It will call the provided function on each byte.
func (e Endian) IterateSmallestToLargest(b Bytes, f ByteIteratorFunc) {
	smallest, largest := e.byteRange(len(b))

	if e == BigEndian {
		for i := smallest; i <= largest; i++ {
			f(&b[i], i)
		}
	} else {
		for i := smallest; i >= largest; i-- {
			f(&b[i], smallest-i)
		}
	}
}

// IterateUint16 iterates over a uint16 as bytes, from smallest to largest.
// Endianness determines if iteration goes from the left-most byte to the
// right-most (big endian), or the right-most byte to the left-most (little
// endian).
func (e Endian) IterateUint16(n uint16, f ByteIteratorFunc) {
	e.iterateNumber(n, f)
}

// IterateUint32 iterates over a uint32 as bytes, from smallest to largest.
// Endianness determines if iteration goes from the left-most byte to the
// right-most (big endian), or the right-most byte to the left-most (little
// endian).
func (e Endian) IterateUint32(n uint32, f ByteIteratorFunc) {
	e.iterateNumber(n, f)
}

// IterateNumber iterates over a number as bytes, from smallest to largest.
func (e Endian) iterateNumber(n interface{}, f ByteIteratorFunc) {
	var smallest, largest int
	var value uint32
	switch v := n.(type) {
	case uint16:
		smallest, largest = e.byteRange(2)
		value = uint32(v)
	case uint32:
		smallest, largest = e.byteRange(4)
		value = v
	}

	if e == BigEndian {
		for i := smallest; i <= largest; i++ {
			shift := (largest - i) * 8
			var intersect uint32 = 0xFF << shift
			b := byte((value & intersect) >> shift)
			f(&b, i)
		}
	} else {
		for i := largest; i <= smallest; i++ {
			shift := i * 8
			var intersect uint32 = 0xFF << shift
			b := byte((value & intersect) >> shift)
			f(&b, i)
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
