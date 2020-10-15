package postgres

import (
	"database/sql"

	"github.com/Flyewzz/tester/models"
)

type TaskManager struct {
	DB *sql.DB
}

func (this TaskManager) SetStatus(userId int, taskId int, status string) (int, error) {
	newAttemptId := 0
	err := this.DB.QueryRow(
		`INSERT INTO attempts (user_id, task_id, status)
		VALUES ($1, $2, $3) RETURNING id`, userId, taskId, status).
		Scan(&newAttemptId)
	return newAttemptId, err
}

func (this TaskManager) GetStatus(userId int, taskId int) (string, error) {
	var status string
	err := this.DB.QueryRow(
		`SELECT status FROM attempts
		WHERE user_id = $1 AND task_id = $2
		ORDER BY time DESC LIMIT 1;`).Scan(&status)
	return status, err
}

func (this TaskManager) GetUserAttempts(userId int) ([]models.Attempt, error) {
	rows, err := this.DB.Query(
		`SELECT id, task_id, status, time
		FROM attempts
		WHERE user_id = $1 ORDER BY time DESC`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var attempts []models.Attempt
	for rows.Next() {
		var attempt models.Attempt
		err = rows.Scan(&attempt.ID, &attempt.TaskId,
			&attempt.Status, &attempt.Time)
		if err != nil {
			continue //! May omit important records
		}
		attempts = append(attempts, attempt)
	}
	return attempts, nil
}
