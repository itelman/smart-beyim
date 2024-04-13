package models

import "time"

type Test struct {
	ID         int `json:"id"`
	UserID     int
	TestType   int       `json:"test_type"`
	Score      float32   `json:"score"`
	Feedback   string    `json:"feedback"`
	TimeSpent  time.Time `json:"time_spent"`
	TimePassed time.Time `json:"time_passed"`
}
