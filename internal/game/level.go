package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ObjectType int64

const (
	PlatformType ObjectType = iota
	WallType
)

type ObjectInfo struct {
	Rect pixel.Rect
	Type ObjectType
}

type Level struct {
	platforms []*Platform
	walls     []*Wall
	grid      *Grid
	items     []*Item
}

func NewLevel(walls []*Wall, platforms []*Platform, items []*Item) *Level {
	level := &Level{
		platforms: platforms,
		walls:     walls,
		grid:      NewGrid(250),
		items:     items,
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
	for _, i := range l.items {
		i.Draw(win)
	}
}

func (l *Level) GetNearby(rect pixel.Rect) []ObjectInfo {
	nearbyRects := l.grid.GetNearby(rect)
	var nearbyObjects []ObjectInfo

	for _, otherRect := range nearbyRects {
		objectType := WallType

		for _, p := range l.platforms {
			if p.GetRect() == otherRect {
				objectType = PlatformType
				break
			}
		}
		nearbyObjects = append(nearbyObjects, ObjectInfo{otherRect, objectType})
	}
	return nearbyObjects
}

func (l *Level) AllItemsCollected() bool {
	for _, item := range l.items {
		if item.IsActive() {
			return false
		}
	}
	return true
}