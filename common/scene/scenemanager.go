package scene

import (
	"github.com/golang-collections/collections/stack"
)

//GOMAXPROCS
type SceneManager struct {
	stack stack.Stack
	now   Scene
}

func (s *SceneManager) Push(scene Scene) {
	s.stack.Push(s.now)
	s.now = scene
	s.now.Enter(s)
}

func (s *SceneManager) Goto(scene Scene) {
	s.now = scene
	s.now.Enter(s)
}

func (s *SceneManager) Pop() {
	s.now = s.stack.Pop().(Scene)
}

func (s *SceneManager) Update() {
	current_scene := s.now
	current_scene.Update()
	if current_scene != s.now {
		current_scene.Leave()
	}
}
