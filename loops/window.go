package loops

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime"
	"strconv"

	gl "github.com/go-gl/gl/v4.1-core/gl"
	glfw "github.com/go-gl/glfw/v3.1/glfw"
)

var (
	defaultWidth  = 400
	defaultHeight = 400
	defaultTitle  = "loop"
	defaultSpeed  = 1.0
	record        = false
)

var programSolid2d = ProgramSource{
	[]ShaderSource{
		ShaderSource{
			gl.VERTEX_SHADER,
			`
				#version 330

				uniform mat4 mat;
				in vec2 v_position;

				void main() {
					gl_Position = mat * vec4(v_position, 0, 1);
				}
			`,
		},
		ShaderSource{
			gl.FRAGMENT_SHADER,
			`
				#version 330

				uniform vec4 color;
				out vec4 fragColor;

				void main() {
					fragColor = color;
				}
			`,
		},
	},
}

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

func init() {
	runtime.LockOSThread()

	flag.BoolVar(&record, "record", false, "Record to gif instead of playing")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: loop [OPTIONS]\n\nOptions:\n")
		flag.PrintDefaults()
	}
}

func NewLoopWindow() *LoopWindow {
	flag.Parse()
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

	gl.ReadPixels(0, 0, int32(self.Width), int32(self.Height),
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(buffer))

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

func (self *LoopWindow) Record(graphics *Graphics) {
	numFrames := 120

	gl.Viewport(0, 0, int32(self.Width), int32(self.Height))
	gl.Clear(gl.COLOR_BUFFER_BIT)

	self.Draw(0.0, graphics)
	self.Window.SwapBuffers()

	for i := 0; i < numFrames; i++ {
		gl.Viewport(0, 0, int32(self.Width), int32(self.Height))
		gl.Clear(gl.COLOR_BUFFER_BIT)

		self.Draw(float64(i)/float64(numFrames), graphics)
		self.Window.SwapBuffers()

		self.Screenshot("frame" + strconv.Itoa(i) + ".png")
	}
}

func (self *LoopWindow) Run() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to init glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(self.Width, self.Height, self.Title, nil, nil)

	if err != nil {
		log.Fatal("Failed to create window")
	}

	self.Window = window

	window.SetKeyCallback(self.KeyCallback)
	window.MakeContextCurrent()

	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		log.Fatal("Failed to init gl")
	}

	program := NewProgram(programSolid2d.ShaderSources...)
	program.Use()

	gl.ClearColor(0.2, 0.2, 0.2, 0)

	// gl.Enable(gl.DEPTH_TEST)
	// gl.DepthFunc(gl.LESS)

	self.Load()

	CheckGLForErrors()

	graphics := NewGraphics(self, &program)
	graphics.SetMat(NewIdentityMat4())
	graphics.SetColor(color.RGBA{255, 255, 255, 255})

	if record {
		self.Record(graphics)
		log.Print("Finished")
		return
	}

	time := glfw.GetTime()
	var elapsed float64

	log.Print("Running loop")
	for !window.ShouldClose() {
		gl.Viewport(0, 0, int32(self.Width), int32(self.Height))
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

		CheckGLForErrors()

		time = newTime

		window.SwapBuffers()
		glfw.PollEvents()
	}

	log.Print("Finished")
}
