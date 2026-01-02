package looker

import (
	"email-send/config"
	"email-send/util"
)

type (
	looker struct {
		c      config.Config
		err    chan (error)
		end    chan (bool)
		isOpen bool
		logger *util.Logger
	}
)
