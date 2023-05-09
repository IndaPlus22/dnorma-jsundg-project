package game
// Player is the player in the game. It can move around and jump.

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
// NewPlayer creates a new player with the given sprite and position.
func NewPlayer(sprite *pixel.Sprite, pos pixel.Vec) *Player {
	return &Player{
		sprite: sprite,
		pos:	pos,
		vel:	pixel.V(300, 300),
	}
}
//Draw draws the player to the window.
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
// Update updates the player's position and velocity.
func (p *Player) Update(input *input.InputState, win *pixelgl.Window) {
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

	p.pos.X += p.vel.X * dt
	p.pos.Y += p.vel.Y * dt
}
// GetRect returns the player's rectangle.
func (p *Player) GetRect() pixel.Rect {
	return pixel.R(p.pos.X, p.pos.Y, p.pos.X+40, p.pos.Y+30)
}
// Jump makes the player jump.
func (p *Player) Jump() {
	if p.grounded {
		p.vel.Y += p.jumpForce
		p.grounded = false
	}
}
// SpeedBoost increases the player's speed.
func (p *Player) SpeedBoost(speed float64) {
	p.velForce += speed
}
//JumpBoost increases the player's jump force.
func (p *Player) JumpBoost(jump float64) {
	p.jumpForce += jump
}
//ResetEffects resets the player's speed and jump force.
func (p *Player) ResetEffects() {
	p.velForce = 0
	p.jumpForce = 1000
}
