package title

import (
	"evarpg/common/ecs"
	"evarpg/common/renderer"
	"evarpg/common/scene"
	"evarpg/common/ui"
	"evarpg/common/utils"

	"github.com/faiface/pixel"
)

type AFrameEntity struct {
	ecs.Identity
	renderer.WithBatchPictureRendererComponent
	renderer.FilepathPictureComponent
	ui.WindowFrameComponent
}

type Scene struct {
	frame_renderer ui.WindowFrameRendererSystem
	entity         AFrameEntity
}

func (s *Scene) Enter(manager *scene.SceneManager) {
	s.frame_renderer.Init(renderer.DefaultTarget)
	s.entity = AFrameEntity{Identity: ecs.NewIdentity()}

	s.entity.FilepathPictureComponent.SetFilename("assets/images/windowskin/WindowSkin.png")
	s.frame_renderer.Init(renderer.DefaultTarget)
}
func (s *Scene) Update() {
	utils.FilepathPictureSystem.Update(&s.entity.FilepathPictureComponent, &s.entity.WithBatchPictureRendererComponent)
	s.frame_renderer.Draw(&s.entity.WindowFrameComponent,
		&s.entity.WithBatchPictureRendererComponent,
		pixel.IM.Moved(utils.Window.Bounds().Center()))
}
func (s *Scene) Leave() {
	s.frame_renderer.Init(renderer.DefaultTarget)
	s.entity = AFrameEntity{Identity: ecs.NewIdentity()}

	s.entity.FilepathPictureComponent.SetFilename("assets/images/windowskin/WindowSkin.png")
	s.frame_renderer.Init(renderer.DefaultTarget)
}
