package game

import (
	"github.com/faiface/pixel/pixelgl"
)

type Level struct {
	platforms []*Platform
	walls []*Wall
}

func NewLevel(walls []*Wall, platforms []*Platform) *Level {
	return &Level{
		platforms: platforms,
		walls: walls,
	}
}

func (l *Level) Draw(win *pixelgl.Window) {
	for _, p := range l.platforms {
		p.Draw(win)
	}
	for _, w := range l.walls {
		w.Draw(win)
	}
}