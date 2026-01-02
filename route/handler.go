package route

import (
	"context"
	"email-send/engine"
	"email-send/looker"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func sendEmailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req  sendEmailReq
			resp sendEmailResp
		)
		err := c.ShouldBindBodyWithJSON(&req)
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
			c.JSON(http.StatusOK, resp)
			return
		}
		err = setSendEmailEngine(req.Time, req.Subject, req.Body)
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
			c.JSON(http.StatusOK, resp)
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
		c.JSON(http.StatusOK, resp)
	}
}

func setSendEmailEngine(t, s, b string) error {
	times := strings.Split(t, "-")
	if len(times) < 7 {
		return fmt.Errorf("1")
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

	ctx, cancel := context.WithDeadline(context.Background(), targetTime)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				e := engine.NewDefaultEmailEngine()
				err := e.SendMail(s, b)
				looker.Looker.Err <- err
			case <-ticker.C:
				continue
			}
		}
	}()

	return nil
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
