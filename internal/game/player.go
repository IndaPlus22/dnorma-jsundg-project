package game

import (
	"dnorma-jsundg-project/internal/input"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
	"github.com/faiface/pixel/imdraw"

)

type Player struct {
	pos pixel.Vec
	vel pixel.Vec
	sprite	*pixel.Sprite
	last	time.Time
}

func NewPlayer(sprite *pixel.Sprite, pos pixel.Vec) *Player{
	return &Player{
		sprite: sprite,
		pos:	pos,
		vel:	pixel.V(300, 300),
	}
}

//Collision check
func (p *Player) CollidingWith(rect pixel.Rect) bool {
	playerRect := pixel.R(p.pos.X, p.pos.Y, p.pos.X + 50, p.pos.Y + 50)
	return playerRect.Intersect(rect).Area() > 0
}

func (p *Player) Draw(win *pixelgl.Window) {
	if p.sprite != nil {
		p.sprite.Draw(win, pixel.IM.Moved(p.pos))
	} else {
		imd := imdraw.New(nil)
		imd.Color = pixel.RGB(0.8, 0.2, 0.2)
		imd.Push(p.pos, p.pos.Add(pixel.V(50, 50)))
		imd.Rectangle(0)
		imd.Draw(win)
	}

}

func (p *Player) Update(input *input.InputState, win *pixelgl.Window){
	dt := time.Since(p.last).Seconds()
	p.last = time.Now()

	if input.Left {
		p.pos.X -= p.vel.X * dt
	}
	if input.Right {
		p.pos.X += p.vel.X * dt
	}
	if input.Up {
		p.pos.Y += p.vel.Y * dt
	}
	if input.Down {
		p.pos.Y -= p.vel.Y * dt
	}
}

func (p *Player) GetRect() pixel.Rect {
	return pixel.R(p.pos.X, p.pos.Y, p.pos.X + 50, p.pos.Y + 50)
}