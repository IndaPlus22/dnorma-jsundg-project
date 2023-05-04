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
	g.player.Update(input, win)
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