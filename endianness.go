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