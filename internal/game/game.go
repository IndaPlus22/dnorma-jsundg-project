package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"dnorma-jsundg-project/internal/input"
)

type GameState struct {
	player		*Player
	platforms	[]*Platform //change to levels later
	playing		bool
	Input		*input.InputState
}
//Currently just a test without level design or anything
func NewGameState() *GameState {
	return &GameState{
		player:		NewPlayer(nil, pixel.V(0, 0)),
		platforms:	[]*Platform{
			NewPlatform(pixel.V(0, 0), 100, 10),
			NewPlatform(pixel.V(0, 100), 100, 10),
			NewPlatform(pixel.V(0, 200), 100, 10),
		},
	}
}

func (g *GameState) DrawGameState(win *pixelgl.Window) {
	g.player.Draw(win)
	for _, p := range g.platforms {
		p.Draw(win)
	}
}

func (g *GameState) UpdateGameState(input *input.InputState, win *pixelgl.Window) {
	g.player.Update(input, win)
	// for _, p := range g.platforms {
	// 			//For later when platforms are moving/disappearing
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