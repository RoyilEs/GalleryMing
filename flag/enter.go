package flag

import FLAG "flag"
import "github.com/fatih/structs"

type Option struct {
	DB bool
}

// Parse 解析命令行参数
func Parse() Option {
	db := FLAG.Bool("db", false, "初始化数据库")
	//解析命令写入注册的flag中
	FLAG.Parse()
	return Option{
		DB: *db,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, val := range maps {
		switch v := val.(type) {
		case string:
			if v != "" {
				f = true
			}
		case bool:
			if v == true {
				f = true
			}
		}
	}
	return
}

func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
		return
	}
}
