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
}

func (s *Scene) Main(manager *scene.SceneManager) {
	var frame_renderer ui.WindowFrameRendererSystem
	frame_renderer.Init(renderer.DefaultTarget)

	entity := AFrameEntity{Identity: ecs.NewIdentity()}
	entity.SetWidth(600)
	entity.SetHeight(200)
	entity.FilepathPictureComponent.SetFilename("assets/images/windowskin/WindowSkin1.png")

	entity2 := AFrameEntity{Identity: ecs.NewIdentity()}
	entity2.SetWidth(300)
	entity2.SetHeight(200)
	entity2.FilepathPictureComponent.SetFilename("assets/images/windowskin/WindowSkin.png")

	frame_renderer.Init(renderer.DefaultTarget)
	i := 0
	for {
		manager.Lock()
		utils.FilepathPictureSystem.Update(&entity.FilepathPictureComponent, &entity.WithBatchPictureRendererComponent)
		frame_renderer.Draw(&entity.WindowFrameComponent,
			&entity.WithBatchPictureRendererComponent,
			pixel.IM.Moved(pixel.V(100, 100)))
		utils.FilepathPictureSystem.Update(&entity2.FilepathPictureComponent, &entity2.WithBatchPictureRendererComponent)
		frame_renderer.Draw(&entity2.WindowFrameComponent,
			&entity2.WithBatchPictureRendererComponent,
			pixel.IM.Moved(pixel.V(300, 200)))
		i++
		if i > 200 {
			manager.Goto(&Scene2{})
			manager.Yield()
			break
		}
		manager.Yield()

	}
}

type Scene2 struct {
}

func (s *Scene2) Main(manager *scene.SceneManager) {
	var frame_renderer ui.WindowFrameRendererSystem
	frame_renderer.Init(renderer.DefaultTarget)

	entity := AFrameEntity{Identity: ecs.NewIdentity()}
	entity.SetWidth(100)
	entity.SetHeight(200)
	entity.FilepathPictureComponent.SetFilename("assets/images/windowskin/WindowSkin1.png")

	frame_renderer.Init(renderer.DefaultTarget)

	for {
		manager.Lock()
		utils.FilepathPictureSystem.Update(&entity.FilepathPictureComponent, &entity.WithBatchPictureRendererComponent)
		frame_renderer.Draw(&entity.WindowFrameComponent,
			&entity.WithBatchPictureRendererComponent,
			pixel.IM.Moved(pixel.V(100, 100)))

		manager.Yield()
	}
}
