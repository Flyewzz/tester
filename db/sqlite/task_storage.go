package sqlite

import (
	"database/sql"

	"github.com/Flyewzz/tester/models"
	_ "github.com/mattn/go-sqlite3"
)

type TaskStorage struct {
	Path string
}

func NewTaskStorage(path string) *TaskStorage {
	return &TaskStorage{
		Path: path,
	}
}

func (this TaskStorage) GetInfo(id int) (*models.TaskInfo, error) {
	db, err := sql.Open("sqlite3", this.Path)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var taskInfo models.TaskInfo
	err = db.QueryRow(
		`SELECT id, text, ram, 
		hdd, time, samples, limitations 
		FROM tasks WHERE id = $1 LIMIT 1`, id).
		Scan(&taskInfo.ID, &taskInfo.Text, &taskInfo.Ram,
			&taskInfo.HDD, &taskInfo.Time, &taskInfo.Samples,
			&taskInfo.Limitations,
		)

	return &taskInfo, err
}
