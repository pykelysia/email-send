package route

import (
	"email-send/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pykelysia/pyketools"
)

func NewG(c *config.Config) *G {
	return &G{
		c:      c,
		server: gin.Default(),
		host:   c.RouteConfig.Host,
		port:   c.RouteConfig.Port,
	}
}

func (g *G) bindRoute() {
	server := g.server
	server.POST("/send", sendEmailHandler(g.c))
}

func (g *G) Run() error {
	g.bindRoute()
	ip := fmt.Sprintf("%s:%s", g.host, g.port)
	err := g.server.Run(ip)
	if err != nil {
		pyketools.Errorf("failed to run gin: %v", err)
	}
	return err
}
