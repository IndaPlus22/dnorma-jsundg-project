package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"dnorma-jsundg-project/internal/input"
	// "dnorma-jsundg-project/internal/assets"
)

type GameState struct {
	player		*Player
	level 		*Level
	playing		bool

}

func NewGameState(level *Level) *GameState {
	// pic, err := assets.LoadPicture("internal/assets/player.png")
	// if err != nil {
	// 	panic(err)
	// }
	// sprite := pixel.NewSprite(pic, pic.Bounds())

	return &GameState{
		player:		NewPlayer(nil, pixel.V(50, 100)),
		level:		level,
		playing:	false,}

}

func (g *GameState) DrawGameState(win *pixelgl.Window) {
	g.player.Draw(win)
	g.level.Draw(win)
}

func (g *GameState) UpdateGameState(input *input.InputState, win *pixelgl.Window) {
	oldPos := g.player.pos
	oldPlayerRect := g.player.GetRect()
	g.player.Update(input, win)

	playerRect := g.player.GetRect()
	nearbyObjects := g.level.GetNearby(playerRect)

	for _, object := range nearbyObjects {
		rect:= object.Rect
		if object.Type == WallType || (object.Type == PlatformType && oldPlayerRect.Min.Y >= rect.Max.Y) {
			if playerRect.Intersect(rect).Area() > 0 {
			g.player.pos = oldPos
			break
			}
		}
	}
	// for _, platform := range g.level.platforms {
	// 	if g.player.CollidingWith(platform.GetRect()) && oldPos.Y >= platform.GetRect().Max.Y {
	// 		g.player.pos = oldPos
	// 		break
	// 	}
	// }
	// for _, wall := range g.level.walls {
	// 	if g.player.CollidingWith(wall.GetRect()) {
	// 		g.player.pos = oldPos
	// 		break
	// 	}
	// }
}

func (g *GameState) StartGame() {
	g.playing = true
}

func (g *GameState) StopGame() {
	g.playing = false
}

func (g *GameState) IsPlaying() bool {
	return g.playing
}