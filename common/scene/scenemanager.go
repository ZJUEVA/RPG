package scene

import "github.com/golang-collections/collections/stack"

type SceneManager struct {
	stack stack.Stack
	now   Scene
}

func (sm *SceneManager) Push(scene Scene) {
	panic("Not Implemented!")
}

func (sm *SceneManager) Goto(scene Scene) {
	panic("Not Implemented!")
}

func (sm *SceneManager) Pop() {
	panic("Not Implemented!")
}

func (sm *SceneManager) Loop() {
	panic("Not Implemented!")
}

func NewSceneManager() *SceneManager {
	panic("Not Implemented!")
}
