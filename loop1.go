package main

import "log"

import (
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/go-gl/glh"
)

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func main() {
	if !glfw.Init() {
		log.Fatal("Failed to init glfw")
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(400, 400, "loop1", nil, nil)
	if err != nil {
		log.Fatal("Failed to create window")
	}

	window.SetKeyCallback(keyCallback)
	window.MakeContextCurrent()

	glfw.SwapInterval(1)

	if gl.Init() != 0 {
		log.Fatal("failed to init gl")
	}

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}

	glh.OpenGLSentinel()
}
