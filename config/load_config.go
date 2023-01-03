package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go_gin/dal/po"
	"go_gin/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var C po.Config
var DB *gorm.DB
var RDB *redis.Client
var once sync.Once

// ConfigFilename 配置文件路径
const ConfigFilename = "/Users/ourschooldev-01/GolandProjects/school/src/go_gin/config/config.yaml"

// InitEnv 初始化环境
func InitEnv() {
	once.Do(func() {
		LoadConfig()
		initMysql()
		initRedis()
	})
}

// LoadConfig 加载配置文件
func LoadConfig() {
	viper.SetConfigFile(ConfigFilename) // 指定配置文件路径

	if err := viper.ReadInConfig(); err != nil { // 查找并读取配置文件
		fmt.Printf("%s----%#v\n", "加载配置文件出错！", err)
		return
	}

	if err := viper.Unmarshal(&C); err != nil {
		fmt.Printf("%s----%#v\n", "配置文件解析到结构体出错！", err)
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	return
}

// initMysql 初始化MYSQL
func initMysql() {
	db, err := gorm.Open(mysql.Open(GetDbDsn()), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库出错：", err)
		return
	}
	DB = db
	query.SetDefault(db.Debug())
	return
}

// initRedis 初始化Redis
func initRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     C.Redis.Host,
		Password: C.Redis.Passwd, // no password set
		DB:       C.Redis.Db,     // use default DB
	})
}

// GetDbDsn 获取DSN链接
func GetDbDsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", C.Mysql.Username, C.Mysql.Passwd, C.Mysql.Host, C.Mysql.Port, C.Mysql.Db)
	return
}
