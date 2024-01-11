package routers

import (
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

	// 静态路由
	router.StaticFS("uploads", http.Dir("uploads"))

	apiGroup := router.Group("api")
	routerGroup := Group{apiGroup}
	routerGroup.ImagesRouter()
	return router
}
