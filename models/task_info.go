package models

type TaskInfo struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Ram         string `json:"ram"`
	HDD         string `json:"hdd"`
	Time        string `json:"time"`
	Samples     string `json:"samples"`
	Limitations string `json:"limitations"`
}
