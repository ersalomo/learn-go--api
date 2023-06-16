package model

import "time"
import "encoding/json"

type User struct {
	ID        int
	Gmail     string `json:"gmail" binding:"required,email"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string      `json:"name" binding:"required"`
	Grade     json.Number `json:"grade" binding:"required,number"`
}
