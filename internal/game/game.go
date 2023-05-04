package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"dnorma-jsundg-project/internal/input"
	// "dnorma-jsundg-project/internal/assets"
)

type GameState struct {
	player		*Player
	level 		*Level
	playing		bool

}

func NewGameState(level *Level) *GameState {
	// pic, err := assets.LoadPicture("internal/assets/player.png")
	// if err != nil {
	// 	panic(err)
	// }
	// sprite := pixel.NewSprite(pic, pic.Bounds())

	return &GameState{
		player:		NewPlayer(nil, pixel.V(50, 50)),
		level:		level,
		playing:	false,}

}

func (g *GameState) DrawGameState(win *pixelgl.Window) {
	g.player.Draw(win)
	g.level.Draw(win)
}

func (g *GameState) UpdateGameState(input *input.InputState, win *pixelgl.Window) {
    oldPos := g.player.pos
    g.player.Update(input, win)

    playerRect := g.player.GetRect()
    nearbyObjects := g.level.GetNearby(playerRect)

    g.player.grounded = false

    for _, object := range nearbyObjects {
        rect := object.Rect
        if object.Type == WallType || (object.Type == PlatformType && oldPos.Y+ 1.0 >= rect.Max.Y) {
            intersection := playerRect.Intersect(rect)
            if intersection.Area() > 0 {
                // Determine the direction of the collision
                xOverlap := intersection.W()
                yOverlap := intersection.H()

                // If the player is inside the object, push them out
                if xOverlap < yOverlap { // Collision is horizontal
                    if playerRect.Center().X > rect.Center().X{ // Player is to the right
                        g.player.pos.X += xOverlap
                    } else { // Player is to the left
                        g.player.pos.X -= xOverlap
                    }
                } else { // Collision is vertical
                    if playerRect.Center().Y > rect.Center().Y {
                        g.player.pos.Y += yOverlap
                        g.player.grounded = true
                        g.player.vel.Y = 0
                    } else {
                        g.player.pos.Y -= yOverlap
                    }
                }
            }
        } 

    }
	for _, item := range g.level.items {
		if item.IsActive() && playerRect.Intersect(item.GetRect()).Area() > 0 {
			item.Collect(g.player)
			item.Deactivate(g.level)
		}
	}
	
}


func (g *GameState) StartGame() {
	g.playing = true
}

func (g *GameState) StopGame() {
	g.playing = false
}

func (g *GameState) IsPlaying() bool {
	return g.playing
}