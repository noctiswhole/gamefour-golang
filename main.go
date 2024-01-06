package main

import (
	"fmt"
	"os"
)

func main() {
	gameWindow := GameWindow{}
	gameGraphics := GameGraphics{}

	if err := gameWindow.Init(); err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(-1)
	}
	defer gameWindow.Destroy()

	if err := gameGraphics.Init(); err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(-1)
	}
	defer gameGraphics.Destroy()

	for !gameWindow.ShouldClose() {
		gameWindow.ProcessInput()
		gameGraphics.Draw()
		gameWindow.SwapBuffer()
	}
}
