package scheduler

import (
	"context"
	"crypto/rand"
	"email-send/config"
	"email-send/engine"
	"fmt"
	"strings"
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
func (s *Scheduler) AddTask(to, subject, body string, sendTime time.Time) (*EmailTask, error) {
	// 解析时间格式
	targetTime := sendTime

	// 检查时间是否在将来
	if targetTime.Before(time.Now()) {
		return nil, fmt.Errorf("发送时间不能在过去")
	}

	task := &EmailTask{
		ID:        s.generateID(),
		To:        to,
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
	// 创建超时上下文
	ctx, cancel := context.WithDeadline(context.Background(), task.SendTime)
	defer cancel()

	// 等待到发送时间
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		// 时间到了，发送邮件
		s.sendEmail(task)
		return
	case <-ticker.C:
		// 每秒检查一次
	}

	// 更新状态为发送中
	s.mu.Lock()
	task.Status = TaskStatusSending
	s.mu.Unlock()

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
		fmt.Printf("[错误] 发送邮件失败: %v\n", err)
	} else {
		s.mu.Lock()
		task.Status = TaskStatusSent
		s.mu.Unlock()
		fmt.Printf("[成功] 邮件已发送至: %s\n", task.To)
	}
}

// CancelTask 取消任务
func (s *Scheduler) CancelTask(taskID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, task := range s.Tasks {
		if task.ID == taskID {
			if task.Status == TaskStatusPending {
				task.Status = TaskStatusCancelled
				return nil
			}
			return fmt.Errorf("任务已发送或正在发送，无法取消")
		}
	}

	return fmt.Errorf("未找到任务")
}

// GetTasks 获取所有任务
func (s *Scheduler) GetTasks() []*EmailTask {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 返回副本
	tasks := make([]*EmailTask, len(s.Tasks))
	copy(tasks, s.Tasks)
	return tasks
}

// GetPendingTasks 获取待发送任务
func (s *Scheduler) GetPendingTasks() []*EmailTask {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var pending []*EmailTask
	for _, task := range s.Tasks {
		if task.Status == TaskStatusPending {
			pending = append(pending, task)
		}
	}
	return pending
}

// CleanCompleted 清理已完成的旧任务
func (s *Scheduler) CleanCompleted() {
	s.mu.Lock()
	defer s.mu.Unlock()

	var remaining []*EmailTask
	for _, task := range s.Tasks {
		if task.Status == TaskStatusPending || task.Status == TaskStatusSending {
			remaining = append(remaining, task)
		}
	}
	s.Tasks = remaining
}

// ParseTime 解析时间字符串（格式: yyyy-mm-dd-hh-mm-ss）
func ParseTime(timeStr string) (time.Time, error) {
	times := strings.Split(timeStr, "-")
	if len(times) < 6 {
		return time.Time{}, fmt.Errorf("时间格式错误，应为 yyyy-mm-dd-hh-mm-ss")
	}

	year := 0
	month := 0
	day := 0
	hour := 0
	minute := 0
	second := 0

	fmt.Sscanf(times[0], "%d", &year)
	fmt.Sscanf(times[1], "%d", &month)
	fmt.Sscanf(times[2], "%d", &day)
	fmt.Sscanf(times[3], "%d", &hour)
	fmt.Sscanf(times[4], "%d", &minute)
	fmt.Sscanf(times[5], "%d", &second)

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local), nil
}
