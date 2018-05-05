package renderer

type ScreenPositionComponent struct {
	screenX int
	screenY int
}

func (spc *ScreenPositionComponent) ScreenX() int {
	return spc.screenX
}
func (spc *ScreenPositionComponent) ScreenY() int {
	return spc.screenY
}
func (spc *ScreenPositionComponent) SetScreenX(x int) {
	panic("Not Implemented!")
}
func (spc *ScreenPositionComponent) SetScreenY(y int) {
	panic("Not Implemented!")
}
