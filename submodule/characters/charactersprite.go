package characters

import (
	"evarpg/common/renderer"
	"evarpg/common/sprite"
)

type CharacterSpriteComponent struct {
	sprite.SpriteComponent
	renderer.BasicPictureRendererComponent
	direction int
	index     int
}

func (csc *CharacterSpriteComponent) Draw() {
	panic("Not Implemented!")
}
