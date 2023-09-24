package models

type UrlConfig struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	Description     string `json:"description"`
	Url             string `json:"url"`
	StatusCode      int    `json:"status_code"`
	BodyContains    string `json:"body_contains"`
	Enabled         bool   `json:"enabled"`
	ScheduleMinutes int    `json:"schedule_minutes"`
}
