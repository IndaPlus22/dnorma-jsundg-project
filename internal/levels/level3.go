package levels

import (
	"dnorma-jsundg-project/internal/game"
	"github.com/faiface/pixel"
)

// Load the third level
func LoadLevel3() *game.Level {
	walls := []*game.Wall{
		game.NewWall(pixel.V(0, 0), 10, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(1590, 0), 10, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(0, 790), 1600, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(100, 0), 60, 600, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(260, 200), 60, 600, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(420, 0), 60, 600, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(420, 700), 60, 100, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(580, 200), 60, 600, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(740, 0), 60, 600, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(1100, 0), 60, 600, pixel.RGB(0.9, 0.2, 0.4)),
	}
	platforms := []*game.Platform{
		game.NewPlatform(pixel.V(0, 0), 800, 30, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(1100, 0), 700, 30, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(640, 200), 100, 10, pixel.RGB(0.5, 0.5, 1)),
		game.NewPlatform(pixel.V(640, 400), 100, 10, pixel.RGB(0.5, 0.5, 1)),
		game.NewPlatform(pixel.V(1095, 515), 5, 10, pixel.RGB(1, 0.75, 0.8)),
	}
	items := []*game.Item{
		game.NewItem(pixel.V(50, 200), 30, pixel.RGB(1, 0, 1) /*Pink*/, game.JumpBoost),
		game.NewItem(pixel.V(480, 200), 100, pixel.RGB(0.9, 0.75, 0.9), game.ResetEffect),
		game.NewItem(pixel.V(1500, 100), 30, pixel.RGB(0, 255, 0), game.WinGame),
	}
	return game.NewLevel(walls, platforms, items)
}
