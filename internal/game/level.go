package game
//Level is the level of the game. It contains all the platforms, walls, and items.

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
)

type ObjectType int64

const (
	PlatformType ObjectType = iota
	WallType
)

type ObjectInfo struct {
	Rect 	pixel.Rect
	Type 	ObjectType
}

type Level struct {
	platforms 	[]*Platform
	walls 		[]*Wall
	grid 		*Grid
}
//NewLevel creates a new level with the given walls, platforms, and items.
func NewLevel(walls []*Wall, platforms []*Platform, items []*Item) *Level {
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
//Draw draws the level to the window.
func (l *Level) Draw(win *pixelgl.Window) {
	for _, p := range l.platforms {
		p.Draw(win)
	}
	for _, w := range l.walls {
		w.Draw(win)
	}
}
//GetNearby returns the nearby objects of the given rectangle.
func (l *Level) GetNearby(rect pixel.Rect) []ObjectInfo {
	nearbyRects := l.grid.GetNearby(rect)
	var nearbyObjects []ObjectInfo

	for _, otherRect := range nearbyRects{
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
//AllItemsCollected returns true if all the items in the level have been collected.
func (l *Level) AllItemsCollected() bool {
	for _, item := range l.items {
		if item.IsActive() {
			return false
		}
	}
	return true
}