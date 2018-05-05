package ui

import (
	"evarpg/common/renderer"

	"github.com/faiface/pixel"
)

type WindowFrameComponent struct {
	width, height int
	changed       bool
}

func (wfc *WindowFrameComponent) Width() int {
	return wfc.width
}
func (wfc *WindowFrameComponent) Height() int {
	return wfc.height
}
func (wfc *WindowFrameComponent) SetWidth(width int) {
	panic("Not implemented!")
}
func (wfc *WindowFrameComponent) SetHeight(width int) {
	panic("Not implemented!")
}

type WindowFrameRendererSystem struct {
	target pixel.Target
}

func (rs *WindowFrameRendererSystem) Draw(rc *renderer.BasicPictureRendererComponent, wfc *WindowFrameComponent) {
	panic("Not Implemented!")
}
