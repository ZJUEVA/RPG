package sprite

import "evarpg/common/renderer"

type SpriteComponent struct {
	column, row int
	index       int
	changed     bool
}

func (sc *SpriteComponent) Column() int {
	return sc.column
}

func (sc *SpriteComponent) Row() int {
	return sc.row
}

func (sc *SpriteComponent) Index() int {
	return sc.index
}

func (sc *SpriteComponent) Changed() bool {
	return sc.changed
}

func (sc *SpriteComponent) SetIndex(index int) {
	if sc.index != index {
		sc.index = index
		sc.changed = true
	}
}
func (sc *SpriteComponent) UnsetChanged() {
	sc.changed = false
}

func NewSpriteComponent(column, row int) *SpriteComponent {
	panic("Not Implemented!")
}

type SpriteSystem struct {
	spriteComponents   []*SpriteComponent
	rendererComponents []*renderer.BasicPictureRendererComponent
}

func (ss *SpriteSystem) Add(sc *SpriteComponent, rc *renderer.BasicPictureRendererComponent) {
	panic("Not Implemented!")
}

func (ss *SpriteSystem) Update() {
	panic("Not Implemented!")
}

func NewSpriteSystem() *SpriteSystem {
	panic("Not Implemented!")
}
