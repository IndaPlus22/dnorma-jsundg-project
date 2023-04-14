package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"dnorma-jsundg-project/internal/game"
	"dnorma-jsundg-project/internal/input"
)

func main()	{
	cfg := pixelgl.WindowConfig{
		Title: "Platformer",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync: true,
	}
	win, err:= pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	game := game.NewGameState()
	win.Clear(pixel.RGB(0, 0, 0))

	//Draw platforms and player
	game.DrawGameState(win)
	in := input.Input{}

	for !win.Closed() {
		
		if win.Pressed(pixelgl.KeyA) {
			in.Left = true
		}
		if win.Pressed(pixelgl.KeyD) {
			in.Right = true
		}
		if win.Pressed(pixelgl.KeySpace) {
			in.Jump = true
		}
		game.UpdateGameState(&in, win)
		game.DrawGameState(win)
		win.Update()

		//TODO: Check if player is on platform
		//TODO: Check collisions
	}
	
	//TODO: Add seperate input handler


}