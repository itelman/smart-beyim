package models

type Reading struct {
	ID       int
	UserID   int
	Score    float32
	Feedback string
	IsSolo   int
}
