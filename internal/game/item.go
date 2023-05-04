package game

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type ItemType int

const	(
	SpeedBoost 	ItemType = iota
	JumpBoost	
	WinGame
)

type Item struct {
	pos    pixel.Vec
	size   float64
	color  pixel.RGBA
	active bool
	itemType ItemType
}

func NewItem(pos pixel.Vec, size float64, color pixel.RGBA, itemType ItemType) *Item {
	return &Item{
		pos:    pos,
		size:   size,
		color:  color,
		active: true,
		itemType: itemType,
	}
}

func (i *Item) Draw(win *pixelgl.Window){
	if !i.active {
		return
	}
	imd := imdraw.New(nil)
	imd.Color = i.color
	imd.Push(i.pos, i.pos.Add(pixel.V(i.size, i.size)))
	imd.Rectangle(0)
	imd.Draw(win)
}

func (i *Item) GetRect() pixel.Rect {
	return pixel.R(i.pos.X, i.pos.Y, i.pos.X+i.size, i.pos.Y+i.size)
}

func (i *Item) IsActive() bool {
	return i.active
}

func (i *Item) Deactivate(level *Level) {
	if !i.active {
		return
	}
	i.active = false
	level.RemoveItem(i)
}

func (i *Item) Collect(player *Player) {
	if !i.active {
		return
	}

	switch i.itemType {
	case SpeedBoost:
		player.SpeedBoost(2)
		fmt.Print("Speed Boost")
	case JumpBoost:
		player.JumpBoost(100)
	case WinGame:
		//TODO
	}
}