package images_api

import (
	"GalleryMing/models/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ImageApi) ImageHtmlView(c *gin.Context) {
	id := c.Param("id")
	if id == "永明" {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	} else {
		res.ResultFailWithMsg("需要神秘代码", c)
	}
}
