package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"fmt"
)

type Player struct {
	pos pixel.Vec
	vel pixel.Vec
	sprite	*pixel.Sprite
}

func NewPlayer(sprite *pixel.Sprite, pos pixel.Vec) *Player{
	return &Player{
		sprite: sprite,
		pos:	pos,
	}
}

func (p *Player) Draw(win *pixelgl.Window) {
	p.sprite.Draw(win, pixel.IM.Moved(p.pos))
}

func (p *Player) Update(dt float64, plat *Platform, win *pixelgl.Window){
	direction := pixel.ZV
	if win.Pressed(pixelgl.KeyA){
		direction.X--
	}
	if win.Pressed(pixelgl.KeyD){
		direction.X++
	}

	p.pos = p.pos.Add(direction.Scaled(100*dt))

	if p.pos.X < plat.pos.X {
		p.pos.X = plat.pos.X
	} else if p.pos.X > plat.pos.X + plat.width {
		p.pos.X = plat.pos.X + plat.width
	}
	
}