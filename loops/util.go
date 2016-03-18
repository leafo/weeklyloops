package loops

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func glErrorString(e uint32) string {
	switch e {
	case gl.NO_ERROR:
		return "NO_ERROR"
	case gl.INVALID_ENUM:
		return "INVALID_ENUM"
	case gl.INVALID_VALUE:
		return "INVALID_VALUE"
	case gl.INVALID_OPERATION:
		return "INVALID_OPERATION"
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		return "INVALID_FRAMEBUFFER_OPERATION"
	case gl.OUT_OF_MEMORY:
		return "OUT_OF_MEMORY"
	case gl.STACK_UNDERFLOW:
		return "OUT_OF_MEMORY"
	case gl.STACK_OVERFLOW:
		return "OUT_OF_MEMORY"
	}

	return "unknown error"
}

func CheckGLForErrors() {
	e := gl.GetError()
	if e != gl.NO_ERROR {
		log.Fatal("opengl error:", glErrorString(e))
	}
}
