package loops

import (
	gl "github.com/go-gl/gl/v4.1-core/gl"
)

var programSolid2d = ProgramSource{
	[]ShaderSource{
		ShaderSource{
			gl.VERTEX_SHADER,
			`
				#version 330

				uniform mat4 object_mat;
				uniform mat4 view_mat;

				in vec2 v_position;

				void main() {
					gl_Position = view_mat * object_mat * vec4(v_position, 0, 1);
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

				uniform mat4 object_mat;
				uniform mat4 view_mat;

				in vec2 v_position;
				in vec4 v_color;
				out vec4 f_color;

				void main() {
					f_color = v_color;
					gl_Position = view_mat * object_mat * vec4(v_position, 0, 1);
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

var programSolid3d = ProgramSource{
	[]ShaderSource{
		ShaderSource{
			gl.VERTEX_SHADER,
			`
				#version 330

				uniform mat4 object_mat;
				uniform mat4 view_mat;

				in vec3 v_position;
				in vec3 v_normal;

				out vec3 f_position;
				out vec3 f_normal;

				void main() {
					f_normal = vec3(object_mat * vec4(v_normal, 0));
					f_position = vec3(object_mat * vec4(v_position, 1));
					gl_Position = view_mat * object_mat * vec4(v_position, 1);
				}
			`,
		},
		ShaderSource{
			gl.FRAGMENT_SHADER,
			`
				#version 330
				uniform vec4 color;

				in vec3 f_normal;
				in vec3 f_position;

				out vec4 fragColor;

				void main() {
					vec3 cam = vec3(0,0,0);
					vec3 at_cam = cam - f_position;
					fragColor = vec4(color.rgb * max(0.3, dot(normalize(at_cam), f_normal)), color.a);
				}
			`,
		},
	},
}
