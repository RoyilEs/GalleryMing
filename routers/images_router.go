package routers

import "github.com/gin-gonic/gin"

func (router Group) ImagesRouter() {
	router.GET("images", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "images"})
	})
}
