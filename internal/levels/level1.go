package levels

import (
	"github.com/faiface/pixel"
	"dnorma-jsundg-project/internal/game"
)
//Load the first level
func LoadLevel1() *game.Level {
	walls := []*game.Wall{
		game.NewWall(pixel.V(0, 0), 10, 800, pixel.RGB(0.5,0.5, 0.2)),
		game.NewWall(pixel.V(1590, 0), 10, 800, pixel.RGB(0.5, 0, 0.2)),
		// game.NewWall(pixel.V(700, 800), 500, 20, pixel.RGB(0.5,0.5, 0.2)),
		// game.NewWall(pixel.V(480, 0), 20, 500, pixel.RGB(0.5, 0, 0.2)),
	}
	platforms := []*game.Platform{
		game.NewPlatform(pixel.V(0, 0), 700, 10, pixel.RGB(0.5, 0.5, 0.2)),
		game.NewPlatform(pixel.V(900, 0), 700, 10, pixel.RGB(0, 0.5, 0)),
		game.NewPlatform(pixel.V(110, 200), 120, 10, pixel.RGB(0.5, 0.5, 1)),
		// game.NewPlatform(pixel.V(300, 300), 20, 30, pixel.RGB(0.5, 0.3, 0.2)),
	}
	items := []*game.Item{
		game.NewItem(pixel.V(100, 100), 10, pixel.RGB(0.5, 0.5, 0.5), game.SpeedBoost),
	}
	return game.NewLevel(walls, platforms, items)
}