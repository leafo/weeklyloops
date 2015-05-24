// This is adapted from https://github.com/go-gl-legacy/glh/blob/master/shader.go

package loops

import (
	"log"
	"strings"

	gl "github.com/go-gl/gl/v4.1-core/gl"
)

type ShaderSource struct {
	Type   uint32
	Source string
}

type Program struct {
	Program uint32
}

type glGetParam func(uint32, uint32, *int32)
type glGetInfoLog func(uint32, int32, *int32, *uint8)

func checkErrorLog(msg string, object uint32, statusEnum uint32, getParam glGetParam, getLog glGetInfoLog) {
	var status int32
	getParam(object, statusEnum, &status)
	if status == gl.FALSE {
		var errorMessageLength int32
		getParam(object, gl.INFO_LOG_LENGTH, &errorMessageLength)
		errorMessage := strings.Repeat("\x00", int(errorMessageLength+1))
		getLog(object, errorMessageLength, nil, gl.Str(errorMessage))
		log.Fatal(msg, errorMessage)
	}
}

func NewProgram(shaders ...ShaderSource) Program {
	program := gl.CreateProgram()

	shaderIds := make([]uint32, 0)

	for _, source := range shaders {
		shader := gl.CreateShader(source.Type)
		csource := gl.Str(source.Source + "\x00")
		gl.ShaderSource(shader, 1, &csource, nil)
		gl.CompileShader(shader)

		checkErrorLog("Failed to compile shader:",
			shader, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog)

		gl.AttachShader(program, shader)
		shaderIds = append(shaderIds, shader)
	}

	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		log.Fatal("Failed to link program")
	}

	for _, shader := range shaderIds {
		gl.DeleteShader(shader)
	}

	return Program{program}
}

func (self *Program) GetUniformLocation(name string) int32 {
	return gl.GetUniformLocation(self.Program, gl.Str(name+"\x00"))
}

func (self *Program) GetAttribLocation(name string) int32 {
	return gl.GetAttribLocation(self.Program, gl.Str(name+"\x00"))
}

func (self *Program) Use() {
	gl.UseProgram(self.Program)
}
