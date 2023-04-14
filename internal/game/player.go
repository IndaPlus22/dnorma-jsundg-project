package game

import (
	"dnorma-jsundg-project/internal/input"
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
		vel:	pixel.V(0, 0),
	}
}

func (p *Player) Draw(win *pixelgl.Window) {
	p.sprite.Draw(win, pixel.IM.Moved(p.pos))
}

func (p *Player) Update(input *input.Input, win *pixelgl.Window){
	if input.Jump {
		// p.Jump()
	}
	if input.Left {
		p.MoveLeft()
	}
	if input.Right {
		p.MoveRight()
	}

	//temporary
	if p.pos.Y < 0 {
		p.pos.Y = 0
		p.onPlat = true
	}
}

// func (p *Player) Jump(){
// 	if p.CheckOnPlatform() {
// 		p.vel.Y = p.jumpStrength
// 		p.onPlat = false
// 	}
// }

func (p *Player) MoveLeft(){
	p.vel.X = -100
}

func (p *Player) MoveRight(){
	p.vel.X = 100
}

func (p *Player) PutOnPlatform(platform *Platform){
	p.pos.Y = (platform.pos.Y + platform.height)
	p.pos.X = platform.pos.X + platform.width/2 - p.sprite.Frame().W()/2
	p.vel = pixel.V(0, 0)
	p.onPlat = true
}
//Function that checks collisions between player and platforms
func (p *Player) CheckCollision(platform *Platform) bool{
	if p.pos.X + p.sprite.Frame().W() > platform.pos.X &&
		p.pos.X < platform.pos.X + platform.width &&
		p.pos.Y + p.sprite.Frame().H() > platform.pos.Y &&
		p.pos.Y < platform.pos.Y + platform.height {
			return true
		}
	return false
}

//Function that checks if player is on a platform
func (p *Player) CheckOnPlatform(platform *Platform) bool{
	if p.pos.X + p.sprite.Frame().W() > platform.pos.X &&
		p.pos.X < platform.pos.X + platform.width &&
		p.pos.Y + p.sprite.Frame().H() == platform.pos.Y {
			p.onPlat = true
			return true
		}
	p.onPlat = false
	return false
}