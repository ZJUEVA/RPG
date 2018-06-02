package resource

import (
	"github.com/faiface/beep"
	"github.com/faiface/pixel"
	"github.com/hnlq715/golang-lru"
	"golang.org/x/image/font"
)

type CachedResourceManager struct {
	pictureCache *lru.ARCCache
}

func (c *CachedResourceManager) LoadPicture(path string) (pixel.Picture, error) {
	var err error
	img, exist := c.pictureCache.Get(path)
	if !exist {
		img, err = loadPicture(path)
		c.pictureCache.Add(path, img)
	}
	if err != nil {
		panic(err)
	}
	return img.(pixel.Picture), err
}

func (c *CachedResourceManager) LoadTTF(path string, size float64) (font.Face, error) {
	return nil, nil
}

func (c *CachedResourceManager) LoadMP3(path string) (beep.Streamer, error) {
	return nil, nil
}

func (c *CachedResourceManager) Init() {
	c.pictureCache, _ = lru.NewARC(100)
}
