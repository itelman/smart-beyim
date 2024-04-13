package models

import "time"

type User struct {
	ID       int
	Name     string
	Password string
	Created  time.Time
}
