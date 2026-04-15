package scheduler

import "time"

// EmailTask 邮件任务
type EmailTask struct {
	ID        string     // 任务 ID
	Subject   string     // 主题
	Body      string     // 正文
	SendTime  time.Time  // 发送时间
	Status    TaskStatus // 状态
	CreatedAt time.Time  // 创建时间
}

// TaskStatus 任务状态
type TaskStatus int

const (
	TaskStatusPending   TaskStatus = iota // 等待中
	TaskStatusSending                     // 发送中
	TaskStatusSent                        // 已发送
	TaskStatusFailed                      // 发送失败
	TaskStatusCancelled                   // 已取消
)

// TaskStatusToString 状态转字符串
func (s TaskStatus) String() string {
	switch s {
	case TaskStatusPending:
		return "等待中"
	case TaskStatusSending:
		return "发送中"
	case TaskStatusSent:
		return "已发送"
	case TaskStatusFailed:
		return "发送失败"
	case TaskStatusCancelled:
		return "已取消"
	default:
		return "未知"
	}
}
