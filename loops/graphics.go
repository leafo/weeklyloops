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

	objectMat    Mat4
	viewMat      Mat4
	currentColor color.RGBA

	programSolid2d   Program
	programColored2d Program
	programSolid3d   Program
}

func NewGraphics(window *LoopWindow) *Graphics {
	return &Graphics{
		LoopWindow:     window,
		buffersCreated: false,

		objectMat: NewIdentityMat4(),
		viewMat:   NewIdentityMat4(),

		programSolid2d:   NewProgram(programSolid2d.ShaderSources...),
		programColored2d: NewProgram(programColored2d.ShaderSources...),
		programSolid3d:   NewProgram(programSolid3d.ShaderSources...),
	}
}

func (self *Graphics) SetMat(mat Mat4) {
	self.objectMat = mat
}

func (self *Graphics) SetViewMat(mat Mat4) {
	self.viewMat = mat
}

func (self *Graphics) SetColor(c color.RGBA) {
	self.currentColor = c
}

func (self *Graphics) Draw(mode uint32, verts []float32) {
	program := self.programSolid2d

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

// v = [ [x,y,r,g,b,a], ... ]
func (self *Graphics) DrawColored(mode uint32, verts []float32) {
	program := self.programColored2d

	numVerts := len(verts) / 6

	self.bindBuffers()
	self.bindProgram(program)

	gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.STATIC_DRAW)

	loc := uint32(program.GetAttribLocation("v_position"))
	gl.EnableVertexAttribArray(loc)
	gl.VertexAttribPointer(loc, 2, gl.FLOAT, false, 6*4, gl.PtrOffset(0))

	loc = uint32(program.GetAttribLocation("v_color"))
	gl.EnableVertexAttribArray(loc)
	gl.VertexAttribPointer(loc, 4, gl.FLOAT, false, 6*4, gl.PtrOffset(2*4))

	gl.DrawArrays(mode, 0, int32(numVerts))
}

func (self *Graphics) DrawShape(shape Shape) {
	self.DrawColored(shape.DrawMode(), shape.Verts())
}

func (self *Graphics) Draw3d(mode uint32, verts []float32, indexes []byte) {
	program := self.programSolid3d
	self.bindBuffers()
	self.bindProgram(program)

	gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.STATIC_DRAW)

	loc := uint32(program.GetAttribLocation("v_position"))
	gl.EnableVertexAttribArray(loc)
	gl.VertexAttribPointer(loc, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(0))

	loc = uint32(program.GetAttribLocation("v_normal"))
	gl.EnableVertexAttribArray(loc)
	gl.VertexAttribPointer(loc, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4))

	gl.DrawElements(mode, int32(len(indexes)), gl.UNSIGNED_BYTE, gl.Ptr(indexes))
}

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
	loc := program.GetUniformLocation("object_mat")
	gl.UniformMatrix4fv(loc, 1, false, &self.objectMat[0])

	loc = program.GetUniformLocation("view_mat")
	gl.UniformMatrix4fv(loc, 1, false, &self.viewMat[0])

	c := self.currentColor
	r := float32(c.R) / float32(255)
	g := float32(c.G) / float32(255)
	b := float32(c.B) / float32(255)
	a := float32(c.A) / float32(255)

	loc = program.GetUniformLocation("color")
	gl.Uniform4f(loc, r, g, b, a)
}

func (self *Graphics) Enable3d() {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
}

func (self *Graphics) Disable3d() {
	gl.Disable(gl.DEPTH_TEST)
}

func (self *Graphics) EnableWireframe() {
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
}

func (self *Graphics) DisableWireframe() {
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
}
