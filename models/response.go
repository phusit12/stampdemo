package models

import "time"

type GetUserRes struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}
type MessageUserRes struct {
	Message string `json:"message"`
}
