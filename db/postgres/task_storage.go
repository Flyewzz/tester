package postgres

import (
	"database/sql"

	"github.com/Flyewzz/tester/models"
)

type TaskStorage struct {
	DB *sql.DB
}

func (this TaskStorage) GetInfo(id int) (*models.TaskInfo, int, error) {
	var taskInfo models.TaskInfo
	err := this.DB.QueryRow(
		`SELECT id, text, ram, 
		hdd, time, samples, limitations 
		FROM tasks WHERE id = $1 LIMIT 1`, id).
		Scan(&taskInfo.ID, &taskInfo.Text, &taskInfo.Ram,
			&taskInfo.HDD, &taskInfo.Time, &taskInfo.Samples,
			&taskInfo.Limitations,
		)
	if err != nil {
		return nil, 0, err
	}
	var taskCount int = 0
	//! Temporary solution. 2 requests is bad. Need to fix it as soon as possible.
	err = this.DB.QueryRow(
		`SELECT COUNT(*) FROM tasks;`).
		Scan(&taskCount)
	return &taskInfo, taskCount, err
}
