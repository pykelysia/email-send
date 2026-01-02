package looker

import (
	"email-send/config"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLooker(t *testing.T) {
	c := config.LoadConfig("../looker_test.yaml")
	l := GetLooker(c)
	l.Start()
	defer l.End()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := range 10 {
			l.Err(fmt.Errorf("test error %v in 1", i))
		}
		wg.Done()
	}()
	go func() {
		for i := range 10 {
			l.Err(fmt.Errorf("test error %v in 2", i))
		}
		wg.Done()
	}()
	wg.Wait()
	time.Sleep(time.Second * 10)
}
