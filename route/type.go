package route

import (
	"email-send/config"

	"github.com/gin-gonic/gin"
)

type (
	G struct {
		c      *config.Config
		server *gin.Engine
		host   string
		port   string
	}
	baseMsg struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	sendEmailReq struct {
		Time    string `json:"time"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}
	sendEmailResp struct {
		BaseMsg baseMsg `json:"base_msg"`
		Data    struct {
			IsSuccess bool `json:"is_success"`
		} `json:"data"`
	}
)
