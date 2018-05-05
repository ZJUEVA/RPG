package scene

const (
	ActionNull = iota
	ActionPush
	ActionGoto
	ActionPop
)

type Result struct {
	Action    int
	NextScene *Scene
}
type Scene interface {
	Enter()
	Update() Result
	Leave()
}
