package ui

import (
	"evarpg/common/renderer"

	"github.com/faiface/pixel"
)

type WindowFrameComponent struct {
	width, height int
	changed       bool
}

func (w *WindowFrameComponent) Width() int {
	return w.width
}
func (w *WindowFrameComponent) Height() int {
	return w.height
}
func (w *WindowFrameComponent) SetWidth(width int) {
	panic("Not implemented!")
}
func (w *WindowFrameComponent) SetHeight(width int) {
	panic("Not implemented!")
}

type WindowFrameRendererSystem struct {
	renderer.WithBatchPictureRendererSystem
}

func (b *WindowFrameRendererSystem) Draw(w *WindowFrameComponent,
	pic *renderer.WithBatchPictureRendererComponent, matrix pixel.Matrix) {

	b.WithBatchPictureRendererSystem.Draw(pic, matrix)
	b.DrawAll(pic, pixel.IM)
}

func (b *WindowFrameRendererSystem) Init(target pixel.Target) {
	b.WithBatchPictureRendererSystem.Init(target)
}
