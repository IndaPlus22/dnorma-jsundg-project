package game

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
)

type Level struct {
	platforms 	[]*Platform
	walls 		[]*Wall
	grid 		*Grid
}

func NewLevel(walls []*Wall, platforms []*Platform) *Level {
	level := &Level{
		platforms: platforms,
		walls: walls,
		grid: NewGrid(100),
	}
	for _, p := range platforms {
		level.grid.Add(p.GetRect())
	}
	for _, w := range walls {
		level.grid.Add(w.GetRect())
	}

	return level
}

func (l *Level) Draw(win *pixelgl.Window) {
	for _, p := range l.platforms {
		p.Draw(win)
	}
	for _, w := range l.walls {
		w.Draw(win)
	}
}

func (l *Level) GetNearby(rect pixel.Rect) []pixel.Rect {
	return l.grid.GetNearby(rect)
}