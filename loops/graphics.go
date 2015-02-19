package loops

import (
	"log"

	gl "github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

type Graphics struct {
	*gl.Program
	*LoopWindow
	defaultBuffer        gl.Buffer
	defaultBufferCreated bool
}

func NewGraphics(window *LoopWindow, program *gl.Program) *Graphics {
	return &Graphics{
		Program:              program,
		LoopWindow:           window,
		defaultBufferCreated: false,
	}
}

func (self *Graphics) SetMat(mat Mat4) {
	loc := self.Program.GetUniformLocation("mat")
	loc.UniformMatrix4f(false, (*[16]float32)(&mat))
}

func (self *Graphics) Draw(mode gl.GLenum, verts []float32) {
	self.bindDefaultBuffer()
	gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(verts), verts, gl.STATIC_DRAW)
	gl.DrawArrays(mode, 0, len(verts))
}

func (self *Graphics) bindDefaultBuffer() {
	if !self.defaultBufferCreated {
		log.Print("Creating default buffer")
		self.defaultBuffer = gl.GenBuffer()
		self.defaultBufferCreated = true
	}

	self.defaultBuffer.Bind(gl.ARRAY_BUFFER)
	loc := self.Program.GetAttribLocation("position")
	loc.EnableArray()
	loc.AttribPointer(2, gl.FLOAT, false, 0, nil)
}
