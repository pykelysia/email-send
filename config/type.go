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
	Config struct {
		UserConfig userConfig `yaml:"UserConfig"`
		EmailTo    toEmail    `yaml:"EmailTo"`
	}
)
