package main

import (
	"dnorma-jsundg-project/internal/game"
	"dnorma-jsundg-project/internal/input"
	"dnorma-jsundg-project/internal/levels"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	//Create window configurations
	cfg := pixelgl.WindowConfig{
		Title:  "Platformer",
		Bounds: pixel.R(0, 0, 1600, 800),
		VSync:  true,
	}
	//Create window with the configurations
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//Load levels and create game state
	level1 := levels.LoadLevel1()
	level2 := levels.LoadLevel2()
	level3 := levels.LoadLevel3()
	levels := []*game.Level{level1, level2, level3}

	gameState := game.NewGameState(level1, 1)
	currentLevel := 0

	//Create input state
	in := input.InitInputState()

	//Game loop
	for !win.Closed() {
		win.Clear(pixel.RGB(1, 0.75, 0.8))
		in.Update(win)
		gameState.UpdateGameState(in, win)

		if gameState.HasWon() {
			currentLevel++

			if currentLevel < len(levels) {
				gameState.LoadNextLevel(levels[currentLevel])
				gameState.ResetWin()
				gameState.ResetPlayer()
			} else {
				//TODO: Win screen/Menu
				break
			}
		}

		gameState.DrawGameState(win)
		win.Update()

	}
}
func main() {
	pixelgl.Run(run)
}
