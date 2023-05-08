package levels

import (
	"dnorma-jsundg-project/internal/game"
	"github.com/faiface/pixel"
)

// Load the second level
func LoadLevel2() *game.Level {
	walls := []*game.Wall{
		game.NewWall(pixel.V(0, 0), 10, 800, pixel.RGB(0.5, 0.5, 0.2)),
		game.NewWall(pixel.V(1590, 0), 10, 800, pixel.RGB(0.5, 0.5, 0.2)),
		game.NewWall(pixel.V(0, 790), 1600, 10, pixel.RGB(0.5, 0.5, 0.2)),
	}
	platforms := []*game.Platform{
		game.NewPlatform(pixel.V(0, 0), 700, 10, pixel.RGB(0, 0, 0)),
		game.NewPlatform(pixel.V(900, 0), 700, 10, pixel.RGB(0, 0, 0)),
	}
	items := []*game.Item{
		game.NewItem(pixel.V(1500, 100), 30, pixel.RGB(1, 0, 1) /*Pink*/, game.JumpBoost),
		game.NewItem(pixel.V(40, 740), 30, pixel.RGB(0, 255, 0), game.WinGame),
	}
	return game.NewLevel(walls, platforms, items)
}
