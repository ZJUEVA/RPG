package sprite

import "evarpg/common/renderer"

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

func NewFilepathPictureComponent(filename string) *FilepathPictureComponent {
	panic("Not Implemented!")
}

type FilepathPictureSystem struct {
	filepathPictureComponents []*FilepathPictureComponent
	rendererComponents        []*renderer.BasicPictureRendererComponent
}

func (fps *FilepathPictureSystem) Add(fpc *FilepathPictureComponent, rc *renderer.BasicPictureRendererComponent) {
	panic("Not Implemented!")
}

func (fps *FilepathPictureSystem) Update() {
	panic("Not Implemented!")
}

func NewFilepathPictureSystem() *FilepathPictureSystem {
	panic("Not Implemented!")
}
