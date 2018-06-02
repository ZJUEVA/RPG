package main

import (
	"evarpg/common/ecs"
	"evarpg/common/renderer"
	"evarpg/common/scene"
	"evarpg/common/ui"
	"evarpg/common/utils"
	"evarpg/submodule/title"
	"fmt"
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type AFrameEntity struct {
	ecs.Identity
	renderer.BasicPictureRendererComponent
	renderer.FilepathPictureComponent
	ui.WindowFrameComponent
}

func run() {
	var manager scene.SceneManager

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	utils.Window = win
	win.Clear(colornames.Skyblue)
	buffer := pixelgl.NewCanvas(win.Bounds())
	renderer.DefaultTarget = buffer
	manager.Goto(&title.Scene{})

	fps := time.Tick(time.Second / 60)
	frames := 0
	second := time.Tick(time.Second)

	for !win.Closed() {

		manager.Update()
		win.Clear(colornames.Skyblue)
		buffer.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		win.Update()
		buffer.Clear(color.RGBA{0, 0, 0, 255})
		<-fps
		updateFPSDisplay(win, &cfg, &frames, second)

	}

}
func main() {
	pixelgl.Run(run)
}
func updateFPSDisplay(win *pixelgl.Window, cfg *pixelgl.WindowConfig, frames *int, second <-chan time.Time) {
	*frames++
	select {
	case <-second:
		win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, *frames))
		*frames = 0
	default:
	}

}
