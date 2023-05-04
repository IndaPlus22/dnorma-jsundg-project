package levels

import (
	"github.com/faiface/pixel"
	"dnorma-jsundg-project/internal/game"
)
//Load the first level
func LoadLevel1() *game.Level {
	walls := []*game.Wall{
		game.NewWall(pixel.V(0, 0), 20, 500, pixel.RGB(0.5,0.5, 0.2)),
		game.NewWall(pixel.V(0, 0),  500, 20, pixel.RGB(0.5, 1, 0.2)),
		game.NewWall(pixel.V(0, 480), 500, 20, pixel.RGB(0.5,0.5, 0.2)),
		game.NewWall(pixel.V(480, 0), 20, 500, pixel.RGB(0.5, 0, 0.2)),
	}
	platforms := []*game.Platform{
		game.NewPlatform(pixel.V(30, 50), 400, 10, pixel.RGB(0.5, 0.5, 0.2)),
		game.NewPlatform(pixel.V(100, 90), 20, 30, pixel.RGB(0, 0.5, 0)),
		game.NewPlatform(pixel.V(200, 200), 20, 30, pixel.RGB(0.5, 0.5, 1)),
		game.NewPlatform(pixel.V(300, 300), 20, 30, pixel.RGB(0.5, 0.3, 0.2)),
	}
	return game.NewLevel(walls, platforms)
}