package models

import "time"

type TestStatus struct {
	Id              int       `json:"id"`
	StartedDateTime time.Time `json:"startedDateTime"`
	EndDateTime     time.Time `json:"endDateTime"`
	Duration        time.Time `json:"duration"`
	StatusCode      int       `json:"status_code"`
	BodyContent     string    `json:"body_content"`
	Successful      bool      `json:"successful"`
}
