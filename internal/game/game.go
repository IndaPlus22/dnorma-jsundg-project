package game
// GameState is the state of the game. It contains the player, the level, and the current level number.

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
// NewGameState creates a new game state with the given levels, level, and level number.
func NewGameState(levels []*Level, level *Level, levelnum int) *GameState {
	pic, err := assets.LoadPicture("internal/assets/datalog_griffin.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())

	return &GameState{
		player:		NewPlayer(nil, pixel.V(50, 100)),
		level:		level,
		playing:	false,}

}
//DrawGameState draws the game state to the window.
func (g *GameState) DrawGameState(win *pixelgl.Window) {
	g.player.Draw(win)
	g.level.Draw(win)
}
//UpdateGameState updates the game state based on the input and window. Also checks for win conditions.
func (g *GameState) UpdateGameState(input *input.InputState, win *pixelgl.Window, levels []*Level) {
	oldPos := g.player.pos
	oldPlayerRect := g.player.GetRect()
	g.player.Update(input, win)

	playerRect := g.player.GetRect()
	nearbyObjects := g.level.GetNearby(playerRect)

	g.player.grounded = false

	// Check for platform/wall collisions
	for _, object := range nearbyObjects {
		rect := object.Rect
		if object.Type == WallType || (object.Type == PlatformType && oldPos.Y+1.0 >= rect.Max.Y) {
			intersection := playerRect.Intersect(rect)
			if intersection.Area() > 0 {
				// Determine the direction of the collision
				xOverlap := intersection.W()
				yOverlap := intersection.H()

				// If the player is inside the object, push them out
				if xOverlap < yOverlap { // Collision is horizontal
					if playerRect.Center().X > rect.Center().X { // Player is to the right
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
						g.player.vel.Y = 0
					}
				}
			}
		}

	}
	// Check for item collisions
	for _, item := range g.level.items {
		if item.IsActive() && playerRect.Intersect(item.GetRect()).Area() > 0 {
			item.Collect(g.player, g)
			item.Deactivate(g.level)
			if item.itemType == WinGame {
				g.ResetWin()
				g.NextLevel()
				if g.GetLevel() < len(levels) {
					g.LoadNextLevel(levels[g.GetLevel()])
					g.ResetPlayer()
				}
			}
		}
	}
	//If the player falls off the map, reset the level
	if g.player.pos.Y < -100 {
		g.ResetLevel(g.level)
	}
	//If the player collects all the items, win the game
	if g.level.AllItemsCollected() {
		g.WinGame()
		if g.GetLevel() < len(levels) - 1{
			g.NextLevel()
			g.LoadNextLevel((levels[g.GetLevel()]))
			g.ResetPlayer()
		}
	}
}
//HasWon returns true if the player has won the game (or finished a level)
func (g *GameState) HasWon() bool {
	return g.player.won
}
//StartGame starts the game
func (g *GameState) StartGame() {
	g.playing = true
}
//StopGame stops the game
func (g *GameState) StopGame() {
	g.playing = false
}
//IsPlaying returns true if the game is ongoing
func (g *GameState) IsPlaying() bool {
	return g.playing
}
//WinGame sets the player's won variable to true
func (g *GameState) WinGame() {
	fmt.Println("\nNice!")
	g.player.won = true
}
//ResetWin sets the player's won variable to false
func (g *GameState) ResetWin() {
	g.player.won = false
}
//LoadLevel loads the given level
func (g *GameState) LoadNextLevel(level *Level) {
	g.level = level
}
//ResetPlayer resets the player to the starting position (depends on which level the player is on)
func (g *GameState) ResetPlayer() {
	g.player.ResetEffects()
	switch g.currentLevel {
	case 0: //level1
		g.player.pos = pixel.V(50, 50)
	case 1://level2
		g.player.pos = pixel.V(400, 50)
	case 2://level3
		g.player.pos = pixel.V(50, 50)
	case 3://level4
		g.player.pos = pixel.V(50, 760)
	case 4://level5
		g.player.pos = pixel.V(1500, 750)
	}
	g.player.grounded = true
}
//ResetLevel resets the level to the given level
func (g *GameState) ResetLevel(level *Level) {
	g.level = level
	g.ResetPlayer()
	g.player.ResetEffects()
	for _, item := range g.level.items {
		item.Reset()
	}
}
//NextLevel increments the current level
func (g *GameState) NextLevel() {
    g.currentLevel++
}
//GetLevel returns the current level
func (g *GameState) GetLevel() int {
	return g.currentLevel
}