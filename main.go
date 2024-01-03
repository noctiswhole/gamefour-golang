package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"os"
	"runtime"
)

type Game struct {
	window *glfw.Window
}

func (g *Game) InitGL() error {
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

func (g *Game) InitWindow() error {
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

func (g *Game) Init() error {
	runtime.LockOSThread()

	if err := g.InitWindow(); err != nil {
		return err
	}
	if err := g.InitGL(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Destroy() {
	glfw.Terminate()
}

func (g *Game) ProcessInput() {
	glfw.PollEvents()
	if g.window.GetKey(glfw.KeyEscape) == glfw.Press {
		g.window.SetShouldClose(true)
	}
}

func (g *Game) ClearBuffer() {
	gl.ClearColor(1, 0, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (g *Game) SwapBuffer() {
	g.window.SwapBuffers()
}

func (g *Game) ShouldClose() bool {
	return g.window.ShouldClose()
}

func main() {
	game := Game{}

	if err := game.Init(); err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(-1)
	}

	defer game.Destroy()

	for !game.ShouldClose() {
		game.ProcessInput()
		game.ClearBuffer()
		game.SwapBuffer()
	}

}
