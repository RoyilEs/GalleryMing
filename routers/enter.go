package routers

import (
	"GalleryMing/api"
	"GalleryMing/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Group struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	//网页模板
	router.LoadHTMLGlob("api/images_api/templates/*")
	router.GET(":id", api.ApiGroupApp.ImagesApi.ImageHtmlView)

	apiGroup := router.Group("datou")
	routerGroup := Group{apiGroup}
	routerGroup.ImagesRouter()
	//静态路由
	router.StaticFS("uploads", http.Dir("uploads"))
	return router
}
