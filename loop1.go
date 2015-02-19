package main

import (
	"fmt"

	gl "github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/leafo/weeklyloops/loops"
)

func main() {
	loop := loops.NewLoopWindow()

	verts := [...]float32{-0.5, -0.5, 0.5, -0.5, 0, 0.5}

	var vertexArray gl.VertexArray
	var triangleBuffer gl.Buffer

	loop.Load = func() {
		fmt.Println("loading the thing")
		vertexArray = gl.GenVertexArray()
		vertexArray.Bind()

		triangleBuffer = gl.GenBuffer()
		triangleBuffer.Bind(gl.ARRAY_BUFFER)

		gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(verts), &verts, gl.STATIC_DRAW)
	}

	loop.Update = func(dt float64) {
		fmt.Println("Update: ", dt)
	}

	loop.Draw = func() {
		gl.DrawArrays(gl.TRIANGLE_STRIP, 0, len(verts))
	}

	loop.Run()
}
