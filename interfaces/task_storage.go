package interfaces

import (
	. "github.com/Flyewzz/tester/models"
)

type TaskStorage interface {
	GetInfo(id int) (*TaskInfo, int, error)
}
