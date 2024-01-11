package flag

import (
	"GalleryMing/global"
	"GalleryMing/models"
	"fmt"
)

func MakeMigrations() {
	var err error

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.ImageModel{},
		)
	if err != nil {
		global.Log.Error("[error] 生成数据库表结构失败", err)
		return
	}
	fmt.Println("MakeMigrations GalleryMing [ ok ]")
}
