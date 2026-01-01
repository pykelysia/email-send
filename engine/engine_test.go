package engine

import (
	"testing"

	"github.com/pykelysia/pyketools"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestNewEmailEngine(t *testing.T) {
	engine := NewEmailEngine("../emailsend.yaml")
	pyketools.Infof("loaded config: %v", engine.config)
}
