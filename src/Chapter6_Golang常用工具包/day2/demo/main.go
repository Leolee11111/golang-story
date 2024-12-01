package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// AppConfig 是应用程序的主配置结构体，采用嵌套结构组织所有配置项
// 对应的 YAML 配置格式示例：
// app:
//
//	name: "MyApp"
//	version: "1.0.0"
//	port: 8080
type AppConfig struct {
	// App 存储应用程序的基本配置信息
	App struct {
		Name    string `mapstructure:"name"`    // 应用名称
		Version string `mapstructure:"version"` // 应用版本号
		Port    int    `mapstructure:"port"`    // 应用运行端口
	} `mapstructure:"app"` // mapstructure 标签用于映射 YAML 中的 "app" 节点

	// Database 存储数据库连接相关的配置
	// 对应 YAML 格式：
	// database:
	//   host: "localhost"
	//   port: 5432
	Database struct {
		Host     string `mapstructure:"host"`     // 数据库主机地址
		Port     int    `mapstructure:"port"`     // 数据库端口
		User     string `mapstructure:"user"`     // 数据库用户名
		Password string `mapstructure:"password"` // 数据库密码
		Name     string `mapstructure:"name"`     // 数据库名称
	} `mapstructure:"database"`

	// Features 存储功能特性的开关配置
	// 对应 YAML 格式：
	// features:
	//   cache: true
	//   metrics: false
	Features struct {
		Cache   bool `mapstructure:"cache"`   // 是否启用缓存功能
		Metrics bool `mapstructure:"metrics"` // 是否启用指标收集
	} `mapstructure:"features"`
}

// 使用示例：
// var config AppConfig
// config.App.Name        // 访问应用名称
// config.Database.Host   // 访问数据库主机
// config.Features.Cache  // 访问缓存特性开关

func main() {
	// 1. 设置默认值
	setDefaults()

	// 2. 初始化配置
	if err := initConfig(); err != nil {
		log.Fatal("无法加载配置:", err)
	}

	// 3. 启用配置热重载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件已变更: %s\n", e.Name)
		printConfig()
	})

	// 4. 绑定环境变量
	viper.BindEnv("app.port", "APP_PORT")
	viper.BindEnv("database.password", "DB_PASSWORD")

	// 5. 读取并打印配置
	printConfig()

	// 保持程序运行以观察配置热重载
	time.Sleep(time.Hour)
}

func setDefaults() {
	viper.SetDefault("app.name", "DefaultService")
	viper.SetDefault("app.port", 3000)
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.port", 5432)
}

func initConfig() error {
	viper.SetConfigName("config") // 配置文件名称（无扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 查找配置文件所在路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			return fmt.Errorf("配置文件未找到: %w", err)
		}
		// 其他错误
		return fmt.Errorf("读取配置文件错误: %w", err)
	}

	return nil
}

func printConfig() {
	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("无法解析配置: %v", err)
		return
	}

	fmt.Println("\n当前配置:")
	fmt.Printf("应用名称: %s\n", config.App.Name)
	fmt.Printf("版本: %s\n", config.App.Version)
	fmt.Printf("端口: %d\n", config.App.Port)
	fmt.Printf("数据库主机: %s\n", config.Database.Host)
	fmt.Printf("数据库端口: %d\n", config.Database.Port)
	fmt.Printf("数据库用户: %s\n", config.Database.User)
	fmt.Printf("缓存启用: %v\n", config.Features.Cache)
	fmt.Printf("指标启用: %v\n", config.Features.Metrics)
}
