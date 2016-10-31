package circular

import (
	"errors"
	"log"
)

const testVersion = 4

// Buffer ..
type Buffer struct {
	data     []byte
	cap      int
	len      int
	writePos int
	readPos  int
}

// NewBuffer ..
func NewBuffer(size int) *Buffer {
	buf := new(Buffer)
	buf.data = make([]byte, size)
	buf.cap = size
	return buf
}

// ReadByte ..
func (buf *Buffer) ReadByte() (byte, error) {
	if buf.len == 0 {
		return 0, errors.New("Buffer is empty!")
	}
	b := buf.data[buf.readPos]
	buf.len--
	buf.readPos++
	if buf.readPos == buf.cap {
		buf.readPos = 0
	}
	return b, nil
}

// WriteByte ..
func (buf *Buffer) WriteByte(c byte) error {
	if buf.cap == buf.len {
		return errors.New("Buffer is full!")
	}
	buf.data[buf.writePos] = c
	buf.len++
	buf.writePos++
	if buf.writePos == buf.cap {
		buf.writePos = 0
	}
	return nil
}

// Overwrite ..
func (buf *Buffer) Overwrite(c byte) {
	if buf.len < buf.cap {
		err := buf.WriteByte(c)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		buf.data[buf.writePos] = c

		buf.readPos++
		if buf.readPos == buf.cap {
			buf.readPos = 0
		}
		buf.writePos++
		if buf.writePos == buf.cap {
			buf.writePos = 0
		}
	}
}

// Reset ..
func (buf *Buffer) Reset() {
	buf.readPos = buf.writePos
	buf.len = 0
}
