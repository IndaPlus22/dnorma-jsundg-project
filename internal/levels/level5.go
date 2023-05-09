package levels

import (
	"dnorma-jsundg-project/internal/game"
	"github.com/faiface/pixel"
)

// Load the first level
func LoadLevel5() *game.Level {
	walls := []*game.Wall{
		game.NewWall(pixel.V(0, 0), 10, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(1590, 0), 10, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(0, 790), 1600, 10, pixel.RGB(0.9, 0.2, 0.4)),
	}
	platforms := []*game.Platform{
		game.NewPlatform(pixel.V(0, 0), 700, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(900, 0), 700, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(110, 200), 140, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(750, 400), 100, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(1400, 600), 400, 10, pixel.RGB(0.9, 0.2, 0.4)),
	}
	items := []*game.Item{
		game.NewItem(pixel.V(100, 100), 30, pixel.RGB(255, 255, 0), game.SpeedBoost),
		game.NewItem(pixel.V(1500, 750), 30, pixel.RGB(0, 255, 0), game.WinGame),
	}
	return game.NewLevel(walls, platforms, items)
}
