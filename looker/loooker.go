package looker

import "time"

var (
	Looker looker
)

func (l looker) Start() {
	go func() {
		ticker := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-l.Err:
				continue
			case <-ticker.C:
				continue
			}
		}
	}()
}
