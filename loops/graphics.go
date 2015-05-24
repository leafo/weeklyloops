package loops

import (
	"image/color"
	"log"

	gl "github.com/go-gl/gl/v4.1-core/gl"
)

type Graphics struct {
	*LoopWindow
	defaultBuffer      *Buffer
	defaultVertexArray *VertexArray
	buffersCreated     bool

	currentMat   Mat4
	currentColor color.RGBA
}

func NewGraphics(window *LoopWindow) *Graphics {
	return &Graphics{
		LoopWindow:     window,
		buffersCreated: false,
	}
}

func (self *Graphics) SetMat(mat Mat4) {
	self.currentMat = mat
}

func (self *Graphics) SetColor(c color.RGBA) {
	self.currentColor = c
}

func (self *Graphics) Draw(mode uint32, verts []float32) {
	program := self.LoopWindow.programSolid2d

	self.bindBuffers()
	self.bindProgram(program)

	numVerts := len(verts) / 2
	gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.STATIC_DRAW)

	loc := uint32(program.GetAttribLocation("v_position"))
	gl.EnableVertexAttribArray(loc)
	gl.VertexAttribPointer(loc, 2, gl.FLOAT, false, 0, nil)

	gl.DrawArrays(mode, 0, int32(numVerts))
}

func (self *Graphics) DrawRect(x, y, w, h float32) {
	self.Draw(gl.TRIANGLE_STRIP, []float32{
		x, y,
		x, y + h,
		x + w, y,
		x + w, y + h,
	})
}

// func (self *Graphics) DrawColored(mode uint32, verts []float32) {
// 	self.bindDefaultBuffer()
// 	gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.STATIC_DRAW)
// 	gl.DrawArrays(mode, 0, int32(len(verts)))
// }

func (self *Graphics) bindBuffers() {
	if !self.buffersCreated {
		log.Print("Creating default buffer")
		self.defaultBuffer = NewBuffer()
		self.defaultVertexArray = NewVertexArray()

		self.buffersCreated = true
	}

	self.defaultBuffer.Bind(gl.ARRAY_BUFFER)
	self.defaultVertexArray.Bind()

}

func (self *Graphics) bindProgram(program Program) {
	program.Use()
	loc := program.GetUniformLocation("mat")
	gl.UniformMatrix4fv(loc, 1, false, &self.currentMat[0])

	c := self.currentColor
	r := float32(c.R) / float32(255)
	g := float32(c.G) / float32(255)
	b := float32(c.B) / float32(255)
	a := float32(c.A) / float32(255)

	loc = program.GetUniformLocation("color")
	gl.Uniform4f(loc, r, g, b, a)
}
