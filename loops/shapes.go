package loops

import (
	"image/color"
	"log"
	"math"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shape interface {
	Verts() []float32
	DrawMode() uint32
}

type Rectangle [4 * 6]float32

func pushVec2(verts []float32, vec Vec2) []float32 {
	return append(verts, vec[0], vec[1])
}

func pushColor(verts []float32, c color.RGBA) []float32 {
	return append(verts,
		float32(c.R)/float32(255),
		float32(c.G)/float32(255),
		float32(c.B)/float32(255),
		float32(c.A)/float32(255))
}

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

type Circle struct {
	verts []float32
}

func NewCircle(radius float32, parts int) Circle {
	if parts < 3 {
		log.Fatal("circle needs 3 parts")
	}

	dTheta := math.Pi * 2.0 / float32(parts)
	verts := []float32{}
	verts = pushVec2(verts, NewVec2(0.0, 0.0))
	verts = pushColor(verts, color.RGBA{255, 255, 255, 255})

	p := NewVec2(0.0, 1.0)
	for i := 0; i < parts; i++ {
		verts = pushVec2(verts, p)
		verts = pushColor(verts, color.RGBA{255, 255, 255, 255})
		p = p.Rotate(dTheta)
	}

	verts = pushVec2(verts, NewVec2(0.0, 1.0))
	verts = pushColor(verts, color.RGBA{255, 255, 255, 255})

	return Circle{verts}
}

func (self Circle) Verts() []float32 {
	return self.verts
}

func (self Circle) DrawMode() uint32 {
	return gl.TRIANGLE_FAN
}
