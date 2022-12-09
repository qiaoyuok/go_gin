package po

import "time"

type Config struct {
	Mysql MysqlConfigPo `json:"mysql"`
	Redis RedisConfigPo `json:"redis"`
	Log   LogConfigPo   `json:"log"`
}

// MysqlConfigPo MySQL数据库配置
type MysqlConfigPo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Prefix   string `json:"prefix"`
}

// RedisConfigPo Redis配置
type RedisConfigPo struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Passwd string `json:"passwd"`
	Db     int    `json:"db"`
}

// LogConfigPo 日志配置
type LogConfigPo struct {
	Path           string        `json:"path"`
	ErrLogFileName string        `json:"err_log_file_name"`
	AppLogFileName string        `json:"app_log_file_name"`
	MaxAge         time.Duration `json:"max_age"`
}
