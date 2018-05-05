package renderer

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font"
)

type BasicTextRendererComponent struct {
	ScreenPositionComponent
	face *font.Face
	text *text.Text
}

func (btrc *BasicTextRendererComponent) Face() *font.Face {
	return btrc.face
}
func (btrc *BasicTextRendererComponent) Text() *text.Text {
	return btrc.text
}

func (btrc *BasicTextRendererComponent) Draw(target pixel.Target) {
	panic("Not Implemented!")
}

func NewBasicTextRendererComponent(face *font.Face, screenX, screenY int) *BasicPictureRendererComponent {
	panic("Not Implemented!")
}

type BasicTextRendererSystem struct {
	target pixel.Target
}

func (rs *BasicTextRendererSystem) Draw(rc *BasicTextRendererComponent) {
	panic("Not Implemented!")
}
