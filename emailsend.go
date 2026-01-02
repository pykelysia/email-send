package main

import (
	"email-send/config"
	"email-send/route"

	"github.com/pykelysia/pyketools"
)

func main() {
	c := config.LoadConfig("./develop.yaml")
	g := route.NewG(c)
	err := g.Run()
	if err != nil {
		pyketools.Fatalf("failed to open net: %v", err)
	}
}
