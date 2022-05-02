package space

type BytesSpace []byte

// Implementation of cache.Sized interface
func (c BytesSpace) Size() int {
	return len(c)
}

func (c BytesSpace) String() string {
	return string(c)
}

func (c BytesSpace) CloneBytes() []byte {
	bytes := make([]byte, len(c))
	copy(bytes, c)

	return c
}
