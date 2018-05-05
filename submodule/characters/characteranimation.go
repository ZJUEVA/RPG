package characters

type CharacterAnimationComponent struct {
	frequency int
	animation bool
}

type CharacterAnimationSystem struct {
}

func (cas *CharacterAnimationSystem) UpdateComponent(cac *CharacterAnimationComponent,
	csc *CharacterSpriteComponent) {
	panic("Not Implemented!")
}

func (cas *CharacterAnimationSystem) Update() {
	panic("Not Implemented!")
}
