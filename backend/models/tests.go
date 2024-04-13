package models

import "time"

type Ielts struct {
	ID         int       `json:"id"`
	TimePassed time.Time `json:"time_passed"`
	UserID     int
	Reading    Test `json:"reading"`
	Writing    Test `json:"writing"`
	Listening  Test `json:"listening"`
	Speaking   Test `json:"speaking"`
}

type Test struct {
	ID         int `json:"id"`
	UserID     int
	TestType   int       `json:"test_type"`
	Score      float32   `json:"score"`
	Feedback   string    `json:"feedback"`
	TimeSpent  time.Time `json:"time_spent"`
	TimePassed time.Time `json:"time_passed"`
}
