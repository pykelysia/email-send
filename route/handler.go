package route

import (
	"email-send/config"
	"email-send/scheduler"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 全局调度器（由 main.go 初始化后注入）
var GlobalScheduler *scheduler.Scheduler

func sendEmailHandler(c *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req  sendEmailReq
			resp sendEmailResp
		)
		err := ctx.ShouldBindBodyWithJSON(&req)
		if err != nil {
			resp = sendEmailResp{
				BaseMsg: baseMsg{
					Code:    http.StatusBadRequest,
					Message: "invalied request body",
				},
				Data: struct {
					IsSuccess bool `json:"is_success"`
				}{IsSuccess: false},
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}
		err = setSendEmailEngine(c, req.Time, req.Subject, req.Body)
		if err != nil {
			resp = sendEmailResp{
				BaseMsg: baseMsg{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("server set task failed: %v", err),
				},
				Data: struct {
					IsSuccess bool `json:"is_success"`
				}{IsSuccess: false},
			}
			ctx.JSON(http.StatusOK, resp)
			return
		}

		resp = sendEmailResp{
			BaseMsg: baseMsg{
				Code:    http.StatusOK,
				Message: "set task success",
			},
			Data: struct {
				IsSuccess bool `json:"is_success"`
			}{IsSuccess: true},
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func setSendEmailEngine(c *config.Config, t, s, b string) error {
	times := strings.Split(t, "-")
	if len(times) < 7 {
		return fmt.Errorf("时间格式错误，需要 yyyy-mm-dd-hh-mm-ss")
	}
	targetTime := time.Date(
		atoi(times[0]),
		time.Month(atoi(times[1])),
		atoi(times[2]),
		atoi(times[3]),
		atoi(times[4]),
		atoi(times[5]),
		atoi(times[6]),
		time.Local,
	)

	// 使用全局调度器添加任务
	_, err := GlobalScheduler.AddTask("", s, b, targetTime)
	return err
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
