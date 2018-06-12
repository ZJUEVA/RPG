package scene

import (
	"github.com/golang-collections/collections/stack"
)

//GOMAXPROCS
type SceneManager struct {
	stack stack.Stack
	now   Scene
	old   Scene
	ch    chan bool
	ch2   chan bool
}

func (s *SceneManager) Push(scene Scene) {
	s.stack.Push(s.now)
	s.now = scene
}

func (s *SceneManager) Goto(scene Scene) {
	s.now = scene
}

func (s *SceneManager) Pop() {
	s.now = s.stack.Pop().(Scene)
}
func (s *SceneManager) Lock() {
	<-s.ch
}
func (s *SceneManager) Yield() {
	<-s.ch
}
func (s *SceneManager) Update() {

	if s.old != s.now {
		go s.now.Main(s)
		s.old = s.now
	}
	s.ch <- true
	s.ch <- true
}

func (s *SceneManager) Init() {
	s.ch = make(chan bool)
}
