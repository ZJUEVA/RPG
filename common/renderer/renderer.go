package renderer

import (
	"github.com/faiface/pixel"
)

type Renderer interface {
	Draw(t pixel.Target, matrix pixel.Matrix)
}

type RendererComponent interface {
	GetRenderer() Renderer
}

type CommonRendererSystem struct {
	target pixel.Target
}

func (b *CommonRendererSystem) Draw(r RendererComponent, mat pixel.Matrix) {
	r.GetRenderer().Draw(b.target, mat)
}

func (b *CommonRendererSystem) Init(target pixel.Target) {
	b.target = target
}
