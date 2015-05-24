package loops

import "github.com/go-gl/gl/v4.1-core/gl"

type Shape interface {
	Verts() []float32
	DrawMode() uint32
}

type Rectangle [4 * 6]float32

func NewRectangle(w, h float32) Rectangle {
	hw := w / 2.0
	hh := h / 2.0

	return Rectangle{
		-hw, -hh, 1, 1, 1, 1,
		-hw, hh, 1, 1, 1, 1,
		hw, -hh, 1, 1, 1, 1,
		hw, hh, 1, 1, 1, 1,
	}
}

func (self Rectangle) Verts() []float32 {
	return self[:]
}

func (self Rectangle) DrawMode() uint32 {
	return gl.TRIANGLE_STRIP
}
