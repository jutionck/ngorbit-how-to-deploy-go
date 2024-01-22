package model

import "time"

type Blog struct {
	ID        int       `json:"id"`
	UserId    int       `json:"userId"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
