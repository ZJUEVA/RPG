package maps

import (
	"evarpg/common/renderer"

	"github.com/faiface/pixel"
)

type TiledmapComponent struct {
	renderer.ScreenPositionComponent
	renderer.BasicPictureRendererComponent
}

func (tc *TiledmapComponent) Draw(target pixel.Target) {

	panic("Not Implemented!")
}
