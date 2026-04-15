package scheduler

import (
	"crypto/rand"
	"email-send/config"
	"email-send/engine"
	"email-send/util"
	"fmt"
	"sync"
	"time"
)

// Scheduler 调度器
type Scheduler struct {
	Tasks       []*EmailTask        // 任务列表
	config      *config.Config      // 配置
	mu          sync.RWMutex        // 读写锁
	emailEngine *engine.EmailEngine // 邮件引擎
}

// NewScheduler 创建调度器
func NewScheduler(c *config.Config) *Scheduler {
	return &Scheduler{
		Tasks:  make([]*EmailTask, 0),
		config: c,

		emailEngine: engine.NewEmailEngine(),
	}
}

// generateID 生成唯一Id
func (s *Scheduler) generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// AddTask 添加任务
func (s *Scheduler) AddTask(subject, body string, sendTime time.Time) (*EmailTask, error) {
	// 解析时间格式
	targetTime := sendTime

	// 检查时间是否在将来
	if targetTime.Before(time.Now()) {
		return nil, fmt.Errorf("发送时间不能在过去")
	}

	task := &EmailTask{
		ID:        s.generateID(),
		Subject:   subject,
		Body:      body,
		SendTime:  targetTime,
		Status:    TaskStatusPending,
		CreatedAt: time.Now(),
	}

	s.mu.Lock()
	s.Tasks = append(s.Tasks, task)
	s.mu.Unlock()

	// 启动定时任务
	go s.runTask(task)

	return task, nil
}

// runTask 运行单个任务
func (s *Scheduler) runTask(task *EmailTask) {
	// 计算等待时间
	waitDuration := time.Until(task.SendTime)

	// 如果已经到时间或过期，立即发送
	if waitDuration <= 0 {
		s.sendEmail(task)
		return
	}

	// 使用 Timer 等待到指定时间
	timer := time.NewTimer(waitDuration)
	defer timer.Stop()

	// 等待定时器触发
	<-timer.C

	// 时间到了，发送邮件
	s.sendEmail(task)
}

// sendEmail 发送邮件
func (s *Scheduler) sendEmail(task *EmailTask) {
	s.mu.Lock()
	task.Status = TaskStatusSending
	s.mu.Unlock()

	err := s.emailEngine.SendMail(task.Subject, task.Body)
	if err != nil {
		s.mu.Lock()
		task.Status = TaskStatusFailed
		s.mu.Unlock()
		util.Errorf("[失败] 发送邮件失败: 主题=%s, 错误=%v", task.Subject, err)
	} else {
		s.mu.Lock()
		task.Status = TaskStatusSent
		s.mu.Unlock()
		util.Infof("[成功] 邮件已发送: 主题=%s", task.Subject)
	}
}
