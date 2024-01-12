package images_api

import (
	"GalleryMing/global"
	"GalleryMing/models"
	"GalleryMing/models/res"
	"GalleryMing/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"strings"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// ImageUploadView 上传图片返回url
func (ImageApi) ImageUploadView(c *gin.Context) {
	// 多张上传
	form, err := c.MultipartForm()
	if err != nil {
		res.ResultFailWithMsg(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.ResultFailWithMsg("没有上传文件", c)
		return
	}
	// 判断路径是否存在
	basePath := global.Config.Upload.Path
	baseSize := global.Config.Upload.Size
	_, err = os.ReadDir(basePath)
	if err != nil {
		global.Log.Error(err)
	}
	var resList []FileUploadResponse
	// 遍历逐个获取
	for _, file := range fileList {

		filename := file.Filename
		nameList := strings.Split(filename, ".")
		// 判断后缀
		if !utils.InList(strings.ToLower(nameList[len(nameList)-1]), global.WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "不支持的文件类型",
			})
			continue
		}
		// 判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(baseSize) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("文件大小不能超过%dM, 当前大小为%.2fM", baseSize, size),
			})
			continue
		}

		//获取文件对象
		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, _ := io.ReadAll(fileObj)
		imageHash := utils.Md5(byteData)
		//去数据库查询这个图片是否存在
		var imageModel models.ImageModel
		err = global.DB.Take(&imageModel, "hash = ?", imageHash).Error
		if err == nil {
			resList = append(resList, FileUploadResponse{
				FileName:  imageModel.Path,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}

		filePath := path.Join(basePath, filename)
		//存储 本地
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "图片上传失败" + err.Error(),
			})
			continue
		}
		//上传成功

		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "图片上传成功",
		})
		//图片入库
		global.DB.Create(&models.ImageModel{
			Path: "/" + filePath,
			Hash: imageHash,
			Name: filename,
		})
	}
	res.ResultOkWithData(resList, c)
}
