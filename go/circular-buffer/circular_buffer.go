package circular

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	items   []byte
	toRead  int
	toWrite int
	used int
}

func NewBuffer(size int) *Buffer {
	result := new(Buffer)
	result.items = make([]byte, size)
	result.toRead = 0
	result.toWrite = 0
	result.used = 0
	return result
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.used == 0{
		return 0, errors.New("buffer is empty")
	}
	result := b.items[b.toRead]
	b.used--
	b.toRead = (b.toRead + 1) % len(b.items)
	return result, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.used == len(b.items) {
		return errors.New("buffer is empty")
	}
	b.items[b.toWrite] = c
	b.used++
	b.toWrite = (b.toWrite + 1) % len(b.items)
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	panic("Please implement the Overwrite function")
}

func (b *Buffer) Reset() {
	panic("Please implement the Reset function")
}
