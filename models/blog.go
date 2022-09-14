package models

import "time"

type Blog struct {
	Id      int64    `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Image   []string `json:"image"`
	//IsApproved bool      `json:"isapproved"`
	//Hidden     bool      `json:"hidden"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Reg_No    int64     `json:"reg_no"`
}
