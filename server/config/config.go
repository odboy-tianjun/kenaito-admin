package config

import (
	"github.com/odboy-tianjun/kenaito-admin/server/util"
	"github.com/spf13/viper"
)

type Config struct {
	Port     int  `mapstructure:"port"`
	Debug    bool `mapstructure:"debug"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Database int    `mapstructure:"database"`
	} `mapstructure:"redis"`
}

// LoadConfig 从指定路径加载配置文件
func LoadConfig(path string) *Config {
	// 设置配置文件名称和路径
	viper.SetConfigName("config") // 配置文件名称（无扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(path)     // 查找配置文件的路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		util.ThrowError("配置文件 config.yml 读取失败", err)
	}

	// 反序列化配置
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		util.ThrowError("配置文件 config.yml 解析失败", err)
	}

	return &config
}
