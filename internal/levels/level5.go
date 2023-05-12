package levels

import (
	"dnorma-jsundg-project/internal/game"
	"github.com/faiface/pixel"
)

// Load the fifth level
func LoadLevel5() *game.Level {
	walls := []*game.Wall{
		game.NewWall(pixel.V(0, 0), 10, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(1590, 0), 10, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(0, 790), 800, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(1400, 80), 10, 700, pixel.RGB(0.9, 0.2, 0.4)),
	}
	platforms := []*game.Platform{
		game.NewPlatform(pixel.V(1395, 0), 205, 30, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(1200, 0), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(1000, 0), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(800, 0), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(600, 0), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(400, 0), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(200, 10), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),

		game.NewPlatform(pixel.V(1200, 260), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(1000, 250), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(800, 250), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(600, 250), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(400, 250), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(200, 250), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),

		game.NewPlatform(pixel.V(1200, 500), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(1000, 500), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(800, 500), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(600, 500), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(400, 500), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(200, 510), 10, 10, pixel.RGB(0.9, 0.2, 0.4)),
	}
	items := []*game.Item{
		game.NewItem(pixel.V(1400, 10), 40, pixel.RGB(255, 255, 0), game.SpeedBoost),
		game.NewItem(pixel.V(10, 750), 30, pixel.RGB(0, 255, 0), game.WinGame),
	}
	return game.NewLevel(walls, platforms, items)
}
