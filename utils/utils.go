package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"reflect"
)

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
func Random(min, max int) int {
	return rand.Intn(max-min) + min
}
