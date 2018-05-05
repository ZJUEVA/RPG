package renderer

import "github.com/faiface/pixel"

type BasicPictureRendererComponent struct {
	ScreenPositionComponent
	picture *pixel.Picture
	sprite  *pixel.Sprite
	rect    pixel.Rect
}

func (brc *BasicPictureRendererComponent) Picture() *pixel.Picture {
	return brc.picture
}

func (brc *BasicPictureRendererComponent) Rect() pixel.Rect {
	return brc.rect
}

func (brc *BasicPictureRendererComponent) SetPicture(pic *pixel.Picture) {
	panic("Not Implemented!")
}

func (brc *BasicPictureRendererComponent) SetRect(rect *pixel.Rect) {
	panic("Not Implemented!")
}

func (brc *BasicPictureRendererComponent) Draw(target pixel.Target) {
	panic("Not Implemented!")
}

func NewBasicPictureRendererComponent(picture *pixel.Picture, screenX, screenY int, rect *pixel.Rect) *BasicPictureRendererComponent {
	panic("Not Implemented!")
}
