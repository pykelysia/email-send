package engine

import "email-send/config"

func NewDefaultEmailEngine() *EmailEngine {
	return &EmailEngine{
		config: config.GetConfig(),
	}
}

func NewEmailEngine() *EmailEngine {
	return &EmailEngine{
		config: config.GetConfig(),
	}
}
