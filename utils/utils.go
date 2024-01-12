package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"reflect"
	"time"
)

// InList 是否存在在列表中
func InList(key string, list []string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

// Md5 加密
func Md5(str []byte) string {
	m := md5.New()
	m.Write(str)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// GetValidMsg 返回结构体中的msg数据
func GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	// err断言为具体类型
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		for _, e := range errs {
			// 根据报错字段名 获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				return f.Tag.Get("msg")
			}
		}
	}
	return err.Error()
}

// Random 随机获取
func Random(list []uint) any {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(list))
	return list[randomIndex]
}
