package input

import "github.com/faiface/pixel/pixelgl"

type InputState struct{
	Left		bool
	Right		bool
	Jump		bool
}

func InitInputState() *InputState {
	return &InputState{
		Left:		false,
		Right:		false,
		Jump:		false,
	}
}

func (i *InputState)UpdateInputState(win *pixelgl.Window) {
	i.Left = win.Pressed(pixelgl.KeyLeft)
	i.Right = win.Pressed(pixelgl.KeyRight)
	i.Jump = win.Pressed(pixelgl.KeySpace)
}
