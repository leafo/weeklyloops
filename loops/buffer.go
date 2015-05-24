package loops

import (
	gl "github.com/go-gl/gl/v4.1-core/gl"
)

type Buffer struct {
	Buffer uint32
}

type VertexArray struct {
	VertexArray uint32
}

func NewBuffer() *Buffer {
	var v uint32
	gl.GenBuffers(1, &v)
	return &Buffer{v}
}

func (self *Buffer) Bind(bufferType uint32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, self.Buffer)
}

func NewVertexArray() *VertexArray {
	var v uint32
	gl.GenVertexArrays(1, &v)
	return &VertexArray{v}
}

func (self *VertexArray) Bind() {
	gl.BindVertexArray(self.VertexArray)
}
