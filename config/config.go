package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// 全局配置实例
var globalConfig *Config

// InitConfig 初始化配置
// 支持从指定路径加载配置文件，或从默认路径加载
func InitConfig(configPath string) (*Config, error) {
	v := viper.New()

	// 设置配置文件的名称（不含扩展名）
	v.SetConfigName("emailsend")

	// 添加配置文件路径
	if configPath != "" {
		v.AddConfigPath(configPath)
	}

	// 从当前目录查找配置
	v.AddConfigPath(".")

	// 设置配置类型
	v.SetConfigType("yaml")

	// 设置默认值
	setDefaults(v)

	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置到结构体
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 设置全局配置
	globalConfig = &config

	return &config, nil
}

// LoadConfig 加载配置（兼容旧接口）
// configName: 配置文件名（不含扩展名），如 "emailsend" 或 "develop"
func LoadConfig(configName string) *Config {
	// 如果已存在全局配置，直接返回
	if globalConfig != nil {
		return globalConfig
	}

	v := viper.New()

	// 设置配置文件的名称（不含扩展名）
	baseName := strings.TrimSuffix(configName, ".yaml")
	v.SetConfigName(baseName)

	// 从当前目录查找配置
	v.AddConfigPath(".")

	// 设置配置类型
	v.SetConfigType("yaml")

	// 设置默认值
	setDefaults(v)

	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败: %v", err))
	}

	// 解析配置到结构体
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("解析配置文件失败: %v", err))
	}

	// 设置全局配置
	globalConfig = &config

	return &config
}

// GetConfig 获取全局配置
func GetConfig() *Config {
	if globalConfig == nil {
		panic("配置未初始化，请先调用 LoadConfig 或 InitConfig")
	}
	return globalConfig
}

// setDefaults 设置默认值
func setDefaults(v *viper.Viper) {
	// 用户配置默认值
	v.SetDefault("UserConfig.UserEmail", "")
	v.SetDefault("UserConfig.EmailPsw", "")
	v.SetDefault("UserConfig.EmailHost", "")

	// 收件人配置默认值
	v.SetDefault("EmailTo.Addresses", []string{})

	// 日志配置默认值
	v.SetDefault("LogConfig.LogPath", "./EmailSend.log")

	// 路由配置默认值
	v.SetDefault("RouteConfig.Host", "")
	v.SetDefault("RouteConfig.Port", "8080")
}

// ReloadConfig 重新加载配置
func ReloadConfig() error {
	if globalConfig == nil {
		return fmt.Errorf("配置未初始化")
	}

	v := viper.New()

	// 获取配置文件路径
	configFile := v.ConfigFileUsed()
	if configFile == "" {
		return fmt.Errorf("无法获取配置文件路径")
	}

	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	// 重新读取
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("重新读取配置文件失败: %w", err)
	}

	// 重新解析
	if err := v.Unmarshal(globalConfig); err != nil {
		return fmt.Errorf("重新解析配置文件失败: %w", err)
	}

	return nil
}
