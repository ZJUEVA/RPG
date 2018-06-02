package resource

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/beep"
	"github.com/faiface/pixel"
	"golang.org/x/image/font"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func LoadTTF(path string, size float64) (font.Face, error) {
	return nil, nil
}

func LoadMP3(path string) (beep.Streamer, error) {
	return nil, nil
}
