package models

type Attempt struct {
	ID     int    `json:"id"`
	TaskId int    `json:"task_id"`
	Status string `json:"status"`
	Time   string `json:"time"`
}
