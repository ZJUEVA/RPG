package characters

type CharacterMoveComponent struct {
	speed int
}

type CharacterMoveSystem struct {
}

func (cms *CharacterMoveSystem) UpdateComponent(cmc *CharacterMoveComponent, csc *CharacterSpriteComponent) {
	panic("Not Implemented!")
}
func (cms *CharacterMoveSystem) Update() {
	panic("Not Implemented!")
}
