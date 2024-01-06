package main

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

type GameWindow struct {
	window       *glfw.Window
	gameGraphics GameGraphics
}

func (g *GameWindow) initGL() error {
	if err := gl.Init(); err != nil {
		return err
	}

	gl.Viewport(0, 0, 1920, 1080)
	if err := gl.Init(); err != nil {
		return err
	}

	g.window.SetFramebufferSizeCallback(
		func(w *glfw.Window, width int, height int) { gl.Viewport(0, 0, int32(width), int32(height)) })

	return nil
}

func (g *GameWindow) initWindow() error {
	if err := glfw.Init(); err != nil {
		return err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	window, err := glfw.CreateWindow(1920, 1080, "hello golang!", nil, nil)
	if err != nil {
		return err
	}

	window.MakeContextCurrent()
	g.window = window
	return nil
}

// Init initializes the window and creates the GL context
func (g *GameWindow) Init() error {
	runtime.LockOSThread()

	if err := g.initWindow(); err != nil {
		return err
	}
	if err := g.initGL(); err != nil {
		return err
	}
	return nil
}

func (g *GameWindow) Destroy() {
	glfw.Terminate()
	g.gameGraphics.Destroy()
}

func (g *GameWindow) ProcessInput() {
	glfw.PollEvents()
	if g.window.GetKey(glfw.KeyEscape) == glfw.Press {
		g.window.SetShouldClose(true)
	}
}

func (g *GameWindow) SwapBuffer() {
	g.window.SwapBuffers()
}

func (g *GameWindow) ShouldClose() bool {
	return g.window.ShouldClose()
}
