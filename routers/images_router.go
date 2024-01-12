package routers

import "GalleryMing/api"

func (router Group) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", imagesApi.ImageUploadView)
	router.GET("image", imagesApi.ImageGetOneView)
}
