package renderer

import (
	"evarpg/common/resource"
)

type FilepathPictureComponent struct {
	filename string
	changed  bool
}

func (fpc *FilepathPictureComponent) Filename() string {
	return fpc.filename
}
func (fpc *FilepathPictureComponent) SetFilename(filename string) {
	if fpc.filename != filename {
		fpc.filename = filename
		fpc.changed = true
	}
}

func (fpc *FilepathPictureComponent) Changed() bool {
	return fpc.changed
}

func (fpc *FilepathPictureComponent) UnsetChanged() {
	fpc.changed = false
}

func (fpc *FilepathPictureComponent) Init(filename string) {
	fpc.filename = filename
	fpc.changed = true
}

type FilepathPictureSystem struct {
	resourceManager resource.ResourceManager
}

func (f *FilepathPictureSystem) Update(fc *FilepathPictureComponent,
	rc PictureSetterComponent) {

	if fc.Changed() {
		fc.UnsetChanged()
		img, err := f.resourceManager.LoadPicture(fc.Filename())
		if err != nil {
			panic(err)
		}
		rc.SetPicture(img)
	}
}

func (f *FilepathPictureSystem) Init(resourceManager resource.ResourceManager) {
	f.resourceManager = resourceManager
}
