package main

import "log"

import (
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/go-gl/glh"
)

var (
	width  = 400
	height = 400
)

func errorCallback(err glfw.ErrorCode, desc string) {
	log.Print("Error: ", err, " ", desc)
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func main() {
	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		log.Fatal("Failed to init glfw")
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(width, height, "loop1", nil, nil)
	if err != nil {
		log.Fatal("Failed to create window")
	}

	window.SetKeyCallback(keyCallback)
	window.MakeContextCurrent()

	glfw.SwapInterval(1)

	if gl.Init() != 0 {
		log.Fatal("failed to init gl")
	}

	gl.ClearColor(0.2, 0.2, 0.2, 0)

	for !window.ShouldClose() {
		gl.Viewport(0, 0, width, height)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	glh.OpenGLSentinel()
}
