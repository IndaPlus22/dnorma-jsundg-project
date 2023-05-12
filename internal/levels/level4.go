package levels

import (
	"dnorma-jsundg-project/internal/game"
	"github.com/faiface/pixel"
)

// Load the fourth level
func LoadLevel4() *game.Level {
	walls := []*game.Wall{
		game.NewWall(pixel.V(0, 0), 10, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(1590, 0), 10, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(0, 790), 1600, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(0, 0), 200, 700, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewWall(pixel.V(790, 0), 200, 700, pixel.RGB(0.9, 0.2, 0.4)),
	}
	platforms := []*game.Platform{
		game.NewPlatform(pixel.V(0, 0), 800, 30, pixel.RGB(0.9, 0.2, 0.4)),

		game.NewPlatform(pixel.V(205, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(250, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(295, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(340, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(385, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(430, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(475, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(520, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(565, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(610, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(655, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(700, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(745, 0), 40, 800, pixel.RGB(0.9, 0.2, 0.4)),

		game.NewPlatform(pixel.V(690, 250), 100, 10, pixel.RGB(0.9, 0.2, 0.4)),
		game.NewPlatform(pixel.V(200, 490), 100, 10, pixel.RGB(0.9, 0.2, 0.4)),
	}
	items := []*game.Item{
		game.NewItem(pixel.V(700, 100), 30, pixel.RGB(255, 255, 0), game.SpeedBoost),
		game.NewItem(pixel.V(1500, 100), 30, pixel.RGB(0, 255, 0), game.WinGame),
	}
	return game.NewLevel(walls, platforms, items)
}
