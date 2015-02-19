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

uniform mat4 mat;
in vec2 position;

void main() {
	gl_Position = mat * vec4(position, 0, 1);
}
`

var defaultFrag = `
#version 330

out vec4 fragColor;

void main() {
	fragColor = vec4(1,1,1,1);
}
`

var AssertErrors func(string)

type UpdateFunc func(float64)
type DrawFunc func(*Graphics)
type LoadFunc func()

type LoopWindow struct {
	Width  int
	Height int
	Update UpdateFunc
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

func (self *LoopWindow) Run() {
	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		log.Fatal("Failed to init glfw")
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(self.Width, self.Height, "loop1", nil, nil)

	if err != nil {
		log.Fatal("Failed to create window")
	}

	self.Window = window

	window.SetKeyCallback(keyCallback)
	window.MakeContextCurrent()

	glfw.SwapInterval(1)

	if gl.Init() != 0 {
		log.Fatal("failed to init gl")
	}

	check := glh.OpenGLSentinel()

	AssertErrors = func(msg string) {
		log.Print(msg)
		check()
	}

	// load shaders
	vertShader := glh.Shader{gl.VERTEX_SHADER, defaultVert}
	fragShader := glh.Shader{gl.FRAGMENT_SHADER, defaultFrag}
	program := glh.NewProgram(vertShader, fragShader)
	program.Use()

	gl.ClearColor(0.2, 0.2, 0.2, 0)

	self.Load()

	check()
	log.Print("Loaded")

	positionLocation := program.GetAttribLocation("position")
	positionLocation.EnableArray()
	positionLocation.AttribPointer(2, gl.FLOAT, false, 0, nil)

	graphics := NewGraphics(self, &program)

	time := glfw.GetTime()

	log.Print("Running loop")
	for !window.ShouldClose() {
		gl.Viewport(0, 0, self.Width, self.Height)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		newTime := glfw.GetTime()
		if time != 0 {
			self.Update(newTime - time)
			self.Draw(graphics)
		}

		check()

		time = newTime

		window.SwapBuffers()
		glfw.PollEvents()
	}

	log.Print("Finished")
}
