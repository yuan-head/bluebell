package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局变量。保存应用配置
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`    // 应用名称
	Mode         string `mapstructure:"mode"`    // 应用运行模式
	Version      string `mapstructure:"version"` // 应用版本
	Port         int    `mapstructure:"port"`    // 应用端口
	*LogConfig   `mapstructure:"log"`            // 日志配置
	*MySQLConfig `mapstructure:"mysql"`          // MySQL配置
	*RedisConfig `mapstructure:"redis"`          // Redis配置
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	MaxSize    int    `mapstructure:"max_size"`    // 日志文件最大尺寸，单位MB
	MaxAge     int    `mapstructure:"max_age"`     // 日志文件最大保存天数
	MaxBackups int    `mapstructure:"max_backups"` // 日志文件最大备份数
	Filename   string `mapstructure:"filename"`    // 日志文件名
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_connections"`
	MaxIdleConns int    `mapstructure:"max_idle_connections"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"` // Redis连接池大小
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml") // 设置配置文件路径
	//viper.SetConfigName("config") // 设置配置文件名
	//viper.SetConfigType("yaml")   // 设置配置文件类型
	viper.AddConfigPath(".")   // 添加当前目录作为配置文件的搜索路径
	err = viper.ReadInConfig() // 读取配置文件
	if err != nil {
		// 如果读取配置文件失败，打印错误信息并退出程序
		fmt.Println("viper.ReadInConfig() failed, err: %v ", err)
		return
		panic("Failed to read config file: " + err.Error())
	}
	// 将配置文件中的内容映射到结构体中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper.Unmarshal() failed, err: %v ", err)
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) { // 配置文件发生变更之后会调用的回调函数
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper.Unmarshal() failed, err: %v ", err)
		}
		fmt.Println("Config file changed:", in.Name)
	})
	return
}
