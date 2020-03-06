package wavreader

import (
	"fmt"
	"io"
)

// Reader ...
type Reader struct {
}

// New ...
func New(r io.ByteReader) (*Reader, error) {
	return nil, fmt.Errorf("not implemented")
}

// Len ...
func (r Reader) Len() uint64 {
	return 0
}

// Rate ...
func (r Reader) Rate() uint {
	return 0
}

// Chans ...
func (r Reader) Chans() uint {
	return 0
}

// At ...
func (r Reader) At(ch uint, offset uint64) float32 {
	return 0.0
}
