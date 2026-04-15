package looker

import (
	"email-send/config"
)

type (
	looker struct {
		c      config.Config
		err    chan error
		end    chan bool
		isOpen bool
	}
)
