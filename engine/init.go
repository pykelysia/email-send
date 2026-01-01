package engine

import "email-send/config"

func NewDefaultEmailEngine() *EmailEngine {
	return &EmailEngine{
		config: config.LoadConfig("./emailsend.yaml"),
	}
}

func NewEmailEngine(filePath string) *EmailEngine {
	return &EmailEngine{
		config: config.LoadConfig(filePath),
	}
}
