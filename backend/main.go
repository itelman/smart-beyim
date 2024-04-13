package main

import (
	"smart-beyim/database"
	"smart-beyim/models"
)

func main() {
	database.InsertUser(models.User{1, "abc", "abc"})
}
