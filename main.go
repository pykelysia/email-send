package main

import (
	"email-send/config"
	"email-send/route"
	"email-send/scheduler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 加载配置
	c := config.LoadConfig("./emailsend.yaml")

	// 创建调度器
	sched := scheduler.NewScheduler(c)

	// 注入全局调度器到 route 包
	route.GlobalScheduler = sched

	// 启动 Web 服务
	go func() {
		g := route.NewG(c)
		log.Printf("启动 Web 服务: %s:%s", c.RouteConfig.Host, c.RouteConfig.Port)
		if err := g.Run(); err != nil {
			log.Printf("Web 服务异常: %v", err)
		}
	}()

	log.Println("Web API 已启动: POST /send")
	log.Println("按 Ctrl+C 退出程序")

	// 处理退出信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("\n正在退出...")
}
