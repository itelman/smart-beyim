package models

type Reading struct {
	ID       int
	UserID   int
	Score    float32
	Feedback string
	IsSolo   int
}

type Listening struct {
	ID       int
	UserID   int
	Score    float32
	Feedback string
	IsSolo   int
}

type Writing struct {
	ID       int
	UserID   int
	Score    float32
	Feedback string
	IsSolo   int
}

type Speaking struct {
	ID       int
	UserID   int
	Score    float32
	Feedback string
	IsSolo   int
}
