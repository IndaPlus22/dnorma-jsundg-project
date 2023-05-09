package input
//Handles user input
import "github.com/faiface/pixel/pixelgl"

type InputState struct {
	Left  bool
	Right bool
	Up    bool
	Down  bool
}

func InitInputState() *InputState {
	return &InputState{}
}
//Updates the input state
func (i *InputState) Update(win *pixelgl.Window) {
	i.Left = win.Pressed(pixelgl.KeyLeft) || win.Pressed(pixelgl.KeyA)
	i.Right = win.Pressed(pixelgl.KeyRight) || win.Pressed(pixelgl.KeyD)
	i.Up = win.Pressed(pixelgl.KeySpace) || win.Pressed(pixelgl.KeyUp) || win.Pressed(pixelgl.KeyW)
	i.Down = win.Pressed(pixelgl.KeyDown) || win.Pressed(pixelgl.KeyS)
}
