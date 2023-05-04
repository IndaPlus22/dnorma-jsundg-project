package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"dnorma-jsundg-project/internal/input"
	// "dnorma-jsundg-project/internal/assets"
)

type GameState struct {
	player		*Player
	platforms	[]*Platform //change to levels later
	playing		bool
	Input		*input.Input
	walls		[]*Wall

}
//Currently just a test without level design or anything
func NewGameState() *GameState {
	// pic, err := assets.LoadPicture("internal/assets/player.png")
	// if err != nil {
	// 	panic(err)
	// }
	// sprite := pixel.NewSprite(pic, pic.Bounds())

	return &GameState{
		player:		NewPlayer(nil, pixel.V(50, 100)),
		platforms:	[]*Platform{
			NewPlatform(pixel.V(300, 500), 200, 20),
			NewPlatform(pixel.V(100, 100), 100, 30),
			NewPlatform(pixel.V(400, 200), 100, 10),
		},
    
		walls:		[]*Wall{
			NewWall(pixel.V(0, 0), 10, 100, pixel.RGB(0, 0, 0)),
			NewWall(pixel.V(100, 0), 10, 100, pixel.RGB(0, 0, 0)),
			NewWall(pixel.V(200, 0), 10, 100, pixel.RGB(0, 0, 0)),
		},
		playing:	false,}

}

func (g *GameState) DrawGameState(win *pixelgl.Window) {
	g.player.Draw(win)
	for _, p := range g.platforms {
		p.Draw(win)
	}
	for _, w := range g.walls {
		w.Draw(win)
	}
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