package utils

import (
	"evarpg/common/renderer"
	"evarpg/common/resource"
)

var ResourceManager resource.ResourceManager
var FilepathPictureSystem renderer.FilepathPictureSystem

func init() {
	ResourceManager = &resource.CachedResourceManager{}
	ResourceManager.(*resource.CachedResourceManager).Init()
	FilepathPictureSystem.Init(ResourceManager)
}
