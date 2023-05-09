package game
/*Creates and draws platforms onto the window.
*/

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

//Platform is a struct that contains the position, color, sprite, width, and height of the platform.
type Platform struct {
	pos 		pixel.Vec
	color		pixel.RGBA
	sprite 		*pixel.Sprite
	width 		float64
	height 		float64
}

//Creates a new platform.
func NewPlatform(pos pixel.Vec, width float64, height float64, color pixel.RGBA) *Platform {
	return &Platform{
		pos:	pos,
		width:	width,
		height:	height,
		color:	color,
	}
}

//Draw the platform onto the window.
func (p *Platform) Draw(win *pixelgl.Window){
	imd := imdraw.New(nil)
	imd.Color = p.color
	imd.Push(p.pos, p.pos.Add(pixel.V(p.width, p.height)))
	imd.Rectangle(0)
	imd.Draw(win)

	//Won't be used for optimization purposes.
	// if p.sprite != nil {
	// 	p.sprite.Draw(win, pixel.IM.Moved(p.rect.Min))
	// 	return
	// }
}
// GetRect returns the rectangle of the platform.
func (p *Platform) GetRect() pixel.Rect {
	return pixel.R(p.pos.X, p.pos.Y, p.pos.X + p.width, p.pos.Y + p.height)
}
