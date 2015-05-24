package loops

import (
	"image/color"
	"log"

	gl "github.com/go-gl/gl/v4.1-core/gl"
)

type Graphics struct {
	*Program
	*LoopWindow
	defaultBuffer      *Buffer
	defaultVertexArray *VertexArray
	buffersCreated     bool
}

func NewGraphics(window *LoopWindow, program *Program) *Graphics {
	return &Graphics{
		Program:        program,
		LoopWindow:     window,
		buffersCreated: false,
	}
}

func (self *Graphics) SetMat(mat Mat4) {
	loc := self.Program.GetUniformLocation("mat")
	gl.UniformMatrix4fv(loc, 1, false, &mat[0])
}

func (self *Graphics) SetColor(c color.RGBA) {
	r := float32(c.R) / float32(255)
	g := float32(c.G) / float32(255)
	b := float32(c.B) / float32(255)
	a := float32(c.A) / float32(255)

	loc := self.Program.GetUniformLocation("color")
	gl.Uniform4f(loc, r, g, b, a)
}

func (self *Graphics) Draw(mode uint32, verts []float32) {
	self.bindDefaultBuffer()
	gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.STATIC_DRAW)
	gl.DrawArrays(mode, 0, int32(len(verts)))
}

func (self *Graphics) DrawRect(x, y, w, h float32) {
	self.Draw(gl.TRIANGLE_STRIP, []float32{
		x, y,
		x, y + h,
		x + w, y,
		x + w, y + h,
	})
}

func (self *Graphics) bindDefaultBuffer() {
	if !self.buffersCreated {
		log.Print("Creating default buffer")
		self.defaultBuffer = NewBuffer()
		self.defaultVertexArray = NewVertexArray()

		self.buffersCreated = true
	}

	self.defaultBuffer.Bind(gl.ARRAY_BUFFER)
	self.defaultVertexArray.Bind()

	loc := uint32(self.Program.GetAttribLocation("v_position"))
	gl.EnableVertexAttribArray(loc)
	gl.VertexAttribPointer(loc, 2, gl.FLOAT, false, 0, nil)
}
