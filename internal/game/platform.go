package game
/*Creates and draws platforms onto the window.
*/

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Platform struct {
	pos		pixel.Vec
	width	float64
	height	float64
	color   pixel.RGBA
}

func NewPlatform(pos pixel.Vec, width float64, height float64) *Platform {
	return &Platform{
		pos:	pos,
		width:	width,
		height:	height,
	}
}

func (p *Platform) Draw(win *pixelgl.Window){
	imd := imdraw.New(nil)
	imd.Color = p.color
	imd.Push(p.pos)
	imd.Push(p.pos.Add(pixel.V(p.width, 0)))
	imd.Push(p.pos.Add(pixel.V(p.width, p.height)))
	imd.Push(p.pos.Add(pixel.V(0, p.height)))
	imd.Polygon(0)
	imd.Draw(win)
}
