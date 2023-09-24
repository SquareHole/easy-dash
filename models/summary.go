package models

import "time"

type Summary struct {
	SummaryDescription string    `json:"summary_description"`
	FromDate           time.Time `json:"from_date"`
	ToDate             time.Time `json:"to_date"`
	NumberOfTests      int       `json:"number_of_tests"`
	SuccessfulTests    int       `json:"successful_tests"`
	FailedTests        int       `json:"failed_tests"`
	AverageDuration    float64   `json:"average_duration"`
}
