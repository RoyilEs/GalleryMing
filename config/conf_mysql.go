package config

import "strconv"

type MySql struct {
	DataBase   string `yaml:"dataBase"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	UserName   string `yaml:"userName"`
	Password   string `yaml:"password"`
	DriverName string `yaml:"driverName"`
	LogLevel   string `yaml:"logLevel"` // 日志等级 dubug info warn error
}

// Dsn 实例化配置 Dsn
func (m *MySql) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DataBase + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
}
