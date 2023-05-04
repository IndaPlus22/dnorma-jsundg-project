package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"dnorma-jsundg-project/internal/game"
	"dnorma-jsundg-project/internal/input"
	"dnorma-jsundg-project/internal/levels"
)

func run() {
	//Create window configurations
	cfg := pixelgl.WindowConfig{
		Title:  "Platformer",
		Bounds: pixel.R(0, 0, 1600, 800),
		VSync: true,
	}
	//Create window with the configurations
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//Load level and create game state
	level := levels.LoadLevel1()
	gameState := game.NewGameState(level)

	//Create input state
	in := input.InitInputState()

	//Game loop
	for !win.Closed() {
		win.Clear(pixel.RGB(173, 255, 230))
		in.Update(win)
		gameState.UpdateGameState(in, win)
		gameState.DrawGameState(win)
		win.Update()

	}
}
func main() {
	pixelgl.Run(run)
}
