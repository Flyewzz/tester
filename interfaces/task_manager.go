package interfaces

import "github.com/Flyewzz/tester/models"

type TaskManager interface {
	SetStatus(userId int, taskId int, status string) (int, error)
	GetStatus(userId int, taskId int) (string, error)
	GetUserAttempts(userId int) ([]models.Attempt, error)
}
