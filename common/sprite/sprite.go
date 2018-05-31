package sprite

import (
	"evarpg/common/renderer"

	"github.com/faiface/pixel"
)

type SpriteComponent struct {
	column, row int
	index       int
	changed     bool
}

func (s *SpriteComponent) Column() int {
	return s.column
}

func (s *SpriteComponent) Row() int {
	return s.row
}

func (s *SpriteComponent) Index() int {
	return s.index
}

func (s *SpriteComponent) Changed() bool {
	return s.changed
}

func (s *SpriteComponent) SetIndex(index int) {
	if s.index != index {
		s.index = index
		s.changed = true
	}
}

func (s *SpriteComponent) UnsetChanged() {
	s.changed = false
}

func (s *SpriteComponent) Init(column, row int) {
	s.column = column
	s.row = row
	s.index = 0
	s.changed = true
}

type SpriteSystem struct {
}

func (s *SpriteSystem) Update(sprite *SpriteComponent, rc *renderer.BasicPictureRendererComponent) {
	bound := rc.Picture().Bounds()
	w, h := bound.W(), bound.H()
	ry, rx := float64(sprite.Index()/sprite.Column()), float64(sprite.Index()%sprite.Column())
	cw, ch := w/float64(sprite.Column()), h/float64(sprite.Row())

	rc.SetRect(pixel.R(rx*cw, ry*ch, (rx+1)*cw, (ry+1)*ch))
}
