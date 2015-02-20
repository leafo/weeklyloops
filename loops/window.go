package loops

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/go-gl/glh"
)

var (
	defaultWidth  = 400
	defaultHeight = 400
	defaultTitle  = "loop"
	defaultSpeed  = 1.0
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

uniform vec4 color;
out vec4 fragColor;

void main() {
	fragColor = color;
}
`

var AssertErrors func(string)

type UpdateFunc func(float64)
type DrawFunc func(float64, *Graphics)
type LoadFunc func()

type LoopWindow struct {
	Width  int
	Height int
	Update UpdateFunc
	Draw   DrawFunc
	Load   LoadFunc
	Window *glfw.Window
	Title  string
	Speed  float64
}

func errorCallback(err glfw.ErrorCode, desc string) {
	log.Print("Error: ", err, " ", desc)
}

func NewLoopWindow() *LoopWindow {
	return &LoopWindow{
		Width:  defaultWidth,
		Height: defaultHeight,
		Title:  defaultTitle,
		Speed:  defaultSpeed,
		Draw:   func(float64, *Graphics) {},
		Update: func(float64) {},
		Load:   func() {},
	}
}

func (self *LoopWindow) KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}

	if key == glfw.KeyS && action == glfw.Press {
		self.Screenshot("screen.png")
	}
}

func (self *LoopWindow) Screenshot(fname string) {
	buffer := make([]uint8, self.Width*self.Height*4)
	gl.ReadPixels(0, 0, self.Width, self.Height, gl.RGBA, gl.UNSIGNED_BYTE, buffer)

	img := image.NewRGBA(image.Rect(0, 0, self.Width, self.Height))

	for y := 0; y < self.Height; y++ {
		for x := 0; x < self.Width; x++ {
			pos := (y*self.Width + x) * 4
			pixel := color.RGBA{buffer[pos], buffer[pos+1], buffer[pos+2], buffer[pos+3]}
			img.SetRGBA(x, y, pixel)
		}
	}

	file, err := os.Create(fname)
	if err != nil {
		log.Fatal("Failed to open file: ", err.Error())
	}

	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal("Failed to write image: ", err.Error())
	}

	log.Print("Took screenshot: ", fname)
}

func (self *LoopWindow) Run() {
	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		log.Fatal("Failed to init glfw")
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(self.Width, self.Height, self.Title, nil, nil)

	if err != nil {
		log.Fatal("Failed to create window")
	}

	self.Window = window

	window.SetKeyCallback(self.KeyCallback)
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

	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()

	self.Load()

	check()

	graphics := NewGraphics(self, &program)
	graphics.SetMat(NewIdentityMat4())
	graphics.SetColor(color.RGBA{255, 255, 255, 255})


	time := glfw.GetTime()
	var elapsed float64

	log.Print("Running loop")
	for !window.ShouldClose() {
		gl.Viewport(0, 0, self.Width, self.Height)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		newTime := glfw.GetTime()
		if time != 0 {
			dt := newTime - time
			self.Update(dt)
			elapsed += dt * self.Speed

			for elapsed > 1.0 {
				elapsed -= 1.0
			}

			self.Draw(elapsed, graphics)
		}

		check()

		time = newTime

		window.SwapBuffers()
		glfw.PollEvents()
	}

	log.Print("Finished")
}
