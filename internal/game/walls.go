package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Wall struct {
	pos		pixel.Vec
	width	float64
	height	float64
	color   pixel.RGBA
}

func NewWall(pos pixel.Vec, width float64, height float64, color pixel.RGBA) *Wall {
	return &Wall{
		pos:	pos,
		width:	width,
		height:	height,
		color:	color,
	}
}

func (w *Wall) Draw(win *pixelgl.Window) {
	
	imd := imdraw.New(nil)
	imd.Color = w.color
	imd.Push(w.pos, w.pos.Add(pixel.V(w.width, w.height)))
	imd.Rectangle(0)
	imd.Draw(win)
}

func (w * Wall) GetRect() pixel.Rect {
	return pixel.R(w.pos.X, w.pos.Y, w.pos.X + w.width, w.pos.Y + w.height)
}