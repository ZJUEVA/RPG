package renderer

import (
	"github.com/faiface/pixel"
)

type RendererComponent interface {
	Draw(target pixel.Target)
}
