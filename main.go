package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"dnorma-jsundg-project/internal/game"
	"dnorma-jsundg-project/internal/input"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Platformer",
		Bounds: pixel.R(0, 0, 1600, 800),
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	gameState := game.NewGameState()
	win.Clear(pixel.RGB(0, 0, 0))

	//Draw platforms and player
	gameState.DrawGameState(win)
	in := input.InitInputState()

	for !win.Closed() {
		win.Clear(pixel.RGB(0, 0, 0))
		in.Update(win)
		gameState.UpdateGameState(in, win)
		gameState.DrawGameState(win)
		win.Update()

		//TODO: Check if player is on platform
		//TODO: Check collisions
	}

	//TODO: Add seperate input handler

}
func main() {
	pixelgl.Run(run)
}
