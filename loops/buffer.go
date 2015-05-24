package loops

import (
	gl "github.com/go-gl/gl/v4.1-core/gl"
)

type Buffer struct {
	Buffer uint32
}

func NewBuffer() *Buffer {
	var v uint32
	gl.GenBuffers(1, &v)
	return &Buffer{v}
}

func (self *Buffer) Bind(bufferType uint32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, self.Buffer)
}
