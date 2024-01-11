package api

import "GalleryMing/api/images_api"

type ApiGroup struct {
	ImagesApi images_api.ImageApi
}

var ApiGroupApp = new(ApiGroup)
