package main

//Entry point of the program
import (
	"dnorma-jsundg-project/internal/game"
	"dnorma-jsundg-project/internal/input"
	"dnorma-jsundg-project/internal/levels"
	"fmt"
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)
//Draws the win screen
func DrawWinScreen(win *pixelgl.Window) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	winText := text.New(pixel.V(win.Bounds().Center().X-50, win.Bounds().Center().Y), basicAtlas)
	winText.Color = pixel.RGB(0, 0, 0)

	_, _ = fmt.Fprint(winText, "You Win!")
	winText.Draw(win, pixel.IM.Scaled(winText.Orig, 4))
}

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
	level4 := levels.LoadLevel4()
	level5 := levels.LoadLevel5()
	levels := []*game.Level{level1, level2, level3, level4, level5}
	gameState := game.NewGameState(levels, level1, 0)

	//Create input state
	in := input.InitInputState()

	//Game loop
	for !win.Closed() {
		win.Clear(pixel.RGB(1, 0.75, 0.8))
		in.Update(win)
		gameState.UpdateGameState(in, win, levels)
		if gameState.HasWon() && gameState.GetLevel() == len(levels) { //If the player has won and is on the last level
			DrawWinScreen(win)
			win.Update()
			time.Sleep(5 * time.Second)
			break

		} else{ //If the player has not won or is not on the last level
			gameState.DrawGameState(win)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
