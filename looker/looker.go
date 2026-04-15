package looker

import (
	"email-send/config"
	"email-send/util"
	"fmt"
	"time"

	"github.com/pykelysia/pyketools"
)

func GetLooker(c *config.Config) *looker {
	err, end := make(chan error), make(chan bool)

	l := util.NewLogger(c)

	return &looker{
		err:    err,
		end:    end,
		isOpen: false,
		logger: l,
	}
}

func (l *looker) Start() {
	l.isOpen = true
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for l.isOpen {
			select {
			case e := <-l.err:
				pyketools.Infof("get error: %v", e)
				l.logger.LogToFile("error", fmt.Sprintf("%v", e))
				continue
			case <-l.end:
				continue
			case <-ticker.C:
				pyketools.Infof("get tick")
				continue
			}
		}
	}()
}

func (l *looker) End() {
	if l.isOpen {
		l.end <- false
		l.isOpen = false
		close(l.end)
		close(l.err)
		l.logger.Close()
	}
}

func (l *looker) Err(err error) {
	if l.isOpen {
		l.err <- err
	}
}
