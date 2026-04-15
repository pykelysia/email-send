package main

import (
	"email-send/config"
	"email-send/route"
	"email-send/scheduler"
	"email-send/util"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 加载配置
	c := config.LoadConfig("./emailsend.yaml")

	// 初始化日志
	if err := util.InitLogger(c); err != nil {
		panic("初始化日志失败: " + err.Error())
	}

	// 创建调度器
	sched := scheduler.NewScheduler(c)

	// 注入全局调度器到 route 包
	route.GlobalScheduler = sched

	// 启动 Web 服务
	go func() {
		g := route.NewG(c)
		util.Infof("启动 Web 服务: %s:%s", c.RouteConfig.Host, c.RouteConfig.Port)
		if err := g.Run(); err != nil {
			util.Errorf("Web 服务异常: %v", err)
		}
	}()

	util.Info("Web API 已启动: POST /send")
	util.Info("按 Ctrl+C 退出程序")

	// 处理退出信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	util.Info("正在退出...")

	// 刷新日志缓冲区
	util.Sync()
}
