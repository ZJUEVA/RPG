package resource

import (
	"github.com/faiface/beep"
	"github.com/faiface/pixel"
	"golang.org/x/image/font"
)

type ResourceManager struct {
}

func (rm *ResourceManager) CachedLoadPicture(path string) (pixel.Picture, error) {
	return nil, nil
}

func (rm *ResourceManager) CachedLoadTTF(path string, size float64) (font.Face, error) {
	return nil, nil
}

func (rm *ResourceManager) CachedLoadMP3(path string) (beep.Streamer, error) {
	return nil, nil
}
