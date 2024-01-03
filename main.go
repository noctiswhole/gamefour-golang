package main

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"os"
)

func resizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func processInput(w *glfw.Window) {
	if w.GetKey(glfw.KeyEscape) == glfw.Press {
		w.SetShouldClose(true)
	}
}

func main() {
	_ = glfw.Init()
	defer glfw.Terminate()
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(1920, 1080, "hello golang!", nil, nil)
	if err != nil {
		os.Exit(-1)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		os.Exit(-1)
	}

	gl.Viewport(0, 0, 1920, 1080)
	window.SetFramebufferSizeCallback(resizeCallback)

	for !window.ShouldClose() {
		processInput(window)
		gl.ClearColor(1, 0, 1, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}

}
