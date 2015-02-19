package main

import (
	"math"

	gl "github.com/go-gl/gl"
	glh "github.com/go-gl/glh"
	"github.com/leafo/weeklyloops/loops"
)

func main() {
	loop := loops.NewLoopWindow()

	verts := [...]float32{-0.5, -0.5, 0.5, -0.5, 0, 0.5}

	var vertexArray gl.VertexArray
	var triangleBuffer gl.Buffer

	var elapsed float64

	loop.Load = func() {
		vertexArray = gl.GenVertexArray()
		vertexArray.Bind()

		triangleBuffer = gl.GenBuffer()
		triangleBuffer.Bind(gl.ARRAY_BUFFER)

		gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(verts), &verts, gl.STATIC_DRAW)
	}

	loop.Update = func(dt float64) {
		elapsed += dt
	}

	loop.Draw = func(g *loops.Graphics) {
		t := loops.NewTranslateMatrix(float32(math.Sin(elapsed)), 0, 0)
		r := loops.NewRotate2DMatrix(float32(elapsed))

		g.SetMat(r.Mul(t))
		gl.DrawArrays(gl.TRIANGLE_STRIP, 0, len(verts))
	}

	loop.Run()
}
