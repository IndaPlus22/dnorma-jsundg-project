package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	pos pixel.Vec
	vel pixel.Vec
	sprite	*pixel.Sprite
	onPlat 		bool
	jumpStrength float64
}

func NewPlayer(sprite *pixel.Sprite, pos pixel.Vec) *Player{
	return &Player{
		sprite: sprite,
		pos:	pos,
		onPlat: false,
		jumpStrength: 100,
	}
}

func (p *Player) Draw(win *pixelgl.Window) {
	p.sprite.Draw(win, pixel.IM.Moved(p.pos))
}

func (p *Player) Update(dt float64, plat *Platform, win *pixelgl.Window){
	p.pos = p.pos.Add(p.vel.Scaled(dt))

	//temporary
	if p.pos.Y < 0 {
		p.pos.Y = 0
		p.onPlat = true
	}
}

func (p *Player) Jump(){
	if p.onPlat {
		p.vel.Y = p.jumpStrength
		p.onPlat = false
	}
}

func (p *Player) MoveLeft(){
	p.vel.X = -100
}

func (p *Player) MoveRight(){
	p.vel.X = 100
}

func (p *Player) Stop(){
	p.vel.X = 0
}

func (p *Player) PutOnPlatform(platform *Platform){
	p.pos.Y = platform.pos.Y + platform.height
	p.onPlat = true
}

func (p *Player) CollidesWithPlatform(platform *Platform) bool{
	//TODO: implement
}