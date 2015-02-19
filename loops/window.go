package loops

import (
	"log"

	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/go-gl/glh"
)

var (
	defaultWidth  = 400
	defaultHeight = 400
)

var defaultVert = `
#version 330

in vec2 position;

void main() {
	gl_Position = vec4(position, 0, 1);
}
`

var defaultFrag = `
#version 330

out vec4 fragColor;

void main() {
	fragColor = vec4(1,1,1,1);
}
`

type DrawFunc func()
type LoadFunc func()

type LoopWindow struct {
	Width  int
	Height int
	Draw   DrawFunc
	Load   LoadFunc
	Window *glfw.Window
}

func errorCallback(err glfw.ErrorCode, desc string) {
	log.Print("Error: ", err, " ", desc)
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func NewLoopWindow() *LoopWindow {
	return &LoopWindow{
		Width:  defaultWidth,
		Height: defaultHeight,
	}
}

func (win *LoopWindow) Run() {
	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		log.Fatal("Failed to init glfw")
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(win.Width, win.Height, "loop1", nil, nil)

	if err != nil {
		log.Fatal("Failed to create window")
	}

	win.Window = window

	window.SetKeyCallback(keyCallback)
	window.MakeContextCurrent()

	glfw.SwapInterval(1)

	if gl.Init() != 0 {
		log.Fatal("failed to init gl")
	}

	// load shaders
	vertShader := glh.Shader{gl.VERTEX_SHADER, defaultVert}
	fragShader := glh.Shader{gl.FRAGMENT_SHADER, defaultFrag}
	program := glh.NewProgram(vertShader, fragShader)
	program.Use()

	gl.ClearColor(0.2, 0.2, 0.2, 0)

	win.Load()

	positionLocation := program.GetAttribLocation("position")
	positionLocation.EnableArray()
	positionLocation.AttribPointer(2, gl.FLOAT, false, 0, nil)

	for !window.ShouldClose() {
		gl.Viewport(0, 0, win.Width, win.Height)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		win.Draw()

		window.SwapBuffers()
		glfw.PollEvents()
	}

	glh.OpenGLSentinel()
}
