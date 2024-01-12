package images_api

import (
	"GalleryMing/global"
	"GalleryMing/models"
	"GalleryMing/models/res"
	"GalleryMing/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func (ImageApi) ImageGetOneView(c *gin.Context) {
	var (
		ImageModel models.ImageModel
		ids        []uint
	)
	//TODO 数据量过大需要分批处理
	err := global.DB.Table("image_models").Pluck("id", &ids).Error
	if err != nil {
		global.Log.Error("ID列表获取失败")
	}

	r := utils.Random(ids)
	err = global.DB.Where("id = ?", r).Find(&ImageModel).Error
	fmt.Println(ImageModel.Path)
	if err != nil {
		res.ResultFailWithMsg("获取失败", c)
		return
	}
	query, b := c.GetQuery("type")
	if b && query == "json" {
		s := c.Request.Host
		res.ResultOkWithData("http://"+s+ImageModel.Path, c)
		return
	}
	nameList := strings.Split(ImageModel.Name, ".")
	data, err := ioutil.ReadFile("." + ImageModel.Path)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.Data(http.StatusOK, "image/"+nameList[len(nameList)-1], data)
}
