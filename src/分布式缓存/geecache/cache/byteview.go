package cache

/*
缓存值的抽象与封装
 */


// A ByteView holds an immutable view of bytes.
type ByteView struct {
	buf []byte
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.buf)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.buf)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.buf)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}