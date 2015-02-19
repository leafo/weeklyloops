package loops

import gl "github.com/go-gl/gl"

type Graphics struct {
	*gl.Program
	*LoopWindow
}

func NewGraphics(window *LoopWindow, program *gl.Program) *Graphics {
	return &Graphics{
		Program:    program,
		LoopWindow: window,
	}
}

func (self *Graphics) SetMat(mat Mat4) {
	loc := self.Program.GetUniformLocation("mat")
	loc.UniformMatrix4f(false, (*[16]float32)(&mat))
}
