package renderer

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

type BasicTextRendererComponent struct {
	atlas *text.Atlas
	text  *text.Text
}

func (b *BasicTextRendererComponent) Atlas() *text.Atlas {
	return b.atlas
}

func (b *BasicTextRendererComponent) Text() *text.Text {
	return b.text
}

func (b *BasicTextRendererComponent) Init(atlas *text.Atlas) {
	b.text = text.New(pixel.V(0, 0), atlas)
	b.atlas = atlas
}

type BasicTextRendererSystem struct {
	BasicRendererSystem
	target pixel.Target
}

func (b *BasicTextRendererSystem) Draw(t *BasicTextRendererComponent, matrix pixel.Matrix) {
	t.text.Draw(b.target, matrix)
}
