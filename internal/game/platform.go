package game
/*Creates and draws platforms onto the window.
*/

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Platform struct {
	rect 		pixel.Rect
	color		pixel.RGBA
	sprite 		*pixel.Sprite
}

func NewPlatform(pos pixel.Vec, width float64, height float64) *Platform {
	rect := pixel.R(pos.X, pos.Y, pos.X + width, pos.Y + height)
	return &Platform{
		rect:	rect,
		color:	pixel.RGB(0.5, 0, 0.2),
	}
}

func (p *Platform) Draw(win *pixelgl.Window){
	if p.sprite != nil {
		p.sprite.Draw(win, pixel.IM.Moved(p.rect.Min))
		return
	} else {
		imd := imdraw.New(nil)
		imd.Color = p.color
		imd.Push(p.rect.Min, p.rect.Max)
		imd.Rectangle(0)
		imd.Draw(win)
	}
}
