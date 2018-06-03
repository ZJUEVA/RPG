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
	w.width = width
	w.changed = true
}
func (w *WindowFrameComponent) SetHeight(height int) {
	w.height = height
	w.changed = true
}
func (w *WindowFrameComponent) Unchange() {
	w.changed = false
}
func (w *WindowFrameComponent) Changed() bool {
	return w.changed
}

type WindowFrameRendererSystem struct {
	renderer.WithBatchPictureRendererSystem
}

func setRect(pic *renderer.WithBatchPictureRendererComponent, x, y, w, h float64) {
	pic.SetRect(pixel.R(x, y, x+w, y+h))
}

func (b *WindowFrameRendererSystem) Draw(w *WindowFrameComponent,
	pic *renderer.WithBatchPictureRendererComponent, matrix pixel.Matrix) {
	if w.Changed() {
		width := w.Width()
		height := w.Height()
		pic.Batch().Clear()
		setRect(pic, 0, 0, 128, 128)
		b.WithBatchPictureRendererSystem.Draw(pic,
			pixel.IM.
				ScaledXY(pixel.V(0, 0), pixel.V(float64(width-8)/128, float64(height-8)/128)).
				Chained(matrix).
				Moved(pixel.V(float64(width-8)/2+4, float64(height-8)/2+4)))
		matrix = matrix.Moved(pixel.V(8, 8))

		//横向填充·下
		for i := 0; i < width/16-1; i += 2 {
			setRect(pic, 128+16, 64, 16, 16)
			b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(float64(i)*16+16, 0)))
		}
		for i := 1; i < width/16-1; i += 2 {
			setRect(pic, 128+32, 64, 16, 16)
			b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(float64(i)*16+16, 0)))
		}

		//纵向填充·右
		for i := 0; i < height/16-1; i += 2 {
			setRect(pic, 128+48, 64+16, 16, 16)
			b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(float64(width-16), float64(i)*16+16)))
		}
		for i := 1; i < height/16-1; i += 2 {
			setRect(pic, 128+48, 64+32, 16, 16)
			b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(float64(width-16), float64(i)*16+16)))
		}

		//横向填充·上
		for i := 0; i < width/16-1; i += 2 {
			setRect(pic, 128+16, 64+48, 16, 16)
			b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(float64(i)*16+16, float64(height-16))))
		}
		for i := 1; i < width/16-1; i += 2 {
			setRect(pic, 128+32, 64+48, 16, 16)
			b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(float64(i)*16+16, float64(height-16))))
		}

		//纵向填充·左
		for i := 0; i < height/16-1; i += 2 {
			setRect(pic, 128, 64+16, 16, 16)
			b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(0, float64(i)*16+16)))
		}
		for i := 1; i < height/16-1; i += 2 {
			setRect(pic, 128, 64+32, 16, 16)
			b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(0, float64(i)*16+16)))
		}
		// 左下
		setRect(pic, 128, 64, 16, 16)
		b.WithBatchPictureRendererSystem.Draw(pic, matrix)
		// 右下
		setRect(pic, 128+48, 64, 16, 16)
		b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(float64(width-16), 0)))
		// 左上
		setRect(pic, 128, 64+48, 16, 16)
		b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(0, float64(height-16))))
		// 右上
		setRect(pic, 128+48, 64+48, 16, 16)
		b.WithBatchPictureRendererSystem.Draw(pic, matrix.Moved(pixel.V(float64(width-16), float64(height-16))))
		w.Unchange()
	}

	b.DrawAll(pic, pixel.IM)
}

func (b *WindowFrameRendererSystem) Init(target pixel.Target) {
	b.WithBatchPictureRendererSystem.Init(target)
}
