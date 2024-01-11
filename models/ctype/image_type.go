package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1 // 本地
	QiNiu ImageType = 2 // 七牛云
)

func (r ImageType) MarshlJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// 枚举
func (r ImageType) String() interface{} {
	var str string
	switch r {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛云"
	default:
		str = "未知"
	}
	return str
}
