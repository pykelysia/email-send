package config

type (
	userConfig struct {
		UserEmail string `yaml:"UserEmail"`
		EmailPsw  string `yaml:"EmailPsw"`
		EmailHost string `yaml:"EmailHost"`
	}
	toEmail struct {
		Addresses []string `yaml:"Addresses"`
	}
	logConfig struct {
		LogPath string `yaml:"LogPath"`
	}
	routeConfig struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	}
	Config struct {
		UserConfig  userConfig  `yaml:"UserConfig"`
		EmailTo     toEmail     `yaml:"EmailTo"`
		LogConfig   logConfig   `yaml:"LogConfig"`
		RouteConfig routeConfig `yaml:"RouteConfig"`
	}
)
