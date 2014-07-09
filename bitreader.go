package bitreader

import "io"

type Reader interface {
	ReadBit() (bool, error)
	PeekBit() (bool, error)
	Trash(uint) error
}

type Reader32 interface {
	Reader
	Read32(uint) (uint32, error)
	Peek32(uint) (uint32, error)
}

func NewReader32(r io.Reader) Reader32 {
	return &simpleReader32{r, make([]byte, 4), 0, 0}
}
