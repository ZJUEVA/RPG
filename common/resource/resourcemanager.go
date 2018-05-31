package resource

import (
	"github.com/faiface/beep"
	"github.com/faiface/pixel"
	"golang.org/x/image/font"
)

type ResourceManager interface {
	LoadPicture(path string) (pixel.Picture, error)
	LoadTTF(path string, size float64) (font.Face, error)
	LoadMP3(path string) (beep.Streamer, error)
}
