package core

import (
	"GalleryMing/config"
	"GalleryMing/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "application.yaml"

// InitConf 初始化读取配置文件
func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %v", err))
	}
	// 读取
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("yaml.Unmarshal: %v", err)
	}
	log.Println("config yamlFile InitConf success")
	//TODO 设置global-config
	global.Config = c
}

func SetYaml() (err error) {
	marshal, err := yaml.Marshal(global.Config)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(ConfigFile, marshal, fs.ModePerm)
	if err != nil {
		return
	}
	global.Log.Info("config yamlFile SetYaml success d=====(￣▽￣*)b")
	return nil
}
