package main

import "github.com/go-gl/gl/v4.6-core/gl"

type ShaderCompileError struct{}

func (sce *ShaderCompileError) Error() string {
	return "Could not compile shader."
}

type ShaderProgramCompileError struct{}

func (sce *ShaderProgramCompileError) Error() string {
	return "Could not build shader program."
}

type GameGraphics struct {
	program  uint32
	vao, vbo uint32
}

func (gg *GameGraphics) createBuffers(vertices []float32) (uint32, uint32) {
	var vao, vbo uint32

	gl.GenBuffers(1, &vbo)
	gl.GenVertexArrays(1, &vao)

	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return vao, vbo
}

func (gg *GameGraphics) createShader(shaderSource string, shaderType uint32) (uint32, error) {
	var status int32

	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(shaderSource)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return 0, &ShaderCompileError{}
	}

	return shader, nil
}

func (gg *GameGraphics) buildProgram(vertexShader uint32, fragShader uint32) (uint32, error) {
	var status int32

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragShader)
	gl.LinkProgram(program)

	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		return 0, &ShaderProgramCompileError{}
	}

	return program, nil
}
func (gg *GameGraphics) Init() error {
	var vertices = []float32{
		-0.5, -0.5, 0.0, // Left
		0.5, -0.5, 0.0, // Right
		0.0, 0.5, 0.0, // Top
	}
	vao, vbo := gg.createBuffers(vertices)
	gg.vbo = vbo
	gg.vao = vao

	const vertexShaderSource = `
	#version 330 core
	layout (location = 0) in vec3 position;
	void main() {
	  gl_Position = vec4(position.x, position.y, position.z, 1.0);
	}` + "\x00"

	const fragShaderSource = `
	#version 330 core
	out vec4 color;
	void main() {
	  color = vec4(1.0f, 1.0f, 0.2f, 1.0f);
	}` + "\x00"

	vertexShader, err := gg.createShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return err
	}
	defer gl.DeleteShader(vertexShader)

	fragShader, err := gg.createShader(fragShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return err
	}
	defer gl.DeleteShader(fragShader)

	gg.program, err = gg.buildProgram(vertexShader, fragShader)
	if err != nil {
		return err
	}

	return nil
}

func (gg *GameGraphics) Draw() {
	gg.Clear()

	gl.UseProgram(gg.program)
	gl.BindVertexArray(gg.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.BindVertexArray(0)
}

func (gg *GameGraphics) Clear() {
	gl.ClearColor(1, 0, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (gg *GameGraphics) Destroy() {
	gl.DeleteVertexArrays(1, &gg.vao)
	gl.DeleteBuffers(1, &gg.vbo)
	gl.DeleteProgram(gg.program)
}
