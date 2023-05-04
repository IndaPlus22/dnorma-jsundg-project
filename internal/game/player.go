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
	grounded bool
	jumpForce float64
	velForce float64
}

func NewPlayer(sprite *pixel.Sprite, pos pixel.Vec) *Player{
	return &Player{
		sprite: sprite,
		pos:	pos,
		vel:	pixel.V(300, 300),
		grounded: true,
		jumpForce: 200,
		velForce: 1,
	}
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

func (p *Player) Update(input *input.InputState, win *pixelgl.Window) {
	dt := time.Since(p.last).Seconds()
	p.last = time.Now()

	if input.Left {
		p.vel.X = -200 * p.velForce
	} else if input.Right {
		p.vel.X = 200 * p.velForce
	} else {
		p.vel.X = 0
	}

	if input.Up {
		p.Jump()
	}

	
	if !p.grounded {
		p.vel.Y -= 300 * dt 
	} else {
		p.vel.Y = 0
	}

	p.pos.X += p.vel.X * dt
	p.pos.Y += p.vel.Y * dt

	// fmt.Printf("Player position: %v\n", p.pos)

	// if input.Down {
	// 	p.pos.Y -= p.vel.Y * dt
	// }
}

func (p *Player) GetRect() pixel.Rect {
	return pixel.R(p.pos.X, p.pos.Y, p.pos.X + 50, p.pos.Y + 50)
}

func (p *Player) Jump() {
	if p.grounded {
		p.vel.Y += p.jumpForce
		p.grounded = false
	}
}

func (p *Player) SpeedBoost(speed float64){
	p.velForce += speed
}

func (p *Player) JumpBoost(jump float64){
	p.jumpForce += jump
}