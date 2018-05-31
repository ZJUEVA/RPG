package renderer

import (
	"github.com/faiface/pixel"
)

type BasicPictureRendererComponent struct {
	picture pixel.Picture
	sprite  *pixel.Sprite
	rect    pixel.Rect
}

func (b *BasicPictureRendererComponent) Picture() pixel.Picture {
	return b.picture
}

func (b *BasicPictureRendererComponent) Rect() pixel.Rect {
	return b.rect
}

func (b *BasicPictureRendererComponent) SetPicture(pic pixel.Picture) {
	b.picture = pic
	b.sprite.Set(pic, b.rect)

}

func (b *BasicPictureRendererComponent) SetRect(rect pixel.Rect) {
	b.rect = rect
	b.sprite.Set(b.picture, rect)
}

func (b *BasicPictureRendererComponent) GetRenderer() Renderer {
	return b.sprite
}

func (b *BasicPictureRendererComponent) Init(picture pixel.Picture, rect pixel.Rect) {
	b.picture = picture
	b.sprite = pixel.NewSprite(picture, rect)
}
