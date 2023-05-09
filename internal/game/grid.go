package game

import (
	"github.com/faiface/pixel"
	"math"
)

type Grid struct {
	cellSize float64
	cells map[[2]int][]pixel.Rect
}

func NewGrid(cellSize float64) *Grid {
	return &Grid{
		cellSize: cellSize,
		cells: make(map[[2]int][]pixel.Rect),
	}
}

func (g *Grid) Add(rect pixel.Rect) {
	cellCoords := g.getCellCoords(rect)
	for _, coord := range cellCoords {
		g.cells[coord] = append(g.cells[coord], rect)
	}
}

func (g *Grid) getCellCoords(rect pixel.Rect) [][2]int {
	minX, minY := int(math.Floor(rect.Min.X / g.cellSize)), int(math.Floor(rect.Min.Y / g.cellSize))
	maxX, maxY := int(math.Floor(rect.Max.X / g.cellSize)), int(math.Floor(rect.Max.Y / g.cellSize))

	var coords [][2]int
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			coords = append(coords, [2]int{x, y})
		}
	}
	return coords
}

func (g *Grid) GetNearby(rect pixel.Rect) []pixel.Rect {
	cellCoords := g.getCellCoords(rect)
	var nearbyRects []pixel.Rect

	for _, coord := range cellCoords {
		for _, otherRect := range g.cells[coord] {
			nearbyRects = append(nearbyRects, otherRect)
		}
	}
	return nearbyRects
}