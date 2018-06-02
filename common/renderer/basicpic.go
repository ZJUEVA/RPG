package renderer

import (
	"github.com/faiface/pixel"
)

type PictureSetterComponent interface {
	SetPicture(pic pixel.Picture)
}

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
	if b.sprite == nil {
		b.rect = b.picture.Bounds()
		b.sprite = pixel.NewSprite(b.picture, b.rect)
	}

	b.sprite.Set(pic, b.rect)
}

func (b *BasicPictureRendererComponent) SetRect(rect pixel.Rect) {
	b.rect = rect
	b.sprite.Set(b.picture, rect)
}

type BasicPictureRendererSystem struct {
	BasicRendererSystem
}

func (b *BasicPictureRendererSystem) Draw(pic *BasicPictureRendererComponent, matrix pixel.Matrix) {
	if pic.sprite != nil {
		pic.sprite.Draw(b.Target, matrix)
	}
}

type WithBatchPictureRendererComponent struct {
	BasicPictureRendererComponent
	batch *pixel.Batch
}

func (b *WithBatchPictureRendererComponent) Batch() *pixel.Batch {
	return b.batch
}

func (b *WithBatchPictureRendererComponent) SetPicture(pic pixel.Picture) {
	if pic != b.picture {
		b.batch = pixel.NewBatch(&pixel.TrianglesData{}, pic)
		b.BasicPictureRendererComponent.SetPicture(pic)
	}
}

type WithBatchPictureRendererSystem struct {
	BasicRendererSystem
}

func (b *WithBatchPictureRendererSystem) Draw(pic *WithBatchPictureRendererComponent, matrix pixel.Matrix) {
	if pic.sprite != nil && pic.batch != nil {
		pic.sprite.Draw(pic.batch, matrix)
	}
}
func (b *WithBatchPictureRendererSystem) DrawAll(pic *WithBatchPictureRendererComponent, matrix pixel.Matrix) {
	if pic.batch != nil {
		pic.batch.Draw(b.Target)
	}
}
