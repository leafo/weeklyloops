package loops

import "github.com/go-gl/gl/v2.1/gl"

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

var programColored2d = ProgramSource{
	[]ShaderSource{
		ShaderSource{
			gl.VERTEX_SHADER,
			`
				#version 330

				uniform mat4 mat;
				in vec2 v_position;
				in vec4 v_color;
				out vec4 f_color;

				void main() {
					f_color = v_color;
					gl_Position = mat * vec4(v_position, 0, 1);
				}
			`,
		},
		ShaderSource{
			gl.FRAGMENT_SHADER,
			`
				#version 330
				uniform vec4 color;
				in vec4 f_color;
				out vec4 fragColor;

				void main() {
					fragColor = color * f_color;
				}
			`,
		},
	},
}