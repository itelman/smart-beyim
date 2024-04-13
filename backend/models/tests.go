package models

import "time"

type Test struct {
	ID         int
	UserID     int
	TestTypeID int
	Score      float32
	Feedback   string
	TimeSpent  time.Time
	TimePassed time.Time
}

type Ielts struct {
	ID           int
	UserID       int
	Reading_ID   float32
	Listening_ID float32
	Writing_ID   float32
	Speaking_ID  float32
}
