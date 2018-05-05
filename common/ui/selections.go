package ui

import (
	"github.com/faiface/pixel"
)

type SelectionComponent struct {
	selections []string
	index      int
}

func (sc *SelectionComponent) Draw(target pixel.Target) {
	panic("Not Implemented!")
}

func (sc *SelectionComponent) Index() int {
	return sc.index
}
func (sc *SelectionComponent) SetIndex(index int) {
	sc.index = index
}

func NewSelectionComponent(selections []string) SelectionComponent {
	panic("Not Implemented!")
}

type SelectionSystem struct {
	component *SelectionComponent
}

func (ss *SelectionSystem) SetComponent(component *SelectionComponent) {
	panic("Not Implemented!")
}
func (ss *SelectionSystem) UpdateComponent(component *SelectionComponent) {
	panic("Not Implemented!")
}
func (ss *SelectionSystem) Update() {
	panic("Not Implemented!")
}

func NewSelectionSystem() {
	panic("Not Implemented!")
}
