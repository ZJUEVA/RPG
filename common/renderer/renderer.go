package renderer

import (
	"github.com/faiface/pixel"
)

type Drawer interface {
	Draw(t pixel.Target, matrix pixel.Matrix)
}

type BasicRendererSystem struct {
	Target pixel.Target
}

func (b *BasicRendererSystem) Init(target pixel.Target) {
	b.Target = target
}

var DefaultTarget pixel.Target
