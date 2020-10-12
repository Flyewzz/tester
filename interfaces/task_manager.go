package interfaces

import "github.com/Flyewzz/tester/models"

type TaskManager interface {
	SetStatus(task *models.TaskInfo, user *models.User) error
	Undone(task *models.TaskInfo, user *models.User) error
}
