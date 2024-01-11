package models

import (
	"GalleryMing/models/ctype"
	"gorm.io/gorm"
)

type ImageModel struct {
	gorm.Model
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `json:"name"`
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片类型 本地 or 七牛云
}
