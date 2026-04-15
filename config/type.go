package config

// UserConfig 用户邮件配置
type UserConfig struct {
	UserEmail string `mapstructure:"UserEmail"`
	EmailPsw  string `mapstructure:"EmailPsw"`
	EmailHost string `mapstructure:"EmailHost"`
}

// ToEmailConfig 收件人配置
type ToEmailConfig struct {
	Addresses []string `mapstructure:"Addresses"`
}

// LogConfig 日志配置
type LogConfig struct {
	LogPath string `mapstructure:"LogPath"`
}

// RouteConfig 路由配置
type RouteConfig struct {
	Host string `mapstructure:"Host"`
	Port string `mapstructure:"Port"`
}

// Config 应用配置
type Config struct {
	UserConfig  UserConfig    `mapstructure:"UserConfig"`
	EmailTo     ToEmailConfig `mapstructure:"EmailTo"`
	LogConfig   LogConfig     `mapstructure:"LogConfig"`
	RouteConfig RouteConfig   `mapstructure:"RouteConfig"`
}
