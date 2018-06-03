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
	entity2        AFrameEntity
	entity3        AFrameEntity
}

func (s *Scene) Enter(manager *scene.SceneManager) {
	s.frame_renderer.Init(renderer.DefaultTarget)

	s.entity = AFrameEntity{Identity: ecs.NewIdentity()}
	s.entity.SetWidth(600)
	s.entity.SetHeight(200)
	s.entity.FilepathPictureComponent.SetFilename("assets/images/windowskin/WindowSkin1.png")

	s.entity2 = AFrameEntity{Identity: ecs.NewIdentity()}
	s.entity2.SetWidth(300)
	s.entity2.SetHeight(200)
	s.entity2.FilepathPictureComponent.SetFilename("assets/images/windowskin/WindowSkin.png")

	s.entity3 = AFrameEntity{Identity: ecs.NewIdentity()}
	s.entity3.SetWidth(100)
	s.entity3.SetHeight(200)
	s.entity3.FilepathPictureComponent.SetFilename("assets/images/windowskin/WindowSkin3.png")
	s.frame_renderer.Init(renderer.DefaultTarget)
}
func (s *Scene) Update() {
	utils.FilepathPictureSystem.Update(&s.entity.FilepathPictureComponent, &s.entity.WithBatchPictureRendererComponent)
	s.frame_renderer.Draw(&s.entity.WindowFrameComponent,
		&s.entity.WithBatchPictureRendererComponent,
		pixel.IM.Moved(pixel.V(100, 100)))
	utils.FilepathPictureSystem.Update(&s.entity2.FilepathPictureComponent, &s.entity2.WithBatchPictureRendererComponent)
	s.frame_renderer.Draw(&s.entity2.WindowFrameComponent,
		&s.entity2.WithBatchPictureRendererComponent,
		pixel.IM.Moved(pixel.V(300, 200)))
	utils.FilepathPictureSystem.Update(&s.entity3.FilepathPictureComponent, &s.entity3.WithBatchPictureRendererComponent)
	s.frame_renderer.Draw(&s.entity3.WindowFrameComponent,
		&s.entity3.WithBatchPictureRendererComponent,
		pixel.IM.Moved(pixel.V(10, 400)))
}
func (s *Scene) Leave() {

}
